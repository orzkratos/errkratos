package erkkratos

import "github.com/go-kratos/kratos/v2/errors"

func NewErkFs(erkFc func(format string, args ...interface{}) *errors.Error, msg string, opt string) func(erx error) *errors.Error {
	return func(erx error) *errors.Error {
		return erkFc("%s%s%s", msg, opt, erx)
	}
}

func NewErkFb(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return NewErkFs(erkFc, msg, " ")
}

func NewErkFc(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return NewErkFs(erkFc, msg, ":")
}

func NewErkFe(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return NewErkFs(erkFc, msg, "=")
}
