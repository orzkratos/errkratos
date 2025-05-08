package erkrequire

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/utils"
	"github.com/stretchr/testify/require"
)

func NoError(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.NoError(t, utils.Adapt(erk), msgAndArgs...)
}

func Error(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.Error(t, utils.Adapt(erk), msgAndArgs...) //这里必须传递个空才行，跟前面的情况相同
}

func Is(t require.TestingT, expected *errors.Error, erk *errors.Error, msgAndArgs ...interface{}) {
	require.Equal(t, expected == nil, erk == nil, msgAndArgs...)
	if expected != nil && erk != nil {
		require.Equal(t, expected.Reason, erk.Reason, msgAndArgs...)
		require.Equal(t, expected.Code, erk.Code, msgAndArgs...)
	}
}
