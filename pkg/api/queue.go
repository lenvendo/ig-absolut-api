package api

import (
	"context"
	"encoding/json"

	"github.com/lenvendo/ig-absolut-api/tools/logging"
	"github.com/go-kit/kit/log"
	kitnats "github.com/go-kit/kit/transport/nats"
	"github.com/nats-io/nats.go"
)

func NewNatsSubscriber(ctx context.Context, nc *nats.Conn, svc Service) error {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "nats handler")
	{
		ra := kitnats.NewSubscriber(
			makeApiUserConfirmEndpoint(svc),
			decodeNATSUserConfirmRequestRequest,
			kitnats.EncodeJSONResponse,
			kitnats.SubscriberErrorLogger(logger),
		)
		if _, err := nc.QueueSubscribe("api-api-apiUserConfirm", "api", ra.ServeMsg(nc)); err != nil {
			return err
		}
	}
	{
		ra := kitnats.NewSubscriber(
			makeApiUserProfileEndpoint(svc),
			decodeNATSUserProfileRequestRequest,
			kitnats.EncodeJSONResponse,
			kitnats.SubscriberErrorLogger(logger),
		)
		if _, err := nc.QueueSubscribe("api-api-apiUserProfile", "api", ra.ServeMsg(nc)); err != nil {
			return err
		}
	}
	{
		ra := kitnats.NewSubscriber(
			makeApiUserRegistrationEndpoint(svc),
			decodeNATSUserRegRequestRequest,
			kitnats.EncodeJSONResponse,
			kitnats.SubscriberErrorLogger(logger),
		)
		if _, err := nc.QueueSubscribe("api-api-apiUserRegistration", "api", ra.ServeMsg(nc)); err != nil {
			return err
		}
	}

	return nil
}

func decodeNATSUserConfirmRequestRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request UserConfirmRequest
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeNATSUserProfileRequestRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request UserProfileRequest
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeNATSUserRegRequestRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request UserRegRequest
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}
	return request, nil
}
