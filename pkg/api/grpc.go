package api

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	pb "github.com/lenvendo/ig-absolut-api/internal/apipb"
	"github.com/lenvendo/ig-absolut-api/tools/logging"
	"github.com/lenvendo/ig-absolut-api/tools/tracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	apiUserConfirm      grpctransport.Handler
	apiUserProfile      grpctransport.Handler
	apiUserRegistration grpctransport.Handler
}

type ContextGRPCKey struct{}

type GRPCInfo struct{}

// NewGRPCServer makes a set of endpoints available as a gRPC apiServer.
func NewGRPCServer(ctx context.Context, s Service) pb.ApiServiceServer {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "grpc handler", "api")
	tracer := tracing.FromContext(ctx)

	options := []grpctransport.ServerOption{
		// grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(grpcToContext()),
		grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "grpc server", logger)),
		grpctransport.ServerFinalizer(closeGRPCTracer()),
	}

	return &grpcServer{
		apiUserConfirm: grpctransport.NewServer(
			makeApiUserConfirmEndpoint(s),
			decodeGRPCApiUserConfirmRequest,
			encodeGRPCApiUserConfirmResponse,
			options...,
		),
		apiUserProfile: grpctransport.NewServer(
			makeApiUserProfileEndpoint(s),
			decodeGRPCApiUserProfileRequest,
			encodeGRPCApiUserProfileResponse,
			options...,
		),
		apiUserRegistration: grpctransport.NewServer(
			makeApiUserRegistrationEndpoint(s),
			decodeGRPCApiUserRegistrationRequest,
			encodeGRPCApiUserRegistrationResponse,
			options...,
		),
	}
}

func JoinGRPC(ctx context.Context, s Service) func(*googlegrpc.Server) {
	return func(g *googlegrpc.Server) {
		pb.RegisterApiServiceServer(g, NewGRPCServer(ctx, s))
	}
}

func grpcToContext() grpc.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		return context.WithValue(ctx, ContextGRPCKey{}, GRPCInfo{})
	}
}

func closeGRPCTracer() grpc.ServerFinalizerFunc {
	return func(ctx context.Context, err error) {
		span := stdopentracing.SpanFromContext(ctx)
		span.Finish()
	}
}

func (s *grpcServer) ApiUserConfirm(ctx context.Context, req *pb.UserConfirmRequest) (*pb.UserConfirmResponse, error) {
	_, rep, err := s.apiUserConfirm.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserConfirmResponse), nil
}

func (s *grpcServer) ApiUserProfile(ctx context.Context, req *pb.UserProfileRequest) (*pb.UserProfileResponse, error) {
	_, rep, err := s.apiUserProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserProfileResponse), nil
}

func (s *grpcServer) ApiUserRegistration(ctx context.Context, req *pb.UserRegRequest) (*pb.UserRegResponse, error) {
	_, rep, err := s.apiUserRegistration.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserRegResponse), nil
}

func decodeGRPCApiUserConfirmRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.UserConfirmRequest)
	if !ok {
		return nil, errors.New("decodeGRPCApiUserConfirmRequest wrong request")
	}

	req := PBToUserConfirmRequest(inReq)
	return *req, nil
}

func decodeGRPCApiUserProfileRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.UserProfileRequest)
	if !ok {
		return nil, errors.New("decodeGRPCApiUserProfileRequest wrong request")
	}

	req := PBToUserProfileRequest(inReq)
	return *req, nil
}

func decodeGRPCApiUserRegistrationRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.UserRegRequest)
	if !ok {
		return nil, errors.New("decodeGRPCApiUserRegistrationRequest wrong request")
	}

	req := PBToUserRegRequest(inReq)
	return *req, nil
}

func encodeGRPCApiUserConfirmResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*UserConfirmResponse)
	if !ok {
		return nil, errors.New("encodeGRPCApiUserConfirmResponse wrong response")
	}

	return UserConfirmResponseToPB(inResp), nil
}

func encodeGRPCApiUserProfileResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*UserProfileResponse)
	if !ok {
		return nil, errors.New("encodeGRPCApiUserProfileResponse wrong response")
	}

	return UserProfileResponseToPB(inResp), nil
}

func encodeGRPCApiUserRegistrationResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*UserRegResponse)
	if !ok {
		return nil, errors.New("encodeGRPCApiUserRegistrationResponse wrong response")
	}

	return UserRegResponseToPB(inResp), nil
}

func UserToPB(d *User) *pb.User {
	if d == nil {
		return nil
	}

	resp := pb.User{
		Id:          d.Id,
		IsConfirmed: d.IsConfirmed,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}

	return &resp
}

func PBToUser(d *pb.User) *User {
	if d == nil {
		return nil
	}

	resp := User{
		Id:          d.Id,
		IsConfirmed: d.IsConfirmed,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}

	return &resp
}

func UserConfirmRequestToPB(d *UserConfirmRequest) *pb.UserConfirmRequest {
	if d == nil {
		return nil
	}

	resp := pb.UserConfirmRequest{
		Code: d.Code,
	}

	return &resp
}

func PBToUserConfirmRequest(d *pb.UserConfirmRequest) *UserConfirmRequest {
	if d == nil {
		return nil
	}

	resp := UserConfirmRequest{
		Code: d.Code,
	}

	return &resp
}

func UserConfirmResponseToPB(d *UserConfirmResponse) *pb.UserConfirmResponse {
	if d == nil {
		return nil
	}

	resp := pb.UserConfirmResponse{
		Session: d.Session,
	}

	return &resp
}

func PBToUserConfirmResponse(d *pb.UserConfirmResponse) *UserConfirmResponse {
	if d == nil {
		return nil
	}

	resp := UserConfirmResponse{
		Session: d.Session,
	}

	return &resp
}

func UserProfileRequestToPB(d *UserProfileRequest) *pb.UserProfileRequest {
	if d == nil {
		return nil
	}

	resp := pb.UserProfileRequest{
		Session: d.Session,
	}

	return &resp
}

func PBToUserProfileRequest(d *pb.UserProfileRequest) *UserProfileRequest {
	if d == nil {
		return nil
	}

	resp := UserProfileRequest{
		Session: d.Session,
	}

	return &resp
}

func UserProfileResponseToPB(d *UserProfileResponse) *pb.UserProfileResponse {
	if d == nil {
		return nil
	}

	resp := pb.UserProfileResponse{
		User: UserToPB(d.User),
	}

	return &resp
}

func PBToUserProfileResponse(d *pb.UserProfileResponse) *UserProfileResponse {
	if d == nil {
		return nil
	}

	resp := UserProfileResponse{
		User: PBToUser(d.User),
	}

	return &resp
}

func UserRegRequestToPB(d *UserRegRequest) *pb.UserRegRequest {
	if d == nil {
		return nil
	}

	resp := pb.UserRegRequest{
		Phone: d.Phone,
	}

	return &resp
}

func PBToUserRegRequest(d *pb.UserRegRequest) *UserRegRequest {
	if d == nil {
		return nil
	}

	resp := UserRegRequest{
		Phone: d.Phone,
	}

	return &resp
}

func UserRegResponseToPB(d *UserRegResponse) *pb.UserRegResponse {
	if d == nil {
		return nil
	}

	resp := pb.UserRegResponse{
		Status: d.Status,
	}

	return &resp
}

func PBToUserRegResponse(d *pb.UserRegResponse) *UserRegResponse {
	if d == nil {
		return nil
	}

	resp := UserRegResponse{
		Status: d.Status,
	}

	return &resp
}
