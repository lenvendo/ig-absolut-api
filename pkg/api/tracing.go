package api

import (
	"context"

	"github.com/lenvendo/ig-absolut-api/tools/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// NewTracingService returns an instance of an instrumenting Service.
func NewTracingService(ctx context.Context, s Service) Service {
	tracer := tracing.FromContext(ctx)
	return &tracingService{tracer, s}
}

type trace interface {
	Span() []interface{}
}

func (s *tracingService) getTrace(req interface{}, resp interface{}) (out []interface{}) {
	if val, ok := interface{}(req).(trace); ok {
		out = append(out, val.Span()...)
	}

	if val, ok := interface{}(resp).(trace); ok {
		out = append(out, val.Span()...)
	}

	return
}

type tracingService struct {
	tracer opentracing.Tracer
	Service
}

func (s *tracingService) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "ApiUserConfirm")
	span.LogFields(log.Object("tracingService", s.getTrace(req, resp)))
	defer span.Finish()
	return s.Service.ApiUserConfirm(ctx, req)
}

func (s *tracingService) ApiUserProfile(ctx context.Context, req *UserProfileRequest) (resp *UserProfileResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "ApiUserProfile")
	span.LogFields(log.Object("tracingService", s.getTrace(req, resp)))
	defer span.Finish()
	return s.Service.ApiUserProfile(ctx, req)
}

func (s *tracingService) ApiUserRegistration(ctx context.Context, req *UserRegRequest) (resp *UserRegResponse, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "ApiUserRegistration")
	span.LogFields(log.Object("tracingService", s.getTrace(req, resp)))
	defer span.Finish()
	return s.Service.ApiUserRegistration(ctx, req)
}
