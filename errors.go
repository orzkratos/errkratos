package errkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// Error 再定义个别名，使用它，能避免与官方的errors包名冲突
type Error = errors.Error

// Erk 使用简单的名字简化代码
type Erk = errors.Error

// As 这里使用As就直接能指定类型，这样能够简便些，毕竟在这个语境下的目标类型确定
func As(err error) (erk *errors.Error, ok bool) {
	if ok = errors.As(err, &erk); ok {
		return erk, true
	}
	return nil, false
}

// Is 这里比较相等，直接使用确定的类型，假如要比较不确定的类型，就请用原始的就行
func Is(erk1 *errors.Error, erk2 *errors.Error) bool {
	if erk1 == nil || erk2 == nil {
		return erk1 == nil && erk2 == nil
	}
	return erk1.Is(erk2)
}

// FromError 就是把 error 转换成 *Error 这里封装避免外部再引用 errors 包名，能避免与官方的errors包名冲突
func FromError(err error) *errors.Error {
	return errors.FromError(err)
}

// From 简化函数名，避免写太长的函数名
func From(err error) *errors.Error {
	return FromError(err)
}
