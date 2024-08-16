package erkrequire

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/stretchr/testify/require"
)

func NoError(t *testing.T, erk *errors.Error, msgAndArgs ...interface{}) {
	require.NoError(t, erk2erx(erk), msgAndArgs...)
}

func Error(t *testing.T, erk *errors.Error, msgAndArgs ...interface{}) {
	require.Error(t, erk2erx(erk), msgAndArgs...) //这里必须传递个空才行，跟前面的情况相同
}

func erk2erx(erk *errors.Error) error {
	if erk == nil {
		return nil //这里必须做这样的转换，因为两个 nil 是不一样的
	}
	return erk
}
