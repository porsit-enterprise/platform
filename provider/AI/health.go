package ai

import (
	"context"
	"time"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

var healthclient healthpb.HealthClient

//──────────────────────────────────────────────────────────────────────────────────────────────────

func SetupHealthCheckClient(com *grpc.ClientConn) {
	if healthclient != nil {
		return
	}

	healthclient = healthpb.NewHealthClient(com)
}

func CheckHealth(config cfg_entities.ProviderAI) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.RequestTimeout)*time.Second)
	defer cancel()

	response, err := healthclient.Check(ctx, &healthpb.HealthCheckRequest{})
	if err != nil {
		return false, err
	}

	return response.GetStatus() == healthpb.HealthCheckResponse_SERVING, nil
}
