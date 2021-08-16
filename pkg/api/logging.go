package api

import (
	"context"
	"time"

	"github.com/lenvendo/ig-absolut-api/tools/logging"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(ctx context.Context, s Service) Service {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "component", "api")
	return &loggingService{logger, s}
}

type logged interface {
	Log() []interface{}
}

type loggingService struct {
	logger log.Logger
	Service
}

func (s *loggingService) getLog(req interface{}, resp interface{}) (out []interface{}) {
	if logger, ok := interface{}(req).(logged); ok {
		out = append(out, logger.Log()...)
	}

	if logger, ok := interface{}(resp).(logged); ok {
		out = append(out, logger.Log()...)
	}

	return
}

func (s *loggingService) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "ApiUserConfirm",
			"took", time.Since(begin),
		)

		m = append(m, s.getLog(req, resp)...)

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
	return s.Service.ApiUserConfirm(ctx, req)
}

func (s *loggingService) ApiUserProfile(ctx context.Context, req *UserProfileRequest) (resp *UserProfileResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "ApiUserProfile",
			"took", time.Since(begin),
		)

		m = append(m, s.getLog(req, resp)...)

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
	return s.Service.ApiUserProfile(ctx, req)
}

func (s *loggingService) ApiUserRegistration(ctx context.Context, req *UserRegRequest) (resp *UserRegResponse, err error) {
	defer func(begin time.Time) {
		m := getInfoFromContext(ctx)
		m = append(m,
			"code", getHTTPStatusCode(err),
			"method", "ApiUserRegistration",
			"took", time.Since(begin),
		)

		m = append(m, s.getLog(req, resp)...)

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
	return s.Service.ApiUserRegistration(ctx, req)
}

func getInfoFromContext(ctx context.Context) []interface{} {
	m := make([]interface{}, 0)
	{
		val := ctx.Value(ContextGRPCKey{})
		if _, ok := val.(GRPCInfo); ok {
			m = append(m, "protocol", "GRPC")
		}
	}

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
