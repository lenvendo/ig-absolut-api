//go:generate mockgen -destination service_mock.go -package api  github.com/lenvendo/ig-absolut-api/pkg/api Service
package api

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
)

type Service interface {
	ApiUserConfirm(context.Context, *UserConfirmRequest) (*UserConfirmResponse, error)
	ApiUserProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error)
	ApiUserRegistration(context.Context, *UserRegRequest) (*UserRegResponse, error)
}
