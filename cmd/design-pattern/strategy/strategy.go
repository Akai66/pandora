package main

import (
	"fmt"

	"github.com/Akai66/pandora/internal/design-pattern/strategy"
)

func main() {
	op := &strategy.Operator{}
	//策略设置为加法
	op.SetStrategy(&strategy.Add{})
	ret := op.Calculate(2, 1)
	fmt.Printf("add:%d\n", ret)

	op.SetStrategy(&strategy.Reduce{})
	ret = op.Calculate(2, 1)
	fmt.Printf("reduce:%d\n", ret)
}
