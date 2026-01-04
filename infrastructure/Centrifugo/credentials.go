package centrifugo

import "context"

//──────────────────────────────────────────────────────────────────────────────────────────────────

type credentialsAuth struct {
	key string
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func (t credentialsAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "apikey " + t.key,
	}, nil
}

func (t credentialsAuth) RequireTransportSecurity() bool {
	return false
}
