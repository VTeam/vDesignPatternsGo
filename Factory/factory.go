package factory

import "fmt"

// 工厂接口
type Factory interface {
	// 对外公开的构建实例的方法
	Create(owner string) Product

	// 实际生成实例的方法
	createProduct(owner string) Product

	// 保存已生成的实例
	registerProduct(product Product)
}

// 产品接口
type Product interface {
	// 获取产品持有人名字
	GetOwner() string

	// 使用
	Use()
}

// 身份证类型，实现了Product接口
type IDCard struct {
	owner string `json:"owner,omitempty"`
}

func (c *IDCard) GetOwner() string {
	return c.owner
}

func (c *IDCard) Use() {
	fmt.Printf("使用了%s的身份证抢火车票!\n", c.owner)
}

// 身份证工厂，实现了Factory接口
type IDCardFactory struct {
	products []Product `json:"products,omitempty"`
}

func (f *IDCardFactory) Create(owner string) Product {
	idCard := f.createProduct(owner)
	f.registerProduct(idCard)
	return idCard
}

func (f *IDCardFactory) createProduct(owner string) Product {
	return &IDCard{owner: owner}
}

func (f *IDCardFactory) registerProduct(idCard Product) {
	f.products = append(f.products, idCard)
}

type LogFileHandler struct {
	name string
}

func (lf *LogFileHandler) GetOwner() string {

	return lf.name
}
func (lf *LogFileHandler) Use() {

}

// type LogStdOutHandler struct{}

// type LogMQHandler struct{}

type LogFactory struct {
	handlers []Product
}

func (t *LogFactory) Create(name string) Product {
	handler := t.createProduct(name)
	t.registerProduct(handler)
	return handler
}

func (t *LogFactory) createProduct(name string) Product {
	return &LogFileHandler{name: name}
}

func (t *LogFactory) registerProduct(h Product) {
	t.handlers = append(t.handlers, h)

}
