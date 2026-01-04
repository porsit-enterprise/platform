package cryptographic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/porsit-enterprise/platform/pkg/cryptographic"
	_ "github.com/porsit-enterprise/platform/test"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestGenerateDigit(t *testing.T) {
	str, err := GenerateDigit(SMS_PASSWORD_LENGTH)
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
	assert.Len(t, str, SMS_PASSWORD_LENGTH)
}

func TestGenerateString(t *testing.T) {
	str := GenerateString(SMS_PASSWORD_LENGTH)
	assert.NotEmpty(t, str)
	assert.Len(t, str, SMS_PASSWORD_LENGTH)
}
