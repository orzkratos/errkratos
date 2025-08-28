// Package errkratos: Advanced Kratos error handling utilities with type-safe operations
// Provides simplified API for Kratos error manipulation with naming conflict avoidance
// Features streamlined error conversion, comparison, and type assertion functions
// Optimized for production Kratos applications with clean error handling patterns
//
// errkratos: 高级 Kratos 错误处理工具，提供类型安全操作
// 为 Kratos 错误操作提供简化的 API，避免命名冲突
// 具有简化的错误转换、比较和类型断言函数
// 为生产环境 Kratos 应用的清晰错误处理模式进行了优化
package errkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// Erk provides simplified type alias for errors.Error to avoid naming conflicts with standard library
// Using concise naming reduces code complexity while maintaining clear semantics
// Prevents import conflicts between standard errors package and Kratos errors
//
// Erk 为 errors.Error 提供简化的类型别名，避免与标准库命名冲突
// 使用简洁命名减少代码复杂性，同时保持清晰的语义
// 防止标准 errors 包与 Kratos errors 包之间的导入冲突
type Erk = errors.Error

// As performs type-safe conversion from generic error to Kratos-specific error type
// Returns the converted error and success flag for safe error handling patterns
// Eliminates need for manual type assertions in error processing workflows
// Optimized for high-frequency error handling in production services
//
// As 执行从通用 error 到 Kratos 特定错误类型的类型安全转换
// 返回转换后的错误和成功标志，用于安全的错误处理模式
// 消除错误处理工作流中手动类型断言的需求
// 为生产服务中的高频错误处理进行了优化
func As(err error) (erk *errors.Error, ok bool) {
	if ok = errors.As(err, &erk); ok {
		return erk, true
	}
	return nil, false
}

// Is compares two Kratos errors for logical equivalence with nil-safe operation
// Performs intelligent comparison based on error reason and code matching
// Handles nil pointer cases gracefully to prevent runtime panics
// Essential for error classification and handling logic in service layers
//
// Is 比较两个 Kratos 错误的逻辑等价性，具有 nil 安全操作
// 基于错误原因和代码匹配执行智能比较
// 优雅处理 nil 指针情况，防止运行时 panic
// 对于服务层中的错误分类和处理逻辑至关重要
func Is(erk1 *errors.Error, erk2 *errors.Error) bool {
	if erk1 == nil || erk2 == nil {
		return erk1 == nil && erk2 == nil
	}
	return erk1.Is(erk2)
}

// FromError converts generic error interface to Kratos error with intelligent handling
// Wraps standard library errors into Kratos error format when necessary
// Preserves existing Kratos error metadata and context information
// Provides unified error handling interface for mixed error sources
//
// FromError 将通用 error 接口转换为 Kratos 错误，具有智能处理
// 必要时将标准库错误包装为 Kratos 错误格式
// 保留现有的 Kratos 错误元数据和上下文信息
// 为混合错误源提供统一的错误处理接口
func FromError(err error) *errors.Error {
	return errors.FromError(err)
}

// From provides concise alias for FromError to reduce function name length
// Maintains identical functionality with shorter syntax for frequent usage
// Optimizes developer experience in error-heavy codebases
//
// From 为 FromError 提供简洁别名，减少函数名长度
// 为频繁使用保持相同功能但语法更短
// 在错误处理密集的代码库中优化开发者体验
func From(err error) *errors.Error {
	return FromError(err)
}
