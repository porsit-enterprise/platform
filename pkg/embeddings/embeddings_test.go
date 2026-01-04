package embeddings_test

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	infra_Ollama "github.com/porsit-enterprise/platform/infrastructure/Ollama"
	. "github.com/porsit-enterprise/platform/pkg/embeddings"
	test_testing "github.com/porsit-enterprise/platform/test"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

var testFile string

func init() {
	testFile = filepath.Join("..", "..", test_testing.TEST_DATA_PATH, "embedding")

	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		panic(err)
	}
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestGenerate(t *testing.T) {
	found, err := test_testing.SetupFoundation(t, test_testing.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	client, err := infra_Ollama.Connect(found.Configuration.Infrastructure.Ollama)
	if err != nil {
		t.Fatal(err)
	}

	result, err := Generate(found.Configuration.Infrastructure.Ollama, client, "porsit")
	assert.NoError(t, err)
	assert.NotZero(t, len(result))

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	content := strings.TrimSpace(string(data))
	if content == "" {
		t.Fatal("embedding test data is empty")
	}

	parts := strings.Split(content, "\n")

	expected := make([]float32, len(parts))

	for i, part := range parts {
		valueStr := strings.TrimSpace(part)
		if valueStr == "" {
			t.Fatal("empty value at position", i)
		}

		value64, err := strconv.ParseFloat(valueStr, 32)
		if err != nil {
			t.Fatalf("failed to parse value '%s' at position %d: %v", valueStr, i, err)
		}

		expected[i] = float32(value64)
	}

	assert.Equal(t, expected, result)
}
