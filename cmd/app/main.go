package main

import (
	"context"
	"fmt"
	"github.com/lenvendo/ig-absolut-api/internal/repository/token"
	"github.com/lenvendo/ig-absolut-api/internal/repository/users"

	"net/http"
	"os"
	"time"

	"github.com/lenvendo/ig-absolut-api/configs"
	"github.com/lenvendo/ig-absolut-api/internal/db"
	"github.com/lenvendo/ig-absolut-api/internal/server"
	"github.com/lenvendo/ig-absolut-api/pkg/api"
	"github.com/lenvendo/ig-absolut-api/pkg/health"
	"github.com/lenvendo/ig-absolut-api/tools/logging"
	"github.com/lenvendo/ig-absolut-api/tools/metrics"
	"github.com/lenvendo/ig-absolut-api/tools/sentry"
	"github.com/lenvendo/ig-absolut-api/tools/tracing"

	"github.com/go-kit/kit/log/level"
	"github.com/nats-io/nats.go"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config
	cfg := configs.NewConfig()
	if err := cfg.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "read config: %s", err)
		os.Exit(1)
	}
	// Print config
	if err := cfg.Print(); err != nil {
		fmt.Fprintf(os.Stderr, "read config: %s", err)
		os.Exit(1)
	}

	logger, err := logging.NewLogger(cfg.Logger.Level, cfg.Logger.TimeFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger: %s", err)
		os.Exit(1)
	}
	ctx = logging.WithContext(ctx, logger)

	if cfg.Tracer.Enabled {
		tracer, closer, err := tracing.NewJaegerTracer(
			ctx,
			fmt.Sprintf("%s:%d", cfg.Tracer.Host, cfg.Tracer.Port),
			cfg.Tracer.Name,
		)
		if err != nil {
			level.Error(logger).Log("err", err, "msg", "failed to init tracer")
		}
		defer closer.Close()
		ctx = tracing.WithContext(ctx, tracer)
	}
	if cfg.Sentry.Enabled {
		if err := sentry.NewSentry(cfg); err != nil {
			level.Error(logger).Log("err", err, "msg", "failed to init sentry")
		}
	}

	if cfg.Metrics.Enabled {
		ctx = metrics.WithContext(ctx)
	}

	dbConn, err := db.Connect(ctx, cfg)
	if err != nil {
		level.Error(logger).Log("init", "server", "err", err)
		os.Exit(1)
	}
	defer db.Close(ctx, dbConn)

	nc, mainErr := nats.Connect(
		fmt.Sprintf("nats://%s:%d", cfg.Nats.Host, cfg.Nats.Port),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(cfg.Nats.RetryLimit),
		nats.ReconnectWait(time.Millisecond*time.Duration(cfg.Nats.WaitLimit)),
		nats.UserInfo(cfg.Nats.UserName, cfg.Nats.Password),
	)
	if mainErr != nil {
		_ = level.Error(logger).Log("failed to connect nats", mainErr)
		return
	}
	defer nc.Close()

	usersRepository := users.NewRepository(ctx, dbConn)
	tokensRepository := token.NewRepository(ctx, dbConn)
	apiService := initApiService(ctx, cfg, usersRepository, tokensRepository, nc)
	healthService := initHealthService(ctx, cfg)

	s, err := server.NewServer(
		server.SetConfig(cfg),
		server.SetLogger(logger),
		server.SetHandler(
			map[string]http.Handler{

				"api":    api.MakeHTTPHandler(ctx, apiService),
				"health": health.MakeHTTPHandler(ctx, healthService),
			}),
		server.SetGRPC(
			api.JoinGRPC(ctx, apiService),
		),
	)
	if err != nil {
		level.Error(logger).Log("init", "server", "err", err)
		os.Exit(1)
	}
	defer s.Close()

	if err := s.AddHTTP(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	if err = s.AddGRPC(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	if err = s.AddMetrics(); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}

	s.AddSignalHandler()
	s.Run()
}

func initApiService(
	ctx context.Context,
	cfg *configs.Config,
	users users.Repository,
	tokens token.Repository,
	nats *nats.Conn,
) api.Service {
	apiService := api.NewApiService(users, tokens, nats)
	if cfg.Metrics.Enabled {
		apiService = api.NewMetricsService(ctx, apiService)
	}
	apiService = api.NewLoggingService(ctx, apiService)
	if cfg.Tracer.Enabled {
		apiService = api.NewTracingService(ctx, apiService)
	}
	if cfg.Sentry.Enabled {
		apiService = api.NewSentryService(apiService)
	}
	return apiService
}

func initHealthService(_ context.Context, _ *configs.Config) health.Service {
	return health.NewHealthService()
}
