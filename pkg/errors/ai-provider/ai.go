package aiprovider

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// CheckError checks the error returned from gRPC calls and returns a more user-friendly error message.
// It returns the error and a boolean indicating whether the error is known or not.
func CheckError(err error, ctx context.Context) (error, bool) {
	s := status.Convert(err)
	c := s.Code()
	m := s.Message()

	if c == codes.OK {
		return nil, true
	}

	if err != nil {
		select {
		case <-ctx.Done():
			return ErrTimeout, true
		default:
		}
	}

	if c == codes.Canceled {
		return ErrTimeout, true
	}
	if c == codes.DeadlineExceeded {
		return ErrTimeout, true
	}

	if c == codes.Unavailable && strings.HasPrefix(m, "connection error:") {
		return ErrConnection, true
	}

	if strings.HasPrefix(m, "1011:") {
		return ErrRateLimit, true
	}
	if strings.HasPrefix(m, "1021:") {
		return ErrAuthentication, true
	}
	if strings.HasPrefix(m, "1031:") {
		return ErrTimeout, true
	}
	if strings.HasPrefix(m, "2003:") {
		return ErrLoading, true
	}

	return errors.New(m), false
}
