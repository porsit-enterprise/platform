package postgresql

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxvec "github.com/pgvector/pgvector-go/pgx"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Connect(config cfg_entities.PostgreSQL, schema []string) (*pgxpool.Pool, error) {
	slog.Info("connecting to PostgreSQL")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnectionTimeout)*time.Second)
	defer cancel()

	pgconfig, err := pgxpool.ParseConfig(config.Connection)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PostgreSQL configuration: %w", err)
	}

	if schema != nil {
		pgconfig.ConnConfig.RuntimeParams["search_path"] = strings.Join(schema, ",") + "," + PUBLIC_SCHEMA
	} else {
		pgconfig.ConnConfig.RuntimeParams["search_path"] = PUBLIC_SCHEMA
	}

	pgconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return pgxvec.RegisterTypes(ctx, conn)
	}

	connection, err := pgxpool.NewWithConfig(ctx, pgconfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create PostgreSQL connection: %w", err)
	}

	if !checkConnection(connection, config) {
		connection.Close()
		return nil, fmt.Errorf("unable to connect to PostgreSQL")
	}

	return connection, nil
}

func Health(connection *pgxpool.Pool, config cfg_entities.PostgreSQL) error {
	slog.Debug("check PostgreSQL health")

	if connection == nil {
		return nil
	}

	if !checkConnection(connection, config) {
		return fmt.Errorf("unable to calling PostgreSQL")
	}

	return nil
}

func Close(connection *pgxpool.Pool) {
	slog.Debug("closing PostgreSQL connection")

	if connection == nil {
		return
	}

	connection.Close()
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func checkConnection(connection *pgxpool.Pool, config cfg_entities.PostgreSQL) bool {
	for range max(1, config.ConnectionRetries) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnectionTimeout)*time.Second)

		err := connection.Ping(ctx)
		cancel()
		if err == nil {
			return true
		}

		slog.Warn("PostgreSQL connection failed, retrying...", slog.Any("error", err))
		time.Sleep(time.Duration(config.ConnectionRetryDelay) * time.Second)
	}
	return false
}
