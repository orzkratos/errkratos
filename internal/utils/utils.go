package utils

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// Adapt 将 errors.Error 转换为 error 接口类型，以便于和其他库或函数兼容
// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是
// 因此当有错误时，需要单独返回 nil 而不是 (*errors.Error)(nil) 这个结果
func Adapt(erk *errors.Error) error {
	if erk != nil {
		return erk
	}
	return nil
}
