package erkbzmust

import (
	"github.com/orzkratos/errkratos/ebzkratos"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done expects no error. Panics if the provided error is non-nil.
// Done 期望没有错误。如果提供的错误不为 nil，则触发 panic。
func Done(ebz *ebzkratos.Ebz) {
	if ebz != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("NO ERROR BUG", zap.Error(ebz.Erk))
	}
}

// Must expects no error. Panics if the provided error is non-nil.
// Must 期望没有错误。如果提供的错误不为 nil，则触发 panic。
func Must(ebz *ebzkratos.Ebz) {
	if ebz != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("ERROR", zap.Error(ebz.Erk))
	}
}
