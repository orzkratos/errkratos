package utils_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestToError(t *testing.T) {
	correctLogic := func() *errors.Error {
		return nil
	}

	{
		erk := correctLogic()
		var err error = erk
		require.Error(t, err) //这里有问题
	}
	{
		erk := correctLogic()
		var err = utils.ToError(erk)
		require.NoError(t, err) //这才是对的
	}
}
