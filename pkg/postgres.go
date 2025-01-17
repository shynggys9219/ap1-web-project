package pkg

import (
	"context"
	"fmt"
	pg "github.com/jackc/pgx/v5"
)

type Config struct {
	Database string `env:"POSTGRESQL_DB"`
	Host     string `env:"POSTGRESQL_URI"`
	Port     uint16 `env:"POSTGRESQL_PORT"`
	Username string `env:"POSTGRESQL_USERNAME"`
	Password string `env:"POSTGRESQL_PASSWORD"`
}

type DB struct {
	Conn *pg.Conn
}

func NewDB(ctx context.Context, cfg Config) (*DB, error) {
	opts, err := pg.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("pg.ParseConfig: %w", err)
	}

	opts.User = cfg.Username
	opts.Password = cfg.Password
	opts.Database = cfg.Database
	opts.Host = cfg.Host
	opts.Port = cfg.Port
	db, err := pg.ConnectConfig(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &DB{
		Conn: db,
	}, nil
}
