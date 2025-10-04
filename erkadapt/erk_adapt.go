// Package erkadapt: Critical Go interface adaptation utilities to handle nil pointers
// Provides sophisticated error interface conversion with correct nil handling semantics
// Solves Go's notorious (*T)(nil) != nil interface trap through intelligent adaptation
// Essential when integrating with external packages expecting standard error interface
//
// erkadapt: 关键的 Go 接口适配工具，用于处理 nil 指针
// 提供复杂的错误接口转换，具有正确的 nil 处理语义
// 通过智能适配解决 Go 臭名昭著的 (*T)(nil) != nil 接口陷阱
// 在集成期望标准错误接口的外部包时至关重要
package erkadapt

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// Adapt performs intelligent conversion from Kratos error to standard error interface
// Solves Go's core nil interface problem where (*T)(nil) != nil
// Returns true nil when given nil pointers to prevent interface pollution
// Essential when working with packages expecting clean error interface semantics
//
// Reference: https://go.dev/doc/faq#nil_error
// Technical Details:
// - Go interfaces are (type, value) pairs where both must be nil for interface to be nil
// - (*errors.Error)(nil) creates interface with (concrete type, nil value) != nil
// - This function ensures correct nil interface creation for safe error handling
//
// Adapt 执行从 Kratos 错误到标准错误接口的智能转换
// 解决 Go 的核心 nil 接口问题，其中 (*T)(nil) != nil
// 在给定 nil 指针时返回真正的 nil，防止接口污染
// 在使用期望干净错误接口语义的包时至关重要
//
// 参考: https://go.dev/doc/faq#nil_error
// 技术细节:
// - Go 接口是 (类型, 值) 对，两者都必须为 nil 接口才为 nil
// - (*errors.Error)(nil) 创建具有 (具体类型, nil 值) 的接口 != nil
// - 此函数确保正确的 nil 接口创建以进行安全错误处理
func Adapt(erk *errors.Error) error {
	if erk != nil {
		return erk
	}
	return nil
}
