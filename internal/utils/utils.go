// Package utils: Critical Go interface adaptation utilities for nil pointer safety
// Provides sophisticated error interface conversion with proper nil handling semantics
// Solves Go's notorious (*T)(nil) != nil interface trap through intelligent adaptation
// Essential for reliable integration with third-party libraries expecting standard error interface
//
// utils: 关键的 Go 接口适配工具，用于 nil 指针安全
// 提供复杂的错误接口转换，具有正确的 nil 处理语义
// 通过智能适配解决 Go 臭名昭著的 (*T)(nil) != nil 接口陷阱
// 对于与期望标准错误接口的第三方库可靠集成至关重要
package utils

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// SafeInterface performs intelligent conversion from Kratos error to standard error interface
// Solves Go's fundamental nil interface problem where (*T)(nil) != nil
// Returns true nil for nil pointers to prevent interface pollution
// Essential for compatibility with libraries expecting clean error interface semantics
//
// Reference: https://go.dev/doc/faq#nil_error
// Technical Details:
// - Go interfaces are (type, value) pairs where both must be nil for interface to be nil
// - (*errors.Error)(nil) creates interface with (concrete type, nil value) != nil
// - This function ensures proper nil interface creation for safe error handling
//
// SafeInterface 执行从 Kratos 错误到标准错误接口的智能转换
// 解决 Go 的基本 nil 接口问题，其中 (*T)(nil) != nil
// 为 nil 指针返回真正的 nil，防止接口污染
// 对于与期望干净错误接口语义的库兼容性至关重要
//
// 参考: https://go.dev/doc/faq#nil_error
// 技术细节:
// - Go 接口是 (类型, 值) 对，两者都必须为 nil 接口才为 nil
// - (*errors.Error)(nil) 创建具有 (具体类型, nil 值) 的接口 != nil
// - 此函数确保正确的 nil 接口创建以进行安全错误处理
func SafeInterface(erk *errors.Error) error {
	if erk != nil {
		return erk
	}
	return nil
}
