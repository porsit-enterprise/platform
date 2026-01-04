package ai

import (
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

const _RETRY_POLICY = `{
    "methodConfig": [{
        "name": [{"service": "ModelService"},{"service": "grpc.health.v1.Health"}],
        "waitForReady": true,
        "retryPolicy": {
            "maxAttempts": 4,
            "initialBackoff": "0.1s",
            "maxBackoff": "1s",
            "backoffMultiplier": 2.0,
            "retryableStatusCodes": ["UNAVAILABLE", "RESOURCE_EXHAUSTED"]
        }
    }]
}`

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Dial(config cfg_entities.ProviderAI) (*grpc.ClientConn, error) {
	slog.Info("connecting to AI provider")

	connection, err := grpc.NewClient(
		config.Connection,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(_RETRY_POLICY),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create AI provider connection: %w", err)
	}

	return connection, nil
}

func Close(connection *grpc.ClientConn) error {
	slog.Debug("closing AI connection")

	if connection == nil {
		return nil
	}
	err := connection.Close()
	if err != nil {
		return fmt.Errorf("unable to close AI provider connection: %w", err)
	}
	return nil
}
