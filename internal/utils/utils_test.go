package utils_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestSafeInterface(t *testing.T) {
	runSuccess := func() *errors.Error {
		return nil
	}

	{
		erk := runSuccess()
		var err error = erk
		require.Error(t, err) //这里有问题
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是
	}
	{
		erk := runSuccess()
		var err = utils.SafeInterface(erk)
		require.NoError(t, err) //这才是对的
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是
	}
}
