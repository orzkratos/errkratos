package erkkratos

import "github.com/go-kratos/kratos/v2/errors"

// NewErkFs 指定错误的前缀让错误打印更加简单
func NewErkFs(erkFc func(format string, args ...interface{}) *errors.Error, msg string, opt string) func(erx error) *errors.Error {
	return func(erx error) *errors.Error {
		return erkFc("%s%s%s", msg, opt, erx).WithCause(erx)
	}
}

func NewErkFsb(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return NewErkFs(erkFc, msg, " ")
}

func NewErkFsc(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return NewErkFs(erkFc, msg, ":")
}

func NewErkFse(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return NewErkFs(erkFc, msg, "=")
}

// NewErkFm 让错误返回的消息能够被前端直接展示，而把错误的细节放在 metadata 里面
func NewErkFm(erkFc func(format string, args ...interface{}) *errors.Error, msg string, erk string) func(erx error) *errors.Error {
	return func(erx error) *errors.Error {
		return erkFc("%s", msg).WithCause(erx).WithMetadata(map[string]string{
			erk: erx.Error(),
		})
	}
}

func NewErkFmx(erkFc func(format string, args ...interface{}) *errors.Error, msg string) func(erx error) *errors.Error {
	return NewErkFm(erkFc, msg, "erx")
}

func As(erx error) (erk *errors.Error, ok bool) {
	return erk, errors.As(erx, &erk)
}

func Is(erx error, target error) (ok bool) {
	return errors.Is(erx, target)
}
