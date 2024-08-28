package erkkratos

import "github.com/go-kratos/kratos/v2/errors"

// NewErkFsK 指定错误的前缀让错误打印更加简单
func NewErkFsK(efc func(format string, args ...interface{}) *errors.Error, startWith string, middleOpt string) func(erx error) *errors.Error {
	return func(erx error) *errors.Error {
		return efc("%s%s%s", startWith, middleOpt, erx).WithCause(erx)
	}
}

func NewErkFsB(efc func(format string, args ...interface{}) *errors.Error, startWith string) func(erx error) *errors.Error {
	return NewErkFsK(efc, startWith, " ")
}

func NewErkFsC(efc func(format string, args ...interface{}) *errors.Error, startWith string) func(erx error) *errors.Error {
	return NewErkFsK(efc, startWith, ":")
}

func NewErkFsE(efc func(format string, args ...interface{}) *errors.Error, startWith string) func(erx error) *errors.Error {
	return NewErkFsK(efc, startWith, "=")
}

// NewErkMtK 让错误返回的消息能够被前端直接展示，而把错误的细节放在 metadata 里面
func NewErkMtK(efc func(format string, args ...interface{}) *errors.Error, message string, metaKeyName string) func(erx error) *errors.Error {
	return func(erx error) *errors.Error {
		return efc("%s", message).WithCause(erx).WithMetadata(map[string]string{
			metaKeyName: erx.Error(),
		})
	}
}

func NewErkMtX(efc func(format string, args ...interface{}) *errors.Error, message string) func(erx error) *errors.Error {
	return NewErkMtK(efc, message, "erx")
}

func As(erx error) (erk *errors.Error, ok bool) {
	return erk, errors.As(erx, &erk)
}

func Is(erx error, target error) (ok bool) {
	return errors.Is(erx, target)
}
