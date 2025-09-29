package postgresql_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/porsit-enterprise/platform/foundation/configuration"
	. "github.com/porsit-enterprise/platform/infrastructure/PostgreSQL"
	test_test "github.com/porsit-enterprise/platform/test"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestConnect(t *testing.T) {
	config, err := configuration.Load(test_test.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	db, err := Connect(config.Infrastructure.PostgreSQL, "")
	defer func() {
		_ = Close(db)
	}()

	assert.NoError(t, err)
	assert.Equal(t, "*pgxpool.Pool", reflect.TypeOf(db).String())
}
