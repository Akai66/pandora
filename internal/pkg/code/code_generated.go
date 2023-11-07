// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by "codegen -type=int"; DO NOT EDIT.

package code

// init register error codes defines in this source code to `github.com/marmotedu/errors`
func init() {
	register(ErrSuccess, 200, "OK")
	register(ErrUnknown, 500, "Internal server error")
	register(ErrBind, 400, "Error occurred while binding the request body to the struct")
	register(ErrValidation, 400, "Validation failed")
	register(ErrTokenInvalid, 401, "Token invalid")
	register(ErrPageNotFound, 404, "Page not found")
	register(ErrDatabase, 500, "Database error")
	register(ErrUserNotFound, 404, "User not found error")
	register(ErrSimpleFactory, 500, "Create simple factory error")
	register(ErrAbstractFactory, 500, "Create abstract factory error")
	register(ErrSwitchStrategy, 500, "Switch strategy error")
	register(ErrCookTomato, 500, "Cook tomato error")
	register(ErrCookEgg, 500, "Cook egg error")
}