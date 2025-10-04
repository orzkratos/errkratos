package erkadapt_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/erkadapt"
	"github.com/stretchr/testify/require"
)

// TestAdapt demonstrates Go's nil interface trap and validates the Adapt function solution
// Compares direct type casting (which fails) vs safe adaptation (which works correct)
// Proves that (*T)(nil) != nil when cast to interface, and shows how Adapt fixes this issue
//
// TestAdapt 演示 Go 的 nil 接口陷阱并验证 Adapt 函数的解决方案
// 比较直接类型转换（失败）与安全适配（正确工作）
// 证明 (*T)(nil) 转换为接口时不等于 nil，并展示 Adapt 如何修复此问题
func TestAdapt(t *testing.T) {
	runSuccess := func() *errors.Error {
		return nil
	}

	t.Run("direct cast", func(t *testing.T) {
		erk := runSuccess()
		var err error = erk
		require.Error(t, err) //这里有问题
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是
	})

	t.Run("adapt", func(t *testing.T) {
		erk := runSuccess()
		var err = erkadapt.Adapt(erk)
		require.NoError(t, err) //这才是对的
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是
	})
}
