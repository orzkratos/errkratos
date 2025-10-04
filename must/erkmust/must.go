// Package erkmust: Production-grade error enforcement utilities for Kratos applications
// Provides fast-exit error checking with structured logging and immediate panic termination
// Implements rigorous error validation for mission-critical production code paths
// Optimized for zero-error-tolerance Kratos services with strict requirements
//
// erkmust: 生产级 Kratos 应用错误强制工具
// 提供快速退出错误检查，具有结构化日志记录和立即 panic 终止
// 为任务关键生产代码路径实现严格的错误验证
// 为零错误容忍度且具有严格要求的 Kratos 服务进行了优化
package erkmust

import (
	"github.com/orzkratos/errkratos"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done enforces operation completion with immediate panic on error presence
// Terminates application execution right away if errors are detected
// Provides structured logging context before panic for debugging assistance
// Essential when validating mission-critical operations that must succeed
//
// Done 强制操作完成，错误存在时立即 panic
// 如果检测到错误则立即终止应用程序执行
// 在 panic 前提供结构化日志上下文以协助调试
// 在验证必须成功的任务关键操作时至关重要
func Done(erk *errkratos.Erk) {
	if erk != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("NO ERROR BUG", zap.Error(erk))
	}
}

// Must enforces error absence with aggressive panic-based error handling
// Crashes application right away if unexpected errors are encountered
// Records comprehensive error information before termination for post-mortem analysis
// Critical in production systems requiring absolute zero-error-tolerance
//
// Must 强制错误缺失，采用激进的基于 panic 的错误处理
// 如果遇到意外错误则立即崩溃应用程序
// 在终止前记录全面的错误信息以供事后分析
// 在需要绝对零错误容忍度的生产系统中至关重要
func Must(erk *errkratos.Erk) {
	if erk != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("ERROR", zap.Error(erk))
	}
}
