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

type Options struct {
	maxTotal int32
	maxIdle  int32
	miniIdle int32
}

func setOptionsDefaults(options Options) Options {

	if options.maxTotal == 0 {
		options.maxTotal = DEFAULT_MAX_TOTAL
	}
	if options.maxIdle == 0 {
		options.maxIdle = DEFAULT_MAX_IDLE
	}
	if options.miniIdle == 0 {
		options.miniIdle = DEFAULT_MIN_IDLE
	}
	return options
}

func NewResourcePoolConfig(name string, options Options) *resourcePoolConfig {

	if name == "" {
		panic("name cannot be empty")
	}

	options = setOptionsDefaults(options)

	return &resourcePoolConfig{
		name:     name,
		maxTotal: options.maxTotal,
		miniIdle: options.miniIdle,
		maxIdle:  options.maxIdle,
	}
}
