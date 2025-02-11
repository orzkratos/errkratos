package erkrequire

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/erkkratos/internal/utils"
	"github.com/stretchr/testify/require"
)

func NoError(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.NoError(t, convert(erk), msgAndArgs...)
}

func Error(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.Error(t, convert(erk), msgAndArgs...) //这里必须传递个空才行，跟前面的情况相同
}

func convert(erk *errors.Error) error {
	if erk == nil {
		return nil //这里必须做这样的转换，因为两个 nil 是不一样的
	}
	return erk
}

func Is(t require.TestingT, expected *errors.Error, erkParam *errors.Error, msgAndArgs ...interface{}) {
	require.Equal(
		t,
		utils.GetValue(convert(expected) == nil),
		utils.GetValue(convert(erkParam) == nil),
		msgAndArgs...,
	)
	if expected != nil && erkParam != nil {
		require.Equal(t, expected.Reason, erkParam.Reason, msgAndArgs...)
		require.Equal(t, expected.Code, erkParam.Code, msgAndArgs...)
	}
}
