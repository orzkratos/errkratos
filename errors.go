package errkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/utils"
)

// Error 再定义个别名，使用它，能避免与官方的errors包名冲突
type Error = errors.Error

// Erk 使用简单的名字简化代码
type Erk = errors.Error

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
