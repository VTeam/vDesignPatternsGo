package factory

import "testing"

// 优点: 扩展性好，符合了开闭原则，新增一种产品时，只需增加改对应的产品类和对应的工厂子类即可。
// 缺点: 当我们新增产品时，还需要提供对应的工厂类，系统中类的个数将会成倍增加，相当于增加了系统的复杂性。
func TestNewLogFactory(t *testing.T) {
	logHandler := NewLogFactory("file")
	if logHandler == nil {
		t.Error("创建日志失败")
		return
	}

	logHandler.Use()
}
