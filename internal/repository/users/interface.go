package users

import "context"

type Repository interface {
	GetUserBySessionId(ctx context.Context, sessionId string) (*User, error)
	SetNewUser(ctx context.Context, phone string) error
	SetConfirmedNewUser(ctx context.Context, phone string) (*string, error)
}

