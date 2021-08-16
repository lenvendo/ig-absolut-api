package api

import (
	"context"
	"encoding/json"

	kitnats "github.com/go-kit/kit/transport/nats"
	"github.com/nats-io/nats.go"
)

func NewNatsPublisher(ctx context.Context, nc *nats.Conn) Service {
	return endpoints{
		ApiUserConfirmEndpoint: kitnats.NewPublisher(
			nc,
			"api-api-apiUserConfirm",
			kitnats.EncodeJSONRequest,
			decodeHTTPApiUserConfirm,
		).Endpoint(),
		ApiUserProfileEndpoint: kitnats.NewPublisher(
			nc,
			"api-api-apiUserProfile",
			kitnats.EncodeJSONRequest,
			decodeHTTPApiUserProfile,
		).Endpoint(),
		ApiUserRegistrationEndpoint: kitnats.NewPublisher(
			nc,
			"api-api-apiUserRegistration",
			kitnats.EncodeJSONRequest,
			decodeHTTPApiUserRegistration,
		).Endpoint(),
	}
}

func decodeHTTPApiUserConfirm(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var response UserConfirmResponse
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeHTTPApiUserProfile(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var response UserProfileResponse
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeHTTPApiUserRegistration(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var response UserRegResponse
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}
	return response, nil
}
