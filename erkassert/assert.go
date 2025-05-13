package erkassert

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/utils"
	"github.com/stretchr/testify/assert"
)

func NoError(t assert.TestingT, erk *errors.Error, msgAndArgs ...interface{}) bool {
	return assert.NoError(t, utils.Adapt(erk), msgAndArgs...)
}

func Error(t assert.TestingT, erk *errors.Error, msgAndArgs ...interface{}) bool {
	return assert.Error(t, utils.Adapt(erk), msgAndArgs...) //这里必须传递个空才行，跟前面的情况相同
}

func Is(t assert.TestingT, expected *errors.Error, erk *errors.Error, msgAndArgs ...interface{}) bool {
	if ok := assert.Equal(t, expected == nil, erk == nil, msgAndArgs...); !ok {
		return false
	}
	if expected != nil && erk != nil {
		if ok := assert.Equal(t, expected.Reason, erk.Reason, msgAndArgs...); !ok {
			return false
		}
		if ok := assert.Equal(t, expected.Code, erk.Code, msgAndArgs...); !ok {
			return false
		}
	}
	return true
}
