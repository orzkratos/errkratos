package utils_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestAdapt(t *testing.T) {
	runSuccess := func() *errors.Error {
		return nil
	}

	{
		erk := runSuccess()
		var err error = erk
		require.Error(t, err) //这里有问题
	}
	{
		erk := runSuccess()
		var err = utils.Adapt(erk)
		require.NoError(t, err) //这才是对的
	}
}
