package health

import "context"

type ReadinessCheck interface {
	Name() string
	IsReady(ctx context.Context) error
}
