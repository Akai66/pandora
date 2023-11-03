# 错误码

！！系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。

## 功能说明

如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：

```json
{
  "code": 100101,
  "message": "Database error"
}
```

上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

## 错误码列表

系统支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| ErrSuccess | 100001 | 200 | OK |
| ErrUnknown | 100002 | 500 | Internal server error |
| ErrBind | 100003 | 400 | Error occurred while binding the request body to the struct |
| ErrValidation | 100004 | 400 | Validation failed |
| ErrTokenInvalid | 100005 | 401 | Token invalid |
| ErrPageNotFound | 100006 | 404 | Page not found |
| ErrDatabase | 100101 | 500 | Database error |
| ErrUserNotFound | 100201 | 404 | User not found error |
| ErrSimpleFactory | 110001 | 500 | Create simple factory error |
| ErrAbstractFactory | 110002 | 500 | Create abstract factory error |
| ErrSwitchStrategy | 110101 | 500 | Switch strategy error |
| ErrCookTomato | 110201 | 500 | Cook tomato error |
| ErrCookEgg | 110202 | 500 | Cook egg error |

