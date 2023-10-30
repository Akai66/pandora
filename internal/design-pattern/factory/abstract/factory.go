// 抽象工厂
// 通过返回接口，在不公开内部实现的情况下，让调用者使用你提供的各种功能
package factory

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi,my name is %s\n", p.name)
}

func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}
