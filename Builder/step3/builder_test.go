package step3

import "testing"

func TestNewResourcePoolConfig(t *testing.T) {
	var config *resourcePoolConfig = NewResourcePoolConfig("dbconnectionpool").
		setMiniIdle(10).
		setMaxIdle(20).
		setMaxTotal(30).
		complete()

	t.Log(config)
}
