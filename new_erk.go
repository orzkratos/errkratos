package erkkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// ErkBottle 封装错误处理所需的属性
type ErkBottle struct {
	efc         func(format string, args ...interface{}) *errors.Error
	caption     string
	punctuation string
}

// NewErkBottle 创建一个新的 ErkBottle 实例
func NewErkBottle(efc func(format string, args ...interface{}) *errors.Error, caption string, punctuation string) *ErkBottle {
	return &ErkBottle{
		efc:         efc,
		caption:     caption,
		punctuation: punctuation,
	}
}

// SetErkFunc 设置 efc 属性并返回自身，以支持链式调用
func (b *ErkBottle) SetErkFunc(efc func(format string, args ...interface{}) *errors.Error) *ErkBottle {
	b.efc = efc
	return b
}

// SetCaption 设置 caption 属性并返回自身，以支持链式调用
func (b *ErkBottle) SetCaption(caption string) *ErkBottle {
	b.caption = caption
	return b
}

// SetPunctuation 设置 punctuation 属性并返回自身，以支持链式调用
func (b *ErkBottle) SetPunctuation(punctuation string) *ErkBottle {
	b.punctuation = punctuation
	return b
}

// Wrap 方法用于包装错误并返回格式化的错误信息
func (b *ErkBottle) Wrap(erx error) *errors.Error {
	return b.efc("%s%s%s", b.caption, b.punctuation, erx).WithCause(erx)
}
