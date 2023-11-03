package code

import (
	"net/http"

	"github.com/marmotedu/errors"
	"github.com/novalagung/gubrak"
)

type ErrCode struct {
	// 对应业务错误码
	C int

	// 对应HTTP协议的状态码
	HTTP int

	// 对外展示的错误信息
	Ext string

	// 错误文档链接
	Ref string
}

// 在编译时验证ErrCode是否实现了errors.Coder接口
var _ errors.Coder = &ErrCode{}

func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return http.StatusInternalServerError
	}
	return coder.HTTP
}

func (coder ErrCode) String() string {
	return coder.Ext
}

func (coder ErrCode) Reference() string {
	return coder.Ref
}

func (coder ErrCode) Code() int {
	return coder.C
}

// register 向errors包的codes字典中注册业务状态码信息
func register(code int, httpStaus int, message string, ref ...string) {
	find, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStaus)
	if !find {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	var reference string
	if len(ref) > 0 {
		reference = ref[0]
	}

	coder := &ErrCode{
		C:    code,
		HTTP: httpStaus,
		Ext:  message,
		Ref:  reference,
	}
	errors.MustRegister(coder)
}
