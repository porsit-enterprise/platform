package json

import (
	"bytes"
	"encoding/json"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Compactor(input []byte) []byte {
	var buf bytes.Buffer
	if json.Compact(&buf, input) != nil {
		return input
	}
	return buf.Bytes()
}
