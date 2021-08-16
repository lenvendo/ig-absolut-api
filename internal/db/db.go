package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/lenvendo/ig-absolut-api/configs"
)

const TimeStampLayout = "2006-01-02 15:04:05"
const ShowTimeStampLayout = "02.01.2006 15:04"
const TransactionCount int = 200

type Connection struct {
	Master  *pgxpool.Pool
	Replica *pgxpool.Pool
}

func Connect(ctx context.Context, cfg *configs.Config) (*Connection, error) {
	var res Connection
	var err error

	// Подключаемся к мастеру
	res.Master, err = conn(ctx, cfg.Postgres.Master)
	if err != nil {
		return nil, errors.Wrap(err, "Master DB connect")
	}

	// Подключаемся к реплике
	res.Replica, err = conn(ctx, cfg.Postgres.Replica)
	if err != nil {
		return nil, errors.Wrap(err, "Replica DB connect")
	}

	return &res, nil
}

// Подключаемся к базе
func conn(ctx context.Context, db configs.Database) (c *pgxpool.Pool, err error) {
	return pgxpool.Connect(ctx, fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DatabaseName,
		db.Secure,
	))
}

func Close(ctx context.Context, c *Connection) {
	c.Master.Close()
	c.Replica.Close()
}

func GetMasterConn(ctx context.Context, c *Connection) (*pgxpool.Conn, error) {
	conn, err := c.Master.Acquire(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get master connection")
	}
	return conn, nil
}

func GetReplicaConn(ctx context.Context, c *Connection) (*pgxpool.Conn, error) {
	conn, err := c.Replica.Acquire(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get replica connection")
	}
	return conn, nil
}

func Ping(ctx context.Context, c *Connection) error {
	conn, err := c.Master.Acquire(ctx)
	if err != nil {
		return errors.Wrap(err, "master ping error")
	}
	conn.Release()

	conn, err = c.Replica.Acquire(ctx)
	if err != nil {
		return errors.Wrap(err, "replica ping error")
	}
	conn.Release()
	return nil
}

func GetConstraintName(ctx context.Context, conn *pgxpool.Conn, tableName string, columnNames []string) string {
	empty := ""
	if len(columnNames) == 0 {
		return empty
	}

	columnNamesQuoted := make([]string, len(columnNames))
	for i, v := range columnNames {
		columnNamesQuoted[i] = fmt.Sprintf("'%v'", v)
	}

	q := `SELECT kcu1.constraint_name
		  FROM information_schema.key_column_usage as kcu1
		  WHERE kcu1.table_name = '%v' AND kcu1.constraint_name IN (
				SELECT DISTINCT kcu.constraint_name
				FROM information_schema.key_column_usage as kcu
				WHERE kcu.table_name = '%v' AND kcu.column_name IN (%v)
		  )
		  GROUP BY kcu1.constraint_name
		  HAVING COUNT(kcu1.constraint_name) = %v`

	query := fmt.Sprintf(q, tableName, tableName, strings.Join(columnNamesQuoted, ", "), len(columnNamesQuoted))

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return empty
	}
	defer rows.Close()

	for rows.Next() {
		constraintName := ""

		err := rows.Scan(&constraintName)
		if err != nil {
			return empty
		}

		return constraintName
	}

	return empty
}