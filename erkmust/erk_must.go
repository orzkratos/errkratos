package erkmust

import (
	"github.com/orzkratos/errkratos"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done expects no error. Panics if the provided error is non-nil.
// Done 期望没有错误。如果提供的错误不为 nil，则触发 panic。
func Done(erk *errkratos.Erk) {
	if erk != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("NO ERROR BUG", zap.Error(erk))
	}
}

// Must expects no error. Panics if the provided error is non-nil.
// Must 期望没有错误。如果提供的错误不为 nil，则触发 panic。
func Must(erk *errkratos.Erk) {
	if erk != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("ERROR", zap.Error(erk))
	}
}
