package errbizkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos"
)

// Ebz 把 erk 包一层，但不实现 golang 的 error 接口，这样就能避免空指针也被当成是有 error 的，避免踩坑也避免时刻保持警惕，让开发者在写业务时保持放松状态
type Ebz struct {
	Erk *errkratos.Erk
}

// NewEbz 返回的 Ebz 只是把 erk 包一层
func NewEbz(erk *errkratos.Erk) *Ebz {
	return &Ebz{
		Erk: erk,
	}
}

// As 这里使用As就直接能指定类型，这样能够简便些，毕竟在这个语境下的目标类型确定
func As(err error) (ebz *Ebz, ok bool) {
	if erk, ok := errkratos.As(err); ok {
		if erk == nil {
			return nil, true
		}
		return NewEbz(erk), true
	}
	return nil, false
}

// Is 这里比较相等，直接使用确定的类型，假如要比较不确定的类型，就请用原始的就行
func Is(ebz1 *Ebz, ebz2 *Ebz) bool {
	if ebz1 == nil || ebz2 == nil {
		return ebz1 == nil && ebz2 == nil
	}
	return errkratos.Is(ebz1.Erk, ebz2.Erk)
}

// FromError 就是把 error 转换成 *Error 这里封装避免外部再引用 errors 包名，能避免与官方的errors包名冲突
func FromError(err error) *Ebz {
	erk := errors.FromError(err)
	if erk != nil {
		return NewEbz(erk)
	}
	return nil
}

// From 简化函数名，避免写太长的函数名
func From(err error) *Ebz {
	return FromError(err)
}
