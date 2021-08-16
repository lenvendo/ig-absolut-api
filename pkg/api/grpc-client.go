package api

import (
	"context"
	"errors"

	pb "github.com/lenvendo/ig-absolut-api/internal/apipb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// NewGRPCClient returns an Service backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func NewGRPCClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) Service {
	// global client middlewares
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	}

	return endpoints{
		// Each individual endpoint is an grpc/transport.Client (which implements
		// endpoint.Endpoint) that gets wrapped with various middlewares. If you
		// made your own client library, you'd do this work there, so your server
		// could rely on a consistent set of client behavior.
		ApiUserConfirmEndpoint: grpctransport.NewClient(
			conn,
			"apipb.ApiService",
			"ApiUserConfirm",
			encodeGRPCUserConfirmRequest,
			decodeGRPCUserConfirmResponse,
			pb.UserConfirmResponse{},
			options...,
		).Endpoint(),
		ApiUserProfileEndpoint: grpctransport.NewClient(
			conn,
			"apipb.ApiService",
			"ApiUserProfile",
			encodeGRPCUserProfileRequest,
			decodeGRPCUserProfileResponse,
			pb.UserProfileResponse{},
			options...,
		).Endpoint(),
		ApiUserRegistrationEndpoint: grpctransport.NewClient(
			conn,
			"apipb.ApiService",
			"ApiUserRegistration",
			encodeGRPCUserRegRequest,
			decodeGRPCUserRegResponse,
			pb.UserRegResponse{},
			options...,
		).Endpoint(),
	}
}

func encodeGRPCUserConfirmRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*UserConfirmRequest)
	if !ok {
		return nil, errors.New("encodeGRPCUserConfirmRequest wrong request")
	}

	return UserConfirmRequestToPB(inReq), nil
}

func encodeGRPCUserProfileRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*UserProfileRequest)
	if !ok {
		return nil, errors.New("encodeGRPCUserProfileRequest wrong request")
	}

	return UserProfileRequestToPB(inReq), nil
}

func encodeGRPCUserRegRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*UserRegRequest)
	if !ok {
		return nil, errors.New("encodeGRPCUserRegRequest wrong request")
	}

	return UserRegRequestToPB(inReq), nil
}

func decodeGRPCUserConfirmResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*pb.UserConfirmResponse)
	if !ok {
		return nil, errors.New("decodeGRPCUserConfirmResponse wrong response")
	}

	resp := PBToUserConfirmResponse(inResp)

	return *resp, nil
}

func decodeGRPCUserProfileResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*pb.UserProfileResponse)
	if !ok {
		return nil, errors.New("decodeGRPCUserProfileResponse wrong response")
	}

	resp := PBToUserProfileResponse(inResp)

	return *resp, nil
}

func decodeGRPCUserRegResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*pb.UserRegResponse)
	if !ok {
		return nil, errors.New("decodeGRPCUserRegResponse wrong response")
	}

	resp := PBToUserRegResponse(inResp)

	return *resp, nil
}
