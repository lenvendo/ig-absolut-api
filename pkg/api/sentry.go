package api

import (
	"context"
	"strconv"

	"github.com/getsentry/sentry-go"
)

func NewSentryService(s Service) Service {
	return &sentryService{s}
}

type sentryService struct {
	Service
}

type sentryLog interface {
	SentryLog() []interface{}
}

func (s *sentryService) getSentryLog(req interface{}, resp interface{}) (out map[string][]interface{}) {
	out = make(map[string][]interface{})
	if sentry, ok := interface{}(req).(sentryLog); ok {
		out["request"] = append(out["request"], sentry.SentryLog()...)
	}

	if sentry, ok := interface{}(resp).(sentryLog); ok {
		out["response"] = append(out["response"], sentry.SentryLog()...)
	}
	return
}

func (s *sentryService) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	defer func() {
		if err != nil {
			log := s.getSentryLog(req, resp)
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetTag("code", strconv.Itoa(getHTTPStatusCode(err)))
				scope.SetTag("method", "ApiUserConfirm")
				scope.SetExtra("request", log["request"])
				scope.SetExtra("response", log["response"])
			})
			sentry.CaptureException(err)
		}
	}()
	return s.Service.ApiUserConfirm(ctx, req)
}

func (s *sentryService) ApiUserProfile(ctx context.Context, req *UserProfileRequest) (resp *UserProfileResponse, err error) {
	defer func() {
		if err != nil {
			log := s.getSentryLog(req, resp)
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetTag("code", strconv.Itoa(getHTTPStatusCode(err)))
				scope.SetTag("method", "ApiUserProfile")
				scope.SetExtra("request", log["request"])
				scope.SetExtra("response", log["response"])
			})
			sentry.CaptureException(err)
		}
	}()
	return s.Service.ApiUserProfile(ctx, req)
}

func (s *sentryService) ApiUserRegistration(ctx context.Context, req *UserRegRequest) (resp *UserRegResponse, err error) {
	defer func() {
		if err != nil {
			log := s.getSentryLog(req, resp)
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetTag("code", strconv.Itoa(getHTTPStatusCode(err)))
				scope.SetTag("method", "ApiUserRegistration")
				scope.SetExtra("request", log["request"])
				scope.SetExtra("response", log["response"])
			})
			sentry.CaptureException(err)
		}
	}()
	return s.Service.ApiUserRegistration(ctx, req)
}
