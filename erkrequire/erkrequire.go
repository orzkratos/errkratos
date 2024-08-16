package erkrequire

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/stretchr/testify/require"
)

func NoError(t *testing.T, erk *errors.Error, msgAndArgs ...interface{}) {
	if erk != nil { //这里必须在非空时才能让它去判断，否则即使是 nil 也会报错
		require.NoError(t, erk, msgAndArgs...)
	}
}

func Error(t *testing.T, erk *errors.Error, msgAndArgs ...interface{}) {
	if erk == nil {
		require.Error(t, nil, msgAndArgs...) //这里必须传递个空才行，跟前面的情况相同
	}
}
