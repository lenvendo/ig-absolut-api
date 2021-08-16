package token

import (
	"context"
	"github.com/segmentio/ksuid"

	"github.com/lenvendo/ig-absolut-api/internal/db"
	"github.com/pkg/errors"
)

var (
	ConnError = errors.New("get connection error")
)

type repository struct {
	dbConn *db.Connection
}

func NewRepository(ctx context.Context, connection *db.Connection) Repository {
	return &repository{
		dbConn: connection,
	}
}

func (r *repository) SetTokenByUserId(ctx context.Context, id *string) (*string, error) {
	conn, err := db.GetMasterConn(ctx, r.dbConn)
	if err != nil {
		return nil, errors.Wrap(ConnError, err.Error())
	}
	defer conn.Release()

	row, err := conn.Query(ctx, `INSERT INTO token(user_id, token) VALUES($1, $2) RETURNING token`, id, ksuid.New().String())
	if err != nil {
		return nil, errors.Wrap(err, "insert error")
	}
	defer row.Close()
	var token string
	for row.Next() {
		row.Scan(&token)
	}

	if row.Err() != nil {
		return nil, errors.Wrap(row.Err(), "parse token")
	}

	return &token, nil
}
