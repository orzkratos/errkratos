package erkkratos

import "github.com/go-kratos/kratos/v2/errors"

func NewErkFc(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return func(erx error) *errors.Error {
		return erkFc("%s=%s", msg, erx)
	}
}
