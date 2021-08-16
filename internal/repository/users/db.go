package users

import (
	"context"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/lenvendo/ig-absolut-api/internal/db"
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

func (r *repository) GetUserBySessionId(ctx context.Context, sessionId string) (*User, error) {
	conn, err := db.GetReplicaConn(ctx, r.dbConn)
	if err != nil {
		return nil, errors.Wrap(ConnError, err.Error())
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT * FROM users WHERE id = (SELECT user_id FROM token WHERE token = $1)`, sessionId)
	if err != nil {
		return nil, errors.Wrap(err, "select from users")
	}
	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.IsConfirmed, &user.CreatedAt, &user.UpdatedAt)
	}
	if rows.Err() != nil {
		return nil, errors.Wrap(err, "parse users")
	}

	return &user, nil
}

func (r *repository) SetNewUser(ctx context.Context, phone string) (err error) {
	conn, err := db.GetMasterConn(ctx, r.dbConn)
	if err != nil {
		return errors.Wrap(ConnError, err.Error())
	}
	defer conn.Release()

	id, err := uuid.FromBytes([]byte(phone))
	if err != nil {
		return errors.Wrap(err, "convert from phone to uuid")
	}

	_, err = conn.Query(ctx, `INSERT INTO users(id, created_at) VALUES($1, now())`, id.String())
	if err != nil {
		return errors.Wrap(err, "insert error")
	}

	return nil
}

func (r *repository) SetConfirmedNewUser(ctx context.Context, phone string) (*string, error) {
	conn, err := db.GetMasterConn(ctx, r.dbConn)
	if err != nil {
		return nil, errors.Wrap(ConnError, err.Error())
	}
	defer conn.Release()

	id, err := uuid.FromBytes([]byte(phone))
	if err != nil {
		return nil, errors.Wrap(err, "convert from phone to uuid")
	}

	row, err := conn.Query(ctx, `UPDATE users SET is_confirmed = true, updated_at = now() WHERE id = $1 RETURNING id`, id.String())
	if err != nil {
		return nil, errors.Wrap(err, "insert error")
	}
	defer row.Close()

	var currentId string
	for row.Next() {
		row.Scan(&currentId)
	}
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &currentId, nil
}
