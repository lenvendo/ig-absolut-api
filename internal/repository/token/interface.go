package token

import "context"

type Repository interface {
	SetTokenByUserId(ctx context.Context, id *string) (*string, error)
}
