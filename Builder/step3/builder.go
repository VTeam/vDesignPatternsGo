package step3

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

type Builder struct {
	name     string
	maxTotal int32
	maxIdle  int32
	miniIdle int32
}

func (b *Builder) setMaxTotal(v int32) *Builder {

	b.maxTotal = v
	return b
}

func (b *Builder) setMaxIdle(v int32) *Builder {

	b.maxIdle = v
	return b
}

func (b *Builder) setMiniIdle(v int32) *Builder {

	b.miniIdle = v
	return b
}

func (b *Builder) complete() *resourcePoolConfig {
	if b.name == "" {
		panic("name must not be empty")
	}
	if b.maxTotal == 0 {
		b.maxTotal = DEFAULT_MAX_TOTAL
	}
	if b.maxIdle == 0 {
		b.maxIdle = DEFAULT_MAX_IDLE
	}
	if b.miniIdle == 0 {
		b.miniIdle = DEFAULT_MIN_IDLE
	}
	return &resourcePoolConfig{
		name:     b.name,
		maxIdle:  b.maxIdle,
		maxTotal: b.maxTotal,
		miniIdle: b.miniIdle,
	}
}

func NewResourcePoolConfig(name string) *Builder {

	return &Builder{name: name}
}
