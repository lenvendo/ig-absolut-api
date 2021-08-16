//go:generate easyjson -all endpoint.go
package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	_ "github.com/mailru/easyjson/gen"
)

//easyjson:json
type User struct {
	Id          string `json:"id,omitempty"`
	IsConfirmed bool   `json:"isConfirmed,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

//easyjson:json
type UserConfirmRequest struct {
	Code string `json:"code,omitempty"`
}

//easyjson:json
type UserConfirmResponse struct {
	Session string `json:"session,omitempty"`
}

//easyjson:json
type UserProfileRequest struct {
	Session string `json:"session,omitempty"`
}

//easyjson:json
type UserProfileResponse struct {
	User *User `json:"user,omitempty"`
}

//easyjson:json
type UserRegRequest struct {
	Phone string `json:"phone,omitempty"`
}

//easyjson:json
type UserRegResponse struct {
	Status string `json:"status,omitempty"`
}

//easyjson:skip
type endpoints struct {
	ApiUserConfirmEndpoint      endpoint.Endpoint
	ApiUserProfileEndpoint      endpoint.Endpoint
	ApiUserRegistrationEndpoint endpoint.Endpoint
}

func (e endpoints) ApiUserConfirm(ctx context.Context, req *UserConfirmRequest) (resp *UserConfirmResponse, err error) {
	response, err := e.ApiUserConfirmEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(UserConfirmResponse)
	return &r, err
}

func (e endpoints) ApiUserProfile(ctx context.Context, req *UserProfileRequest) (resp *UserProfileResponse, err error) {
	response, err := e.ApiUserProfileEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(UserProfileResponse)
	return &r, err
}

func (e endpoints) ApiUserRegistration(ctx context.Context, req *UserRegRequest) (resp *UserRegResponse, err error) {
	response, err := e.ApiUserRegistrationEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	r := response.(UserRegResponse)
	return &r, err
}

func makeApiUserConfirmEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserConfirmRequest)
		return s.ApiUserConfirm(ctx, &req)
	}
}

func makeApiUserProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserProfileRequest)
		return s.ApiUserProfile(ctx, &req)
	}
}

func makeApiUserRegistrationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRegRequest)
		return s.ApiUserRegistration(ctx, &req)
	}
}
