package valkey_test

import (
	"reflect"
	"testing"

	"github.com/porsit-enterprise/platform/foundation/configuration"
	. "github.com/porsit-enterprise/platform/infrastructure/Valkey"
	test_test "github.com/porsit-enterprise/platform/test"
	"github.com/stretchr/testify/assert"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestNew(t *testing.T) {
	config, err := configuration.Load(test_test.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	c, err := New(config.Infrastructure.Valkey, "")
	defer Close(c)

	assert.NoError(t, err)
	assert.Equal(t, "*valkey.singleClient", reflect.TypeOf(c).String())
}
