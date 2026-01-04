package centrifugo

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Connect(config cfg_entities.Centrifugo) (*grpc.ClientConn, error) {
	slog.Info("connect to Centrifugo")

	credauth := credentialsAuth{
		key: config.ApiKey,
	}

	connection, err := grpc.NewClient(
		config.ConnectionGRPC,
		grpc.WithDefaultServiceConfig(_RETRY_POLICY),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(credauth),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create Centrifugo connection: %w", err)
	}

	return connection, nil
}

func Health(connection *grpc.ClientConn, config cfg_entities.Centrifugo) error {
	slog.Debug("check Centrifugo health")

	if connection == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnectionTimeout)*time.Second)
	defer cancel()

	path, err := url.JoinPath(config.ConnectionHTTP, "/health")
	if err != nil {
		return fmt.Errorf("error in creating Centrifugo health endpoint URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("error in creating Centrifugo health request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("unable to calling Centrifugo: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed Centrifugo health check: status %d,", resp.StatusCode)
	}

	return nil
}

func Close(connection *grpc.ClientConn) error {
	slog.Debug("closing Centrifugo connection")

	if connection == nil {
		return nil
	}

	err := connection.Close()
	if err != nil {
		return fmt.Errorf("unable to close Centrifugo connection: %w", err)
	}

	return nil
}
