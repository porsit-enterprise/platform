package sms

import (
	"errors"
	"net/http"
	"testing"

	"github.com/kavenegar/kavenegar-go"
	"github.com/stretchr/testify/assert"

	"github.com/porsit-enterprise/platform/foundation/configuration"
	test_testing "github.com/porsit-enterprise/platform/test"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Test_newKavenegar(t *testing.T) {
	config, err := configuration.Load(test_testing.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	api := newKavenegar(config.Provider.SMS.Kavenegar)
	assert.NotEmpty(t, api)
	assert.NotNil(t, api.client)

	account, err := api.client.Account.Info()
	if err != nil {
		var apiErr = new(kavenegar.APIError)
		if errors.As(err, &apiErr) {
			if apiErr.Status == 403 {
				t.Skip("skipping Test_newKavenegar due to invalid API key.")
			}
		}
		t.Fatal(err)
	}
	assert.NotEmpty(t, account)
	assert.NotEmpty(t, account.Remaincredit > 0)
	t.Logf("kavenegar account: %+v", account)
}

func TestKavenegarConectivity(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestKavenegarConectivity in short mode.")
	}

	resp, err := http.Get("https://api.kavenegar.com/v1/0/utils/getdate.json")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
