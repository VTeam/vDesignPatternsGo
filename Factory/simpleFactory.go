package factory

import "fmt"

type LogFactorySimple interface {
	Use()
}

type LogFile struct{}

func (l *LogFile) Use() {
	fmt.Println("Use LogFile handle log")
}

type LogStdOut struct{}

func (l *LogStdOut) Use() {
	fmt.Println("Use LogStdOut handle log")
}

type LogMQ struct{}

func (l *LogMQ) Use() {
	fmt.Println("Use LogMQ handle log")

}

// 日志工厂，根据输入字符串实例化对应的日志类
// 优点： 简单工厂模式可以根据需求，动态生成使用者所需类的对象，而使用者不用去知道怎么创建对象，使得各个模块各司其职，降低了系统的耦合性
// 缺点： 扩展性差，违背了开闭原则,新增日志类型时，需要修改工厂方法代码。
func NewLogFactory(name string) LogFactorySimple {
	switch name {
	case "file":
		return &LogFile{}
	case "stdout":
		return &LogStdOut{}
	case "mq":
		return &LogMQ{}
	default:
		return nil
	}

}
