package erkkratos

import "github.com/go-kratos/kratos/v2/errors"

// EmkBottle 封装错误处理所需的属性
type EmkBottle struct {
	makeErkFunc func(format string, args ...interface{}) *errors.Error
	message     string
	metaKeyName string
}

// NewEmkBottle 创建一个新的 EmkBottle 实例
func NewEmkBottle(makeErkFunc func(format string, args ...interface{}) *errors.Error, message string, metaKeyName string) *EmkBottle {
	return &EmkBottle{
		makeErkFunc: makeErkFunc,
		message:     message,
		metaKeyName: metaKeyName,
	}
}

// SetErkFunc 设置 makeErkFunc 属性并返回自身，以支持链式调用
func (e *EmkBottle) SetErkFunc(makeErkFunc func(format string, args ...interface{}) *errors.Error) *EmkBottle {
	e.makeErkFunc = makeErkFunc
	return e
}

// SetMessage 设置 message 属性并返回自身，以支持链式调用
func (e *EmkBottle) SetMessage(message string) *EmkBottle {
	e.message = message
	return e
}

// SetMetaKeyName 设置 metaKeyName 属性并返回自身，以支持链式调用
func (e *EmkBottle) SetMetaKeyName(metaKeyName string) *EmkBottle {
	e.metaKeyName = metaKeyName
	return e
}

// Wrap 方法用于包装错误并返回格式化的错误信息和元数据
func (e *EmkBottle) Wrap(erx error) *errors.Error {
	return e.makeErkFunc("%s", e.message).WithCause(erx).WithMetadata(map[string]string{
		e.metaKeyName: erx.Error(),
	})
}
