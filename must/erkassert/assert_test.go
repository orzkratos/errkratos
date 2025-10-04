package erkassert_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/orzkratos/errkratos/must/erkassert"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
)

// TestNoError validates erkassert.NoError with proper nil interface handling
// Demonstrates why testify/assert.NoError fails with Kratos errors
// Shows correct usage pattern for Kratos error absence validation
//
// TestNoError 验证 erkassert.NoError 的正确 nil 接口处理
// 演示为何 testify/assert.NoError 对 Kratos 错误失败
// 展示 Kratos 错误缺失验证的正确使用模式
func TestNoError(t *testing.T) {
	var erk *errors.Error
	// assert.NoError(t, erk) // 这是不符合预期的
	erkassert.NoError(t, erk) // 需要使用这个函数
}

// TestError validates erkassert.Error function with error presence detection
// Tests assertion success return value for programmatic validation
// Verifies correct error existence validation with structured error messages
//
// TestError 验证 erkassert.Error 函数的错误存在检测
// 测试断言成功返回值用于编程验证
// 验证结构化错误消息的正确错误存在验证
func TestError(t *testing.T) {
	var erk = errorspb.ErrorServerDbError("msg=%s", erero.New("wac"))
	ok := erkassert.Error(t, erk)
	must.TRUE(ok)
}

// TestIs validates erkassert.Is function with Kratos error type comparison
// Verifies that errors with same reason and code are considered equal
// Tests assertion return value for integration with other validation logic
//
// TestIs 验证 erkassert.Is 函数的 Kratos 错误类型比较
// 验证具有相同 reason 和 code 的错误被视为相等
// 测试断言返回值以集成其他验证逻辑
func TestIs(t *testing.T) {
	erkA := errorspb.ErrorServerDbError("a")
	erkB := errorspb.ErrorServerDbError("b")
	ok := erkassert.Is(t, erkA, erkB)
	must.TRUE(ok)
}
