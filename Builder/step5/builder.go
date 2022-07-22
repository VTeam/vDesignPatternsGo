package step5

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
	name     string
	maxTotal int32
	maxIdle  int32
	miniIdle int32
}

type Option func(*Options)

func NewResourcePoolConfig(name string, opts ...Option) *resourcePoolConfig {

	if name == "" {
		panic("name must not be empty")
	}
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	config := &resourcePoolConfig{
		name:     name,
		maxTotal: options.maxTotal,
		maxIdle:  options.maxIdle,
		miniIdle: options.miniIdle,
	}

	return config

}
