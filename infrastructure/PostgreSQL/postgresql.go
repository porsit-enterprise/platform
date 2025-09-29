package postgresql

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxvec "github.com/pgvector/pgvector-go/pgx"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Connect(config cfg_entities.PostgreSQL, schema string) (*pgxpool.Pool, error) {
	slog.Info("connecting to PostgreSQL")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnectionTimeout)*time.Second)
	defer cancel()

	pgconfig, err := pgxpool.ParseConfig(config.Connection)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PostgreSQL configuration: %w", err)
	}

	if schema != "" {
		pgconfig.ConnConfig.RuntimeParams["search_path"] = schema
	}

	pgconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return pgxvec.RegisterTypes(ctx, conn)
	}

	connection, err := pgxpool.NewWithConfig(ctx, pgconfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create PostgreSQL connection: %w", err)
	}

	err = connection.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping PostgreSQL: %w", err)
	}

	return connection, nil
}

func Close(connection *pgxpool.Pool) {
	slog.Debug("closing PostgreSQL connection")

	if connection == nil {
		return
	}
	connection.Close()
}
