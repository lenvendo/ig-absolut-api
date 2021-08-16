package api

import (
	"context"
)

type apiService struct {
}

func NewApiService() Service {
	return &apiService{}
}

func (s *apiService) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	a := UserConfirmResponse{}
	if err != nil {
		return &a, err
	}
	return &a, nil
}

func (s *apiService) ApiUserProfile(ctx context.Context, req *UserProfileRequest) (resp *UserProfileResponse, err error) {
	a := UserProfileResponse{}
	if err != nil {
		return &a, err
	}
	return &a, nil
}

func (s *apiService) ApiUserRegistration(ctx context.Context, req *UserRegRequest) (resp *UserRegResponse, err error) {
	a := UserRegResponse{}
	if err != nil {
		return &a, err
	}
	return &a, nil
}
