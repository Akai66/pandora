// 策略模式：定义一组算法，将每个算法都封装起来，并且使它们之间可以互换
package strategy

type IStrategy interface {
	do(int, int) int
}

type Add struct{}

func (*Add) do(a, b int) int {
	return a + b
}

type Reduce struct{}

func (*Reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

func (o *Operator) SetStrategy(str IStrategy) {
	o.strategy = str
}

func (o *Operator) Calculate(a, b int) int {
	return o.strategy.do(a, b)
}
