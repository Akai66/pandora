package main

import (
	"fmt"

	"github.com/Akai66/pandora/internal/design-pattern/template"
)

func main() {
	// 炒西红柿
	template.DoCook(&template.FryTomato{})

	fmt.Println("炒下一道菜")

	//炒鸡蛋
	template.DoCook(&template.FryEgg{})
}
