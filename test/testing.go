package test_testing

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/porsit-enterprise/platform/core"
	"github.com/porsit-enterprise/platform/foundation"
	found_conf "github.com/porsit-enterprise/platform/foundation/configuration"
	infra "github.com/porsit-enterprise/platform/infrastructure"
	infra_Centrifugo "github.com/porsit-enterprise/platform/infrastructure/Centrifugo"
	infra_Ollama "github.com/porsit-enterprise/platform/infrastructure/Ollama"
	infra_PostgreSQL "github.com/porsit-enterprise/platform/infrastructure/PostgreSQL"
	infra_Valkey "github.com/porsit-enterprise/platform/infrastructure/Valkey"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func init() {
	testing.Testing()

	if err := os.Setenv(core.ENVIRONMENT, core.ENVIRONMENT_TEST); err != nil {
		log.Fatal("can't set test envrionment variable")
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))

	ConfigPath = filepath.Join("..", "..")
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TurnOnLog(t *testing.T) {
	if t != nil {
		t.Helper()
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
}

func SetupFoundation(t *testing.T, path string) (foundation.Foundation, error) {
	if t != nil {
		t.Helper()
	}

	var err error

	cfg, err := found_conf.Load(path)
	if err != nil {
		return foundation.Foundation{}, err
	}

	foundation := foundation.Foundation{
		Settings:      nil,
		Configuration: cfg,
	}

	return foundation, nil
}

func SetupInfrastructure(t *testing.T, found foundation.Foundation) (infra.Infrastructure, error) {
	if t != nil {
		t.Helper()
	}

	var err error

	db, err := infra_PostgreSQL.Connect(found.Configuration.Infrastructure.PostgreSQL, nil)
	if err != nil {
		return infra.Infrastructure{}, err
	}

	valkey, err := infra_Valkey.Connect(found.Configuration.Infrastructure.Valkey, NAME)
	if err != nil {
		return infra.Infrastructure{}, err
	}

	ollama, err := infra_Ollama.Connect(found.Configuration.Infrastructure.Ollama)
	if err != nil {
		return infra.Infrastructure{}, err
	}

	centrifugo, err := infra_Centrifugo.Connect(found.Configuration.Infrastructure.Centrifugo)
	if err != nil {
		return infra.Infrastructure{}, err
	}

	return infra.Infrastructure{
		PostgreSQL: db,
		Valkey:     valkey,
		Ollama:     ollama,
		Centrifugo: centrifugo,
	}, nil
}

func CloseInfrastructure(t *testing.T, infra infra.Infrastructure) {
	if t != nil {
		t.Helper()
	}

	infra_PostgreSQL.Close(infra.PostgreSQL)
	infra_Valkey.Close(infra.Valkey)
}
