package cryptographic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/porsit-enterprise/platform/pkg/cryptographic"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestGenerateDigit(t *testing.T) {
	str, err := cryptographic.GenerateDigit(cryptographic.SMS_PASSWORD_LENGTH)
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
	assert.Len(t, str, cryptographic.SMS_PASSWORD_LENGTH)
}

func TestGenerateString(t *testing.T) {
	str := cryptographic.GenerateString(cryptographic.SMS_PASSWORD_LENGTH)
	assert.NotEmpty(t, str)
	assert.Len(t, str, cryptographic.SMS_PASSWORD_LENGTH)
}
