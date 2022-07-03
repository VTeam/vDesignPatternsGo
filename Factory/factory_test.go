package factory

import "testing"

func TestIDCardFactory(*testing.T) {
	var factory = &IDCardFactory{
		products: make([]Product, 0), //golang 在声明结构体时，不会为结构体内的字段申请内存空间
	}

	card1 := factory.Create("小明")
	card1.Use()

	card2 := factory.Create("小红")
	card2.Use()
}

func TestLogFactory(t *testing.T) {

	var factory = &LogFactory{
		handlers: make([]Product, 0),
	}

	handler1 := factory.Create("file")
	handler1.Use()
	handler2 := factory.Create("stdout")
	handler2.Use()

}
