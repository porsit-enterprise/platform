package valkey_test

import (
	"reflect"
	"testing"

	"github.com/porsit-enterprise/platform/foundation/configuration"
	. "github.com/porsit-enterprise/platform/infrastructure/Valkey"
	test_testing "github.com/porsit-enterprise/platform/test"
	"github.com/stretchr/testify/assert"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestConnect(t *testing.T) {
	config, err := configuration.Load(test_testing.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	c, err := Connect(config.Infrastructure.Valkey, test_testing.NAME)
	defer Close(c)

	assert.NoError(t, err)
	assert.Equal(t, "*valkey.singleClient", reflect.TypeOf(c).String())
}

func TestHealth(t *testing.T) {
	config, err := configuration.Load(test_testing.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	c, err := Connect(config.Infrastructure.Valkey, test_testing.NAME)
	defer Close(c)

	assert.NoError(t, err)

	err = Health(c, config.Infrastructure.Valkey)
	assert.NoError(t, err)
}
