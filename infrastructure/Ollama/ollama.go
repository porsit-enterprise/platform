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
	defaultClient.Timeout = time.Duration(config.ConnectionTimeout)

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
	return client, nil
}

func Health(client *ollama_api.Client) error {
	slog.Debug("check Ollama health")
	if client == nil {
		return nil
	}

	_, err := client.Version(context.Background())
	if err != nil {
		return fmt.Errorf("unable to ping Ollama: %w", err)
	}
	return nil
}
