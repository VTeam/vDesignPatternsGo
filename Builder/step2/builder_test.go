package step2

import "testing"

func TestNewResourcePoolConfig(*testing.T) {
	config := NewResourcePoolConfig("dbconnecionpool")
	config.setMaxIdle(10)
	config.setMaxTotal(20)
	config.setMiniIdle(30)
}
