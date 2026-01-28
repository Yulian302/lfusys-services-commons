package health

type ReadinessCheck interface {
	Name() string
	IsReady() bool
}
