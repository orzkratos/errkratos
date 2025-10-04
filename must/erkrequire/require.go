// Package erkrequire: Advanced testify/require wrapper for strict Kratos error validation
// Provides fail-fast assertion functions with immediate test termination on failure
// Implements intelligent nil interface handling for reliable error state verification
// Optimized for critical path testing in production Kratos service validation workflows
//
// erkrequire: 高级 testify/require 包装器，用于严格的 Kratos 错误验证
// 提供快速失败断言函数，失败时立即终止测试
// 实现智能 nil 接口处理，用于可靠的错误状态验证
// 为生产 Kratos 服务验证工作流中的关键路径测试进行了优化
package erkrequire

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/erkadapt"
	"github.com/stretchr/testify/require"
)

// NoError requires that Kratos error is nil with immediate test failure on violation
// Terminates test execution immediately if error is present, preventing cascading failures
// Uses sophisticated adaptation to handle Go's complex nil interface semantics correctly
// Essential for validating prerequisite conditions in multi-step integration tests
//
// NoError 要求 Kratos 错误为 nil，违反时立即测试失败
// 如果存在错误则立即终止测试执行，防止级联失败
// 使用复杂的适配来正确处理 Go 复杂的 nil 接口语义
// 对于多步骤集成测试中验证前提条件至关重要
func NoError(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.NoError(t, erkadapt.Adapt(erk), msgAndArgs...)
}

// Error requires that Kratos error is present with fail-fast validation
// Immediately stops test execution if expected error is absent
// Provides robust error presence validation through internal adaptation mechanisms
// Critical for negative testing scenarios in distributed service architectures
//
// Error 要求 Kratos 错误存在，具有快速失败验证
// 如果预期错误不存在则立即停止测试执行
// 通过内部适配机制提供健壮的错误存在验证
// 对于分布式服务架构中的负面测试场景至关重要
func Error(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.Error(t, erkadapt.Adapt(erk), msgAndArgs...)
}

// Is requires exact equality between two Kratos errors with comprehensive comparison
// Performs strict validation of nil states, error reasons, and HTTP codes
// Terminates test immediately on any inequality to prevent unreliable test results
// Provides definitive error matching verification for complex service interaction testing
//
// Is 要求两个 Kratos 错误完全相等，具有全面比较
// 对 nil 状态、错误原因和 HTTP 代码执行严格验证
// 在任何不相等时立即终止测试，防止不可靠的测试结果
// 为复杂服务交互测试提供明确的错误匹配验证
func Is(t require.TestingT, expected *errors.Error, erk *errors.Error, msgAndArgs ...interface{}) {
	require.Equal(t, expected == nil, erk == nil, msgAndArgs...)
	if expected != nil && erk != nil {
		require.Equal(t, expected.Reason, erk.Reason, msgAndArgs...)
		require.Equal(t, expected.Code, erk.Code, msgAndArgs...)
	}
}
