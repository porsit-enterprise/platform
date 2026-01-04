package ollama

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	ollama_api "github.com/ollama/ollama/api"
	"github.com/ollama/ollama/envconfig"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Connect(config cfg_entities.Ollama) (*ollama_api.Client, error) {
	slog.Info("connecting to Ollama")

	defaultClient := &http.Client{}
	defaultClient.Timeout = time.Duration(config.ConnectionTimeout) * time.Second

	var mainUrl *url.URL

	if config.Connection == "" {
		mainUrl = envconfig.Host()
	} else {
		err := os.Setenv("OLLAMA_HOST", config.Connection)
		if err != nil {
			return nil, fmt.Errorf("unable to set Ollama environment variable: %w", err)
		}
		mainUrl = envconfig.Host()
	}

	client := ollama_api.NewClient(mainUrl, defaultClient)

	if !checkConnection(client, config) {
		return nil, fmt.Errorf("unable to connect to Ollama")
	}

	return client, nil
}

func Health(client *ollama_api.Client, config cfg_entities.Ollama) error {
	slog.Debug("check Ollama health")

	if client == nil {
		return nil
	}

	if !checkConnection(client, config) {
		return fmt.Errorf("unable to calling Ollama")
	}
	return nil
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func checkConnection(client *ollama_api.Client, config cfg_entities.Ollama) bool {
	for range max(1, config.ConnectionRetries) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnectionTimeout)*time.Second)

		err := client.Heartbeat(ctx)
		cancel()
		if err == nil {
			return true
		}

		slog.Warn("Ollama connection failed, retrying...", slog.Any("error", err))
		time.Sleep(time.Duration(config.ConnectionRetryDelay) * time.Second)
	}
	return false
}
