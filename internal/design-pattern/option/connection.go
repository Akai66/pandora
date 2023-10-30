//选项模式：go的函数不支持给参数设置默认值，因此需要使用选项模式，创建一个带有默认值的struct变量，在需要时，选择性的修改struct变量中一些参数的值
package option

import "time"

const (
	defaultCache   = false
	defaultTimeout = 10 * time.Second
)

type Connection struct {
	addr    string //addr 必填参数
	cache   bool
	timeout time.Duration
}

// 定义可选参数结构体
type options struct {
	cache   bool
	timeout time.Duration
}

// 定义应用可选项参数的接口
type option interface {
	apply(*options)
}

// 定义optFunc实现option接口
type optFunc func(*options)

func (f optFunc) apply(opts *options) {
	f(opts)
}

// optFunc类型实现了option接口，实现的apply方法实际就是调用optFunc本身，将参数opts.cache修改为自定义的cache值
func WithCache(cache bool) option {
	return optFunc(func(opts *options) {
		opts.cache = cache
	})
}

func WithTimeout(timeout time.Duration) option {
	return optFunc(func(opts *options) {
		opts.timeout = timeout
	})
}

func NewConnection(addr string, opts ...option) (*Connection, error) {
	//先设置默认参数值
	finOpts := &options{
		cache:   defaultCache,
		timeout: defaultTimeout,
	}

	//依次执行apply方法，修改可选参数值
	for _, opt := range opts {
		opt.apply(finOpts)
	}

	//使用opts中最终的参数值，创建Connection实例
	return &Connection{
		addr:    addr,
		cache:   finOpts.cache,
		timeout: finOpts.timeout,
	}, nil
}
