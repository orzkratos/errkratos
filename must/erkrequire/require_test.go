package erkrequire_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/orzkratos/errkratos/must/erkrequire"
	"github.com/yyle88/erero"
)

// TestNoError validates erkrequire.NoError with fail-fast error checking
// Demonstrates incompatibility of testify/require.NoError with Kratos errors
// Shows proper usage pattern for critical path Kratos error validation
//
// TestNoError 验证 erkrequire.NoError 的快速失败错误检查
// 演示 testify/require.NoError 与 Kratos 错误的不兼容性
// 展示关键路径 Kratos 错误验证的正确使用模式
func TestNoError(t *testing.T) {
	var erk *errors.Error
	// require.NoError(t, erk) // 这是不符合预期的
	erkrequire.NoError(t, erk) // 需要使用这个函数
}

// TestError validates erkrequire.Error with immediate test termination on failure
// Verifies strict error presence requirement for critical test paths
// Tests integration with structured error message formatting
//
// TestError 验证 erkrequire.Error 在失败时立即终止测试
// 验证关键测试路径的严格错误存在要求
// 测试与结构化错误消息格式化的集成
func TestError(t *testing.T) {
	var erk = errorspb.ErrorServerDbError("msg=%s", erero.New("wac"))
	erkrequire.Error(t, erk)
}

// TestIs validates erkrequire.Is with strict Kratos error type matching
// Ensures test fails immediately if error types don't match expectations
// Verifies comprehensive error equality checking for production test suites
//
// TestIs 验证 erkrequire.Is 的严格 Kratos 错误类型匹配
// 确保错误类型不符合预期时测试立即失败
// 验证生产测试套件的全面错误等价性检查
func TestIs(t *testing.T) {
	erkA := errorspb.ErrorServerDbError("a")
	erkB := errorspb.ErrorServerDbError("b")
	erkrequire.Is(t, erkA, erkB)
}
