package embeddings

import (
	"context"
	"fmt"
	"time"

	ollama "github.com/ollama/ollama/api"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
	"github.com/porsit-enterprise/platform/infrastructure"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Generate(cfg cfg_entities.Ollama, infr infrastructure.Infrastructure, keyword string) ([]float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.OperationTimeout)*time.Second)
	defer cancel()

	kreq := &ollama.EmbeddingRequest{
		Model:  EMBEDDING_MODEL_NAME,
		Prompt: keyword,
	}
	kresp, err := infr.Ollama.Embeddings(ctx, kreq)
	if err != nil {
		return nil, fmt.Errorf("error in getting embedding: %w", err)
	}

	embedding := make([]float32, len(kresp.Embedding))
	for i, f := range kresp.Embedding {
		embedding[i] = float32(f)
	}

	return embedding, nil
}
