package erkassert_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/erkassert"
	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
)

func TestNoError(t *testing.T) {
	var erk *errors.Error
	// assert.NoError(t, erk) // 这是不符合预期的
	erkassert.NoError(t, erk) // 需要使用这个函数
}

func TestError(t *testing.T) {
	var erk = errors_example.ErrorServerDbError("msg=%s", erero.New("wac"))
	ok := erkassert.Error(t, erk)
	must.TRUE(ok)
}

func TestIs(t *testing.T) {
	erkA := errors_example.ErrorServerDbError("a")
	erkB := errors_example.ErrorServerDbError("b")
	ok := erkassert.Is(t, erkA, erkB)
	must.TRUE(ok)
}
