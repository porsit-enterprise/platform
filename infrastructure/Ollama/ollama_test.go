package ollama_test

import (
	"reflect"
	"testing"

	"github.com/porsit-enterprise/platform/foundation/configuration"
	. "github.com/porsit-enterprise/platform/infrastructure/Ollama"
	test_testing "github.com/porsit-enterprise/platform/test"
	"github.com/stretchr/testify/assert"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestConnect(t *testing.T) {
	config, err := configuration.Load(test_testing.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	c, err := Connect(config.Infrastructure.Ollama)

	assert.NoError(t, err)
	assert.Equal(t, "*api.Client", reflect.TypeOf(c).String())
}

func TestHealth(t *testing.T) {
	config, err := configuration.Load(test_testing.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	c, err := Connect(config.Infrastructure.Ollama)

	assert.NoError(t, err)

	err = Health(c, config.Infrastructure.Ollama)
	assert.NoError(t, err)
}
