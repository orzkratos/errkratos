// Package erkmust: Production-grade error enforcement utilities for Kratos applications
// Provides fail-fast error checking with structured logging and immediate panic termination
// Implements rigorous error validation for critical production code paths
// Optimized for high-reliability Kratos services where error tolerance is zero
//
// erkmust: 生产级 Kratos 应用错误强制工具
// 提供快速失败错误检查，具有结构化日志记录和立即 panic 终止
// 为关键生产代码路径实现严格的错误验证
// 为错误容忍度为零的高可靠性 Kratos 服务进行了优化
package erkmust

import (
	"github.com/orzkratos/errkratos"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done enforces successful operation completion with immediate panic on error presence
// Terminates application execution instantly if any error is detected
// Provides structured logging context before panic for debugging assistance
// Essential for validating critical operations where failure is unacceptable
//
// Done 强制成功完成操作，错误存在时立即 panic
// 如果检测到任何错误则立即终止应用程序执行
// 在 panic 前提供结构化日志上下文以协助调试
// 对于验证失败不可接受的关键操作至关重要
func Done(erk *errkratos.Erk) {
	if erk != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("NO ERROR BUG", zap.Error(erk))
	}
}

// Must enforces error absence with aggressive panic-based error handling
// Immediately crashes application if unexpected error is encountered
// Records comprehensive error information before termination for post-mortem analysis
// Critical for production systems requiring absolute error intolerance
//
// Must 强制错误缺失，采用激进的基于 panic 的错误处理
// 如果遇到意外错误则立即崩溃应用程序
// 在终止前记录全面的错误信息以供事后分析
// 对于需要绝对错误不容忍的生产系统至关重要
func Must(erk *errkratos.Erk) {
	if erk != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("ERROR", zap.Error(erk))
	}
}
