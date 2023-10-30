// 验证工厂模式
package main

import (
	"fmt"

	abstractfactory "github.com/Akai66/pandora/internal/design-pattern/factory/abstract"
	methodfactory "github.com/Akai66/pandora/internal/design-pattern/factory/method"
	simplefactory "github.com/Akai66/pandora/internal/design-pattern/factory/simple"
)

func main() {
	// 简单工厂
	p1 := simplefactory.NewPerson("jack", 19)
	fmt.Printf("p1:%T\n", p1)
	p1.Greet()

	//抽象工厂
	p2 := abstractfactory.NewPerson("rose", 20)
	fmt.Printf("p2:%T\n", p2)
	p2.Greet()

	//工厂方法
	//返回一个创建日用品实例的工厂方法
	newDailyGoods := methodfactory.NewProductFactory("日用品")
	//通过返回的子工厂，日用品工厂，创建具体的日用品实例
	pd1 := newDailyGoods("肥皂")
	pd1.Desc()
	pd2 := newDailyGoods("牙刷")
	pd2.Desc()

	//返回一个创建体育用品实例的工厂方法
	newSportGoods := methodfactory.NewProductFactory("体育用品")
	//通过返回的子工厂，体育用品工厂，创建具体的体育用品实例
	pd3 := newSportGoods("篮球")
	pd4 := newSportGoods("乒乓球")

	pd3.Desc()
	pd4.Desc()

}
