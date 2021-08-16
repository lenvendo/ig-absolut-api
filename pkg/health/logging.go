package health

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/lenvendo/ig-absolut-api/tools/logging"
)

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(ctx context.Context, s Service) Service {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "component", "health")
	return &loggingService{logger, s}
}

type loggingService struct {
	logger log.Logger
	Service
}

func (s *loggingService) Liveness(ctx context.Context, req *LivenessRequest) (resp *LivenessResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "Liveness",
			"took", time.Since(begin),
		)

		if getHTTPStatusCode(err) == 404 {
			m = append(m, "msg", err)
			level.Warn(s.logger).Log(m...)
		} else if err != nil {
			m = append(m, "err", err)
			level.Error(s.logger).Log(m...)
		} else {
			level.Info(s.logger).Log(m...)
		}
	}(time.Now())
	return s.Service.Liveness(ctx, req)
}

func (s *loggingService) Readiness(ctx context.Context, req *ReadinessRequest) (resp *ReadinessResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "Readiness",
			"took", time.Since(begin),
		)

		if getHTTPStatusCode(err) == 404 {
			m = append(m, "msg", err)
			level.Warn(s.logger).Log(m...)
		} else if err != nil {
			m = append(m, "err", err)
			level.Error(s.logger).Log(m...)
		} else {
			level.Info(s.logger).Log(m...)
		}
	}(time.Now())
	return s.Service.Readiness(ctx, req)
}

func (s *loggingService) Version(ctx context.Context, req *VersionRequest) (resp *VersionResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "Version",
			"took", time.Since(begin),
		)

		if getHTTPStatusCode(err) == 404 {
			m = append(m, "msg", err)
			level.Warn(s.logger).Log(m...)
		} else if err != nil {
			m = append(m, "err", err)
			level.Error(s.logger).Log(m...)
		} else {
			level.Info(s.logger).Log(m...)
		}
	}(time.Now())
	return s.Service.Version(ctx, req)
}

func getInfoFromContext(ctx context.Context) []interface{} {
	m := make([]interface{}, 0)

	{
		val := ctx.Value(ContextHTTPKey{})
		if i, ok := val.(HTTPInfo); ok {
			m = append(m,
				// "protocol", i.Protocol,
				// "http_method", i.Method,
				// "from", i.From,
				"url", i.URL,
			)
		}
	}

	return m
}
