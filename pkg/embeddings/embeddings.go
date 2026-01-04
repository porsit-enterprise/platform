package embeddings

import (
	"context"
	"fmt"
	"time"

	ollama "github.com/ollama/ollama/api"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Generate(cfg cfg_entities.Ollama, client *ollama.Client, keyword string) ([]float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.OperationTimeout)*time.Second)
	defer cancel()

	kreq := &ollama.EmbeddingRequest{
		Model:  EMBEDDING_MODEL_NAME,
		Prompt: keyword,
	}
	kresp, err := client.Embeddings(ctx, kreq)
	if err != nil {
		return nil, fmt.Errorf("error in generate embedding: %w", err)
	}

	embedding := make([]float32, len(kresp.Embedding))
	for i, f := range kresp.Embedding {
		embedding[i] = float32(f)
	}

	return embedding, nil
}
