package step2

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

func (c *resourcePoolConfig) setMaxTotal(v int32) {
	c.maxTotal = v
}
func (c *resourcePoolConfig) setMaxIdle(v int32) {
	c.maxIdle = v
}
func (c *resourcePoolConfig) setMiniIdle(v int32) {
	c.miniIdle = v
}

func NewResourcePoolConfig(name string) *resourcePoolConfig {

	if name == "" {
		panic("name must not be empty")
	}
	return &resourcePoolConfig{
		name: name,
	}
}
