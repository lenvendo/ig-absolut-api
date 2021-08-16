//go:generate mockgen -destination service_mock.go -package health  github.com/lenvendo/ig-absolut-api/pkg/health Service
package health

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
)

type Service interface {
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}
