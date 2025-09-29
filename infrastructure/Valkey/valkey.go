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

func New(cfg cfg_entities.Valkey, name string) (valkey.Client, error) {
	slog.Info("connecting to Valkey")

	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{cfg.Connection},
		Username:    cfg.Username,
		Password:    cfg.Password,
		ClientName:  name,
	})
	if err != nil {
		return nil, fmt.Errorf("error in connect to Valkey: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ConnectionTimeout)*time.Second)
	defer cancel()

	res := client.Do(ctx, client.B().Ping().Build())
	if res.Error() != nil {
		return nil, fmt.Errorf("error in calling Valkey: %w", res.Error())
	}

	return client, nil
}

func Close(client valkey.Client) {
	slog.Info("closing Valkey connection")
	if client == nil {
		return
	}
	client.Close()
}
