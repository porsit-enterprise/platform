package provider

import (
	"google.golang.org/grpc"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Provider struct {
	AI *grpc.ClientConn
}
