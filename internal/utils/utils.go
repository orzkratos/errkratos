package utils

import "github.com/go-kratos/kratos/v2/errors"

func ToError(erk *errors.Error) error {
	if erk == nil {
		return nil //这里必须做这样的转换，因为两个 nil 是不一样的
	}
	return erk
}
