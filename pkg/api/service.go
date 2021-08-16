package api

import (
	"context"
	"github.com/lenvendo/ig-absolut-api/internal/repository/token"
	"github.com/lenvendo/ig-absolut-api/internal/repository/users"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type apiService struct {
	userRepo  users.Repository
	tokenRepo token.Repository
	nats      *nats.Conn
}

func NewApiService(userRepo users.Repository, tokenRepo token.Repository, nats *nats.Conn) Service {
	return &apiService{userRepo, tokenRepo, nats}
}

func (s *apiService) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	a := UserConfirmResponse{}
	if err != nil {
		return &a, err
	}
	/**
	ToDo
	сообщаем сервису, по grpc, FakeSms полученный код, возврщаем телефон
	*/
	phone := string("phone")

	id, err := s.userRepo.SetConfirmedNewUser(ctx, phone)
	if err != nil {
		return &a, errors.Wrap(err, "set confirmed users")
	}
	tokenId, err := s.tokenRepo.SetTokenByUserId(ctx, id)
	if err != nil {
		return &a, errors.Wrap(err, "create token")
	}

	a.Session = *tokenId
	return &a, nil
}

func (s *apiService) ApiUserProfile(ctx context.Context, req *UserProfileRequest) (resp *UserProfileResponse, err error) {
	a := UserProfileResponse{}
	if err != nil {
		return &a, err
	}
	user, err := s.userRepo.GetUserBySessionId(ctx, req.Session)
	a.User = UnMarshallUser(user)
	return &a, nil
}

func (s *apiService) ApiUserRegistration(ctx context.Context, req *UserRegRequest) (resp *UserRegResponse, err error) {
	if req == nil {
		return
	}
	if err = s.userRepo.SetNewUser(ctx, req.Phone); err != nil {
		return nil, errors.Wrap(err, "set new user")
	}
	if err := s.nats.Publish("phone", []byte(req.Phone)); err != nil {
		return nil, errors.Wrap(err, "publish phone in queue")
	}

	return &UserRegResponse{Status: "Success"}, nil
}

func UnMarshallUser(user *users.User) *User {
	return &User{
		Id:          user.Id,
		IsConfirmed: user.IsConfirmed,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
