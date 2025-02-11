package erkkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// ErkBottle 封装错误处理所需的属性
type ErkBottle struct {
	makeErkFunc func(format string, args ...interface{}) *errors.Error
	caption     string //caption 的意思是 标题、字幕、说明文字。
	punctuation string //在编程或者文本处理中，punctuation 通常指的是像 .,!?;:"'()[]{} 这些标点符号。
}

// NewErkBottle 创建一个新的 ErkBottle 实例
func NewErkBottle(makeErkFunc func(format string, args ...interface{}) *errors.Error, caption string, punctuation string) *ErkBottle {
	return &ErkBottle{
		makeErkFunc: makeErkFunc,
		caption:     caption,
		punctuation: punctuation,
	}
}

// SetErkFunc 设置 makeErkFunc 属性并返回自身，以支持链式调用
func (b *ErkBottle) SetErkFunc(makeErkFunc func(format string, args ...interface{}) *errors.Error) *ErkBottle {
	b.makeErkFunc = makeErkFunc
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
	return b.makeErkFunc("%s%s%s", b.caption, b.punctuation, erx).WithCause(erx)
}
