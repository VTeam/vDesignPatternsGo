package step4

const (
	DEFAULT_MAX_TOTAL = 8
	DEFAULT_MAX_IDLE  = 8
	DEFAULT_MIN_IDLE  = 0
)

type resourcePoolConfig struct {
	name     string
	maxTotal int32
	maxIdle  int32
	miniIdle int32
}

type S1 struct {
	maxIdle int32
}

func (s *S1) applyForOptions(r *resourcePoolConfig) {

	r.maxIdle = s.maxIdle

}

type S2 struct {
	maxTotal int32
}

func (s *S2) applyForOptions(r *resourcePoolConfig) {
	r.maxTotal = s.maxTotal
}

type S3 struct {
	miniIdle int32
}

func (s *S3) applyForOptions(r *resourcePoolConfig) {
	r.miniIdle = s.miniIdle
}

type ResourcePoolConfigOptions interface {
	applyForOptions(*resourcePoolConfig)
}

func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptions) *resourcePoolConfig {

	if name == "" {
		panic("name must not be empty")
	}

	config := &resourcePoolConfig{
		name: name,
	}
	for _, opt := range opts {
		opt.applyForOptions(config)
	}
	return config

}
