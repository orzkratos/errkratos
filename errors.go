package errkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/utils"
)

// Error 再定义个别名，使用它，能避免与官方的errors包名冲突
type Error = errors.Error

// Erk 使用简单的名字简化代码
type Erk = errors.Error

// Ero 把 erk 包一层，但不实现 golang 的 error 接口，这样就能避免空指针也被当成是有 error 的，避免踩坑也避免时刻保持警惕，让开发者在写业务时保持放松状态
type Ero struct {
	Erk *Erk
}

// NewEro 返回的 Ero 只是把 erk 包一层
func NewEro(erk *Erk) *Ero {
	return &Ero{
		Erk: erk,
	}
}

// FmtEro 返回的 Ero 只是把 erk 包一层
func FmtEro(newErk func(format string, args ...interface{}) *errors.Error, format string, args ...interface{}) *Ero {
	return NewEro(newErk(format, args...))
}

// As 这里使用As就直接能指定类型，这样能够简便些，毕竟在这个语境下的目标类型确定
func As(err error) (erk *errors.Error, ok bool) {
	ok = errors.As(err, &erk)
	if !ok {
		return nil, false
	}
	return erk, true
}

// Is 这里比较相等，直接使用确定的类型，假如要比较不确定的类型，就请用原始的就行
func Is(erk *errors.Error, target *errors.Error) bool {
	if erk == nil || target == nil {
		return utils.GetValue(erk == nil && target == nil)
	}
	return erk.Is(target)
}
