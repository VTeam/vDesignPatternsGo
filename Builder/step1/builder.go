package step1

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

type resourcePoolConfigOptions struct {
	maxTotal int32
	maxIdle  int32
	miniIdle int32
}

func NewResourcePoolConfig(name string, opts resourcePoolConfigOptions) *resourcePoolConfig {
	if name == "" {
		panic("name cannot be empty")
	}
	if opts.maxTotal == 0 {
		opts.maxTotal = DEFAULT_MAX_TOTAL
	}
	if opts.maxIdle == 0 {
		opts.maxIdle = DEFAULT_MAX_IDLE
	}
	if opts.miniIdle == 0 {
		opts.miniIdle = DEFAULT_MIN_IDLE
	}

	return &resourcePoolConfig{
		name:     name,
		maxTotal: opts.maxTotal,
		miniIdle: opts.miniIdle,
		maxIdle:  opts.maxIdle,
	}
}
