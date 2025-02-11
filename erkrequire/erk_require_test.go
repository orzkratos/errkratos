package erkrequire_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/erkkratos/erkrequire"
	"github.com/orzkratos/erkkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestNoError(t *testing.T) {
	var erk *errors.Error
	require.Error(t, erk)      // 这是不符合预期的
	erkrequire.NoError(t, erk) // 需要使用这个函数
}

func TestError(t *testing.T) {
	var erk = errors_example.ErrorServerDbError("msg=%s", erero.New("wac"))
	erkrequire.Error(t, erk)
}

func TestIs(t *testing.T) {
	erkA := errors_example.ErrorServerDbError("a")
	erkB := errors_example.ErrorServerDbError("b")
	erkrequire.Is(t, erkA, erkB)
}
