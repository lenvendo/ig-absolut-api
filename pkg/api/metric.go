package api

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
	tool "github.com/lenvendo/ig-absolut-api/tools/metrics"
)

// NewMetricService returns an instance of an instrumenting Service.
func NewMetricsService(ctx context.Context, s Service) Service {
	counter, latency := tool.FromContext(ctx)
	return &metricService{counter, latency, s}
}

type metricService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (s *metricService) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "api", "handler", "ApiUserConfirm", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "api", "handler", "ApiUserConfirm", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.ApiUserConfirm(ctx, req)
}

func (s *metricService) ApiUserProfile(ctx context.Context, req *UserProfileRequest) (resp *UserProfileResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "api", "handler", "ApiUserProfile", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "api", "handler", "ApiUserProfile", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.ApiUserProfile(ctx, req)
}

func (s *metricService) ApiUserRegistration(ctx context.Context, req *UserRegRequest) (resp *UserRegResponse, err error) {
	defer func(begin time.Time) {
		go func() {
			s.requestCount.With("service", "api", "handler", "ApiUserRegistration", "code", strconv.Itoa(getHTTPStatusCode(err))).Add(1)
			s.requestLatency.With("service", "api", "handler", "ApiUserRegistration", "code", strconv.Itoa(getHTTPStatusCode(err))).Observe(time.Since(begin).Seconds())
		}()
	}(time.Now())
	return s.Service.ApiUserRegistration(ctx, req)
}
