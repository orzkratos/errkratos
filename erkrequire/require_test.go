package erkrequire_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/erkrequire"
	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/yyle88/erero"
)

func TestNoError(t *testing.T) {
	var erk *errors.Error
	// require.NoError(t, erk) // 这是不符合预期的
	erkrequire.NoError(t, erk) // 需要使用这个函数
}

func TestError(t *testing.T) {
	var erk = errorspb.ErrorServerDbError("msg=%s", erero.New("wac"))
	erkrequire.Error(t, erk)
}

func TestIs(t *testing.T) {
	erkA := errorspb.ErrorServerDbError("a")
	erkB := errorspb.ErrorServerDbError("b")
	erkrequire.Is(t, erkA, erkB)
}
