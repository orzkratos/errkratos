// Package erkassert: Advanced testify/assert wrapper for Kratos error testing
// Provides type-safe assertion functions specifically designed for Kratos error validation
// Solves nil interface problems through intelligent error adaptation mechanisms
// Optimized for comprehensive test coverage in Kratos-based microservice architectures
//
// erkassert: 高级 testify/assert 包装器，专为 Kratos 错误测试设计
// 提供专门为 Kratos 错误验证设计的类型安全断言函数
// 通过智能错误适配机制解决 nil 接口问题
// 为基于 Kratos 的微服务架构中的全面测试覆盖进行了优化
package erkassert

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/erkadapt"
	"github.com/stretchr/testify/assert"
)

// NoError asserts that Kratos error is nil with proper nil interface handling
// Uses intelligent adaptation to prevent Go's notorious (*T)(nil) != nil trap
// Ensures reliable error absence validation in production test suites
// Essential for validating successful operations in Kratos service testing
//
// NoError 断言 Kratos 错误为 nil，具有正确的 nil 接口处理
// 使用智能适配防止 Go 臭名昭著的 (*T)(nil) != nil 陷阱
// 确保生产测试套件中可靠的错误缺失验证
// 对于 Kratos 服务测试中验证成功操作至关重要
func NoError(t assert.TestingT, erk *errors.Error, msgAndArgs ...interface{}) bool {
	return assert.NoError(t, erkadapt.Adapt(erk), msgAndArgs...)
}

// Error asserts that Kratos error is present with intelligent nil handling
// Performs safe error validation through internal adaptation layer
// Prevents false negatives caused by Go interface nil pointer semantics
// Critical for error condition testing in distributed service environments
//
// Error 断言 Kratos 错误存在，具有智能 nil 处理
// 通过内部适配层执行安全的错误验证
// 防止由 Go 接口 nil 指针语义引起的假阴性
// 对于分布式服务环境中的错误条件测试至关重要
func Error(t assert.TestingT, erk *errors.Error, msgAndArgs ...interface{}) bool {
	return assert.Error(t, erkadapt.Adapt(erk), msgAndArgs...)
}

// Is performs deep equality comparison between two Kratos errors with comprehensive validation
// Validates both nil state consistency and detailed error content matching
// Compares error reason codes and HTTP status codes for complete verification
// Provides thorough error equivalence testing for complex service interaction scenarios
//
// Is 在两个 Kratos 错误之间执行深度相等比较，具有全面验证
// 验证 nil 状态一致性和详细错误内容匹配
// 比较错误原因代码和 HTTP 状态代码以进行完整验证
// 为复杂服务交互场景提供彻底的错误等价测试
func Is(t assert.TestingT, expected *errors.Error, erk *errors.Error, msgAndArgs ...interface{}) bool {
	if ok := assert.Equal(t, expected == nil, erk == nil, msgAndArgs...); !ok {
		return false
	}
	if expected != nil && erk != nil {
		if ok := assert.Equal(t, expected.Reason, erk.Reason, msgAndArgs...); !ok {
			return false
		}
		if ok := assert.Equal(t, expected.Code, erk.Code, msgAndArgs...); !ok {
			return false
		}
	}
	return true
}
