package provider

import (
	"google.golang.org/grpc"

	sms "github.com/porsit-enterprise/platform/provider/SMS"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Provider struct {
	AI  *grpc.ClientConn
	SMS sms.SmsProvider
}
