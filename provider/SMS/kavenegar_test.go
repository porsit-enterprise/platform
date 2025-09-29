package sms

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/porsit-enterprise/platform/foundation/configuration"
	test_test "github.com/porsit-enterprise/platform/test"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Test_newKavenegar(t *testing.T) {
	config, err := configuration.Load(test_test.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	api := newKavenegar(config.Provider.SMS.Kavenegar)
	assert.NotEmpty(t, api)
	assert.NotNil(t, api.client)

	accaount, err := api.client.Account.Info()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, accaount)
	assert.NotEmpty(t, accaount.Remaincredit > 0)
	t.Logf("kavenegar account: %+v", accaount)
}

func TestKavenegarConectivity(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestKavenegarConectivity in short mode.")
	}

	resp, err := http.Get("https://api.kavenegar.com/v1/0/utils/getdate.json")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
