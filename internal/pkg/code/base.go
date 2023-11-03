package code

//go:generate codegen -type=int
//go:generate codegen -type=int -doc -output ../../../docs/guide/zh-CN/api/error_code_generated.md

// 10-基础服务，00-常规模块，从01开始
const (
	// ErrSuccess - 200: OK.
	ErrSuccess int = iota + 100001

	// ErrUnknown - 500: Internal server error.
	ErrUnknown

	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind

	// ErrValidation - 400: Validation failed.
	ErrValidation

	// ErrTokenInvalid - 401: Token invalid.
	ErrTokenInvalid

	// ErrPageNotFound - 404: Page not found.
	ErrPageNotFound
)

// 10-基础服务，01-数据库模块，从01开始
const (
	// ErrDatabase - 500: Database error.
	ErrDatabase int = iota + 100101
)

// 10-基础服务，02-用户模块，从01开始
const (
	// ErrUserNotFound - 404: User not found error.
	ErrUserNotFound int = iota + 100201
)
