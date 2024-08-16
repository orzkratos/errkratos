package erkrequire_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/erkkratos/erkrequire"
	"github.com/orzkratos/erkkratos/internal/errors_example"
	"github.com/yyle88/erero"
)

func TestNoError(t *testing.T) {
	var erk *errors.Error
	// require.NoError(t, erk) 这里是不对的
	// 需要使用以下的函数
	erkrequire.NoError(t, erk)
}

func TestError(t *testing.T) {
	var erk = errors_example.ErrorServerDbError("erx=%s", erero.New("wac"))
	erkrequire.Error(t, erk)
}
