package errkratos_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos"
	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

// TestAs verifies type-safe conversion from generic error to Kratos error type
// Tests both non-nil error conversion and nil pointer interface handling
// Validates that conversion preserves error information and handles Go's nil interface trap
//
// TestAs 验证从通用 error 到 Kratos 错误类型的类型安全转换
// 测试非 nil 错误转换和 nil 指针接口处理
// 验证转换保留错误信息并处理 Go 的 nil 接口陷阱
func TestAs(t *testing.T) {
	t.Run("non nil", func(t *testing.T) {
		erk := errorspb.ErrorServerDbError("wrong")
		var err error = erk
		// t.Log(erk != nil) // true
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := errkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.NotNil(t, res)
	})

	t.Run("nil value", func(t *testing.T) {
		var erk *errors.Error
		var err error = erk
		// t.Log(erk != nil) // false
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := errkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.Nil(t, res)
	})
}

// TestIs validates error comparison logic with reason and code matching
// Verifies that errors with same type are considered equal regardless of message
// Tests compatibility with both Kratos errors.Is and erero.Ise functions
//
// TestIs 验证基于 reason 和 code 匹配的错误比较逻辑
// 验证具有相同类型的错误无论消息如何都被视为相等
// 测试与 Kratos errors.Is 和 erero.Ise 函数的兼容性
func TestIs(t *testing.T) {
	erk1 := errorspb.ErrorServerDbError("wrong-1")
	erk2 := errorspb.ErrorServerDbError("wrong-2")
	require.True(t, errkratos.Is(erk1, erk2))

	require.True(t, errors.Is(erk1, erk1)) //还是相等
	require.True(t, erero.Ise(erk1, erk1)) //依然相等
}

// TestFrom validates conversion from generic error to Kratos error format
// Tests that converted errors maintain compatibility with original errors
// Verifies error equivalence after conversion through Is comparison
//
// TestFrom 验证从通用 error 到 Kratos 错误格式的转换
// 测试转换后的错误与原始错误保持兼容性
// 通过 Is 比较验证转换后的错误等价性
func TestFrom(t *testing.T) {
	erk1 := errorspb.ErrorServerDbError("wrong")
	var err error = erk1
	erk2 := errkratos.From(err)
	require.True(t, errkratos.Is(erk1, erk2))
}
