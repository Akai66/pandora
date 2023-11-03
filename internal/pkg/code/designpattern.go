package code

//go:generate codegen -type=int

// 11-设计模式服务，00-工厂模块，从01开始
const (
	// ErrSimpleFactory - 500: Create simple factory error.
	ErrSimpleFactory int = iota + 110001

	// ErrAbstractFactory - 500: Create abstract factory error.
	ErrAbstractFactory
)

// 11-设计模式服务，01-策略模块，从01开始
const (
	// ErrSwitchStrategy - 500: Switch strategy error.
	ErrSwitchStrategy int = iota + 110101
)

// 11-设计模式服务，02-模板模块，从01开始
const (
	// ErrCookTomato - 500: Cook tomato error.
	ErrCookTomato int = iota + 110201

	// ErrCookEgg - 500: Cook egg error.
	ErrCookEgg
)
