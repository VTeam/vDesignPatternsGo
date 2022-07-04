package factory

import "testing"

func TestNewIphone(t *testing.T) {
	logHandler := NewLogFactory("file")
	if logHandler == nil {
		t.Error("创建日志失败")
		return
	}

	logHandler.Use()
}
