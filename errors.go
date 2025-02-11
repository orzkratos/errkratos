package erkkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/erkkratos/internal/utils"
)

// As 这里使用As就直接能指定类型，这样能够简便些，毕竟在这个语境下的目标类型确定
func As(erx error) (erk *errors.Error, ok bool) {
	ok = errors.As(erx, &erk)
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
