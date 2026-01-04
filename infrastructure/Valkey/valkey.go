package valkey

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/valkey-io/valkey-go"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Connect(cfg cfg_entities.Valkey, name string) (valkey.Client, error) {
	slog.Info("connecting to Valkey")

	for range max(1, cfg.ConnectionRetries) {
		client, err := valkey.NewClient(valkey.ClientOption{
			InitAddress: []string{cfg.Connection},
			Username:    cfg.Username,
			Password:    cfg.Password,
			ClientName:  name,
		})
		if err == nil {
			return client, nil
		}

		slog.Warn("error in connect to Valkey, retrying...", "error", err)
		time.Sleep(time.Duration(cfg.ConnectionRetryDelay) * time.Second)
	}

	return nil, fmt.Errorf("error in connect to Valkey")
}

func Health(client valkey.Client, config cfg_entities.Valkey) error {
	slog.Debug("check Valkey health")

	if client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnectionTimeout)*time.Second)
	defer cancel()

	res := client.Do(ctx, client.B().Ping().Build())
	if res.Error() != nil {
		return fmt.Errorf("error in calling Valkey: %w", res.Error())
	}
	return nil
}

func Close(client valkey.Client) {
	slog.Debug("closing Valkey connection")

	if client == nil {
		return
	}

	client.Close()
}
