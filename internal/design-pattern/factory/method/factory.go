// 工厂方法，可以通过返回工厂函数，将对象创建由一个工厂负责所有具体实例的创建，变为由一群子工厂负责对实例的创建
package factory

import "fmt"

type Product struct {
	Type string
	Name string
}

func (p Product) Desc() {
	fmt.Printf("产品名称:%s,产品类别:%s\n", p.Name, p.Type)
}

// NewProduct 返回工厂函数，再通过工厂函数，创建各种实例
func NewProductFactory(tp string) func(string) Product {
	return func(name string) Product {
		return Product{
			Type: tp,
			Name: name,
		}
	}
}
