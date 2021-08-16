package api

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/lenvendo/ig-absolut-api/internal/repository/token"
	"github.com/lenvendo/ig-absolut-api/internal/repository/users"
	"github.com/lenvendo/ig-absolut-api/internal/verification"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"math/big"
	"strconv"
)

type apiService struct {
	userRepo  users.Repository
	tokenRepo token.Repository
	verify    verification.MemoryService
	nats      *nats.Conn
}

func NewApiService(userRepo users.Repository, tokenRepo token.Repository, verify verification.MemoryService, nats *nats.Conn) Service {
	return &apiService{userRepo, tokenRepo, verify, nats}
}

func (s *apiService) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	a := UserConfirmResponse{}
	if err != nil {
		return &a, err
	}

	code, err := strconv.Atoi(req.Code)
	if err != nil {
		return nil, errors.Wrap(err, "not valid code")
	}
	phone, err := s.verify.Verify(uint8(code))
	if err != nil {
		return nil, errors.Wrap(err, "phone is not valid")
	}

	id, err := s.userRepo.SetConfirmedNewUser(ctx, *phone)
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
	r, err := rand.Int(rand.Reader, big.NewInt(9999))
	if err != nil {
		return nil, err
	}
	if err := s.verify.SetPhoneAndCode(req.Phone, uint8(r.Int64())); err != nil {
		return nil, errors.Wrap(err, "set phone in verification service layer")
	}

	if err := s.nats.Publish("tasks", []byte(fmt.Sprintf("phone:%s;code:%v", req.Phone, r.Int64()))); err != nil {
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
