package cryptographic

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func GenerateDigit(length int) (string, error) {
	otp := make([]byte, length)
	for i := range length {
		num, err := crand.Int(crand.Reader, big.NewInt(int64(len(_DIGIT_TABLE))))
		if err != nil {
			return "", err
		}
		otp[i] = _DIGIT_TABLE[num.Int64()]
	}
	return string(otp), nil
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// GenerateString generate a random string
//
// Base on "https://stackoverflow.com/a/31832326"
func GenerateString(length int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(length)

	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(_ALPHA_TABLE) {
			sb.WriteByte(_ALPHA_TABLE[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

//──────────────────────────────────────────────────────────────────────────────────────────────────
