package step5

import "testing"

func TestNewResourcePoolConfig(t *testing.T) {

	config := NewResourcePoolConfig("dbpool", func(o *Options) {
		if o.maxIdle == 0 {
			o.maxIdle = 10
		}
		if o.maxTotal == 0 {
			o.maxTotal = 20
		}
		if o.miniIdle == 0 {
			o.miniIdle = 30
		}
	})

	t.Log(config)
}
