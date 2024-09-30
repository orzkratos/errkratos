package erkkratos

import "github.com/go-kratos/kratos/v2/errors"

// EmtBottle 封装错误处理所需的属性
type EmtBottle struct {
	efc         func(format string, args ...interface{}) *errors.Error
	message     string
	metaKeyName string
}

// NewEmtBottle 创建一个新的 EmtBottle 实例
func NewEmtBottle(efc func(format string, args ...interface{}) *errors.Error, message string, metaKeyName string) *EmtBottle {
	return &EmtBottle{
		efc:         efc,
		message:     message,
		metaKeyName: metaKeyName,
	}
}

// SetErkFunc 设置 efc 属性并返回自身，以支持链式调用
func (e *EmtBottle) SetErkFunc(efc func(format string, args ...interface{}) *errors.Error) *EmtBottle {
	e.efc = efc
	return e
}

// SetMessage 设置 message 属性并返回自身，以支持链式调用
func (e *EmtBottle) SetMessage(message string) *EmtBottle {
	e.message = message
	return e
}

// SetMetaKeyName 设置 metaKeyName 属性并返回自身，以支持链式调用
func (e *EmtBottle) SetMetaKeyName(metaKeyName string) *EmtBottle {
	e.metaKeyName = metaKeyName
	return e
}

// Wrap 方法用于包装错误并返回格式化的错误信息和元数据
func (e *EmtBottle) Wrap(erx error) *errors.Error {
	return e.efc("%s", e.message).WithCause(erx).WithMetadata(map[string]string{
		e.metaKeyName: erx.Error(),
	})
}
