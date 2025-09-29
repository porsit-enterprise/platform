package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"
	ollama "github.com/ollama/ollama/api"
	"github.com/valkey-io/valkey-go"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Infrastructure struct {
	PostgreSQL *pgxpool.Pool
	Valkey     valkey.Client
	Ollama     *ollama.Client
}
