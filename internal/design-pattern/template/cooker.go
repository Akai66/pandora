// 模板模式，定义一个操作中算法的骨架，将一些不能共用的步逐延迟到子类实现，这种方法让子类在不改变算法结构的情况下，就能重新定义该算法的某些特定步逐
package template

import "fmt"

type cooker interface {
	fire()
	cook()
	outfire()
}

//类似一个抽象类，实现共用的方法
type cookMenu struct{}

func (cookMenu) fire() {
	fmt.Println("开火...")
}

func (cookMenu) outfire() {
	fmt.Println("关火...")
}

//子类实现cooke方法,将一些非共用方法延迟到子类实现
type FryTomato struct {
	cookMenu
}

func (*FryTomato) cook() {
	fmt.Println("炒西红柿...")
}

type FryEgg struct {
	cookMenu
}

func (*FryEgg) cook() {
	fmt.Println("炒鸡蛋...")
}

//定义算法骨架
func DoCook(c cooker) {
	c.fire()
	c.cook()
	c.outfire()
}
