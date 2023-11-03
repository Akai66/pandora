package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmotedu/errors"

	"github.com/Akai66/pandora/internal/pkg/code"
)

type response struct {
	httpStatus int
	code       int
	message    string
	reference  string
	data       interface{}
}

type optError struct {
	author string
	msg    string
}

func (o *optError) Error() string {
	return o.msg
}

func main() {
	name := "jack"
	err := bindUser(name)
	// %s:返回可以直接展示给用户的错误信息，就是coder对应的ext
	fmt.Println("====================> %s <====================")
	fmt.Printf("%s\n\n", err)                //print:Error occurred while binding the request body to the struct
	fmt.Printf("%s\n\n", errors.Unwrap(err)) //print:User not found error

	// %v:alias for %s
	fmt.Println("====================> %v <====================")
	fmt.Printf("%v\n\n", err)
	fmt.Printf("%v\n\n", errors.Unwrap(err))

	// %-v:打印出调用栈，错误码，展示给用户的错误信息，展示给研发的错误信息(只展示错误链中最后一个错误)
	fmt.Println("====================> %-v <====================")
	fmt.Printf("%-v\n\n", err)
	fmt.Printf("%-v\n\n", errors.Unwrap(err))

	// %+v:打印出调用栈，错误码，展示给用户的错误信息，展示给研发的错误信息(展示错误链中的所有错误)
	fmt.Println("====================> %+v <====================")
	fmt.Printf("%+v\n\n", err)

	// %#-v:json格式打印出调用栈，错误码，展示给用户的错误信息，展示给研发的错误信息(只展示错误链中最后一个错误)
	fmt.Println("====================> %#-v <====================")
	fmt.Printf("%#-v\n\n", err)

	// %#+v:json格式打印出调用栈，错误码，展示给用户的错误信息，展示给研发的错误信息(展示错误链中的所有错误)
	fmt.Println("====================> %#+v <====================")
	fmt.Printf("%#+v\n\n", err)

	/*
		[
		{"caller":"#2 /Users/liukai/oriMacPro/code/pandora/cmd/error-withcode/main.go:42 (main.bindUser)","code":100003,"error":"binding user jack failed","message":"Error occurred while binding the request body to the struct"},
		{"caller":"#1 /Users/liukai/oriMacPro/code/pandora/cmd/error-withcode/main.go:50 (main.getUser)","code":100201,"error":"get user jack failed","message":"User not found error"},
		{"caller":"#0 /Users/liukai/oriMacPro/code/pandora/cmd/error-withcode/main.go:57 (main.queryUser)","code":100201,"error":"user jack not found","message":"User not found error"}
		]
	*/

	res := writeResponse(err, nil)
	fmt.Printf("%+v\n\n", res)

	// ------------------------------------------------------------------
	// 测试Is和As方法

	fmt.Println(errors.IsCode(err, code.ErrUserNotFound))                                         // true，只要error链中包含code.ErrUserNotFound，就会返回true
	fmt.Println(errors.Is(err, errors.WithCode(code.ErrUserNotFound, "user %s not found", name))) // false 底层的Is方法，先判断二者是否是可比较类型，其次再判断是否实现了Is方法

	errAny := &optError{
		author: "rose",
		msg:    "test error",
	}
	//errAny := errors.New("test")
	var errOpt *optError
	// errOpt本身是一个指针类型，未初始化，所以指向的内容是nil，虽然指向的内容是nil，但是还会存储指向内容的类型信息
	// 这里应该传&errOpt，指针自己的地址，内部通过地址找到该指针，如果errAny指向的内容和该指针指向的内容的类型相同，就将指针指向的内容设置为errAny指向的内容
	fmt.Println(errors.As(errAny, &errOpt)) // true
	errOpt.author = "jack"
	fmt.Println(errAny.author, errOpt.author) // jack jack
	//fmt.Println(errOpt.author)

}

func bindUser(name string) error {
	if err := getUser(name); err != nil {
		// WrapC不管err类型是啥，都会返回withCode类型
		return errors.WrapC(err, code.ErrBind, "binding user %s failed", name)
	}
	return nil
}

func getUser(name string) error {
	if err := queryUser(name); err != nil {
		// 如果被Wrap的err是withCode类型，那么也会返回withCode类型，否则返回withStack类型
		return errors.Wrapf(err, "get user %s failed", name)
	}
	return nil
}

func queryUser(name string) error {
	// coder对象中的Ext字段是对外展示的错误信息(不能包含敏感信息)，后面的user %s not found是才是error本身的错误信息，后续可以通过log.Errorf()记录下来
	return errors.WithCode(code.ErrUserNotFound, "user %s not found", name)
}

func writeResponse(err error, data interface{}) response {
	if err != nil {
		// 底层只需要将错误一层层向上返回，最后统一在最上层记录错误日志
		log.Printf("%#+v", err)
		// 将withCode类型的error，根据code值从codes字典中获取Coder实例
		coder := errors.ParseCoder(err)
		return response{
			httpStatus: coder.HTTPStatus(),
			code:       coder.Code(),
			message:    coder.String(),
			reference:  coder.Reference(),
		}
	}
	return response{
		httpStatus: http.StatusOK,
		code:       code.ErrSuccess,
		message:    "OK",
		data:       data,
	}
}
