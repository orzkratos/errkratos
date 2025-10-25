[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/errkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/errkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/errkratos)](https://pkg.go.dev/github.com/orzkratos/errkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/errkratos/main.svg)](https://coveralls.io/github/orzkratos/errkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/errkratos.svg)](https://github.com/orzkratos/errkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/errkratos)](https://goreportcard.com/report/github.com/orzkratos/errkratos)

# errkratos

高级 Kratos 错误处理包，提供类型安全操作和 nil 接口陷阱防护。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **类型安全错误处理**: 为 Kratos 错误操作提供简化的 API，避免命名冲突  
⚡ **安全错误包装**: 通过智能适配解决 Go 臭名昭著的 (*T)(nil) != nil 陷阱  
🔄 **测试集成**: 为 Kratos 错误提供完整的 testify/assert 和 testify/require 包装器

## 安装

```bash
go get github.com/orzkratos/errkratos
```

## 使用方法

### 基础错误处理

```go
import "github.com/orzkratos/errkratos"

// 类型安全的错误转换
err := someFunction()
if erk, ok := errkratos.As(err); ok {
    fmt.Printf("Kratos 错误: %s (代码: %d)\n", erk.Reason, erk.Code)
}

// 错误比较
erk1 := errors.BadRequest("INVALID_INPUT", "缺少字段")
erk2 := errors.BadRequest("INVALID_INPUT", "格式错误")
if errkratos.Is(erk1, erk2) {
    // 相同的错误类型（原因和代码匹配）
}

// 将任何错误转换为 Kratos 错误
erk := errkratos.From(err)
```

### 简洁错误创建 (newerk)

```go
import "github.com/orzkratos/errkratos/newerk"

// 配置原因码字段名用于存储枚举数值
newerk.SetReasonCodeFieldName("numeric_reason_code_enum")

// 使用枚举创建类型安全的错误
erk := newerk.NewError(404, ErrorReason_USER_NOT_FOUND, "用户 %d 未找到", userID)

// 检查错误类型
if newerk.IsError(err, ErrorReason_USER_NOT_FOUND, 404) {
    // 处理用户未找到错误
}
```

### 使用 Assert 测试

```go
import "github.com/orzkratos/errkratos/must/erkassert"

func TestSomething(t *testing.T) {
    var erk *errors.Error
    
    // 断言没有错误（正确处理 nil 接口）
    erkassert.NoError(t, erk)
    
    // 断言错误存在
    erk = errors.InternalServer("SERVER_ERROR", "数据库失败")
    erkassert.Error(t, erk)
    
    // 断言错误相等
    expected := errors.BadRequest("INVALID_INPUT", "测试")
    erkassert.Is(t, expected, erk)
}
```

### 使用 Require 测试

```go
import "github.com/orzkratos/errkratos/must/erkrequire"

func TestCritical(t *testing.T) {
    var erk *errors.Error
    
    // 要求没有错误（如果存在错误立即失败）
    erkrequire.NoError(t, erk)
    
    // 只有在没有错误时继续...
}
```

### 生产环境错误强制执行

```go
import "github.com/orzkratos/errkratos/must/erkmust"

func criticalOperation() {
    erk := doSomethingImportant()
    
    // 如果存在错误则 panic（带结构化日志）
    erkmust.Done(erk)
    
    // 或使用 Must（相同行为，不同名称）
    erkmust.Must(erk)
}
```

## 包结构

```
errkratos/
├── errors.go           # 核心 API (As, Is, From)
├── newerk/             # 简洁错误创建 API
├── erkadapt/           # Nil 接口适配
├── must/               # 测试和强制执行工具
│   ├── erkassert/      # testify/assert 包装器
│   ├── erkrequire/     # testify/require 包装器
│   └── erkmust/        # 生产环境 panic 工具
└── internal/
    └── errorspb/       # 错误定义示例
```

## 为什么使用 errkratos？

### Nil 接口问题

Go 有一个众所周知的问题，当类型化的 nil 指针转换为接口时不等于 nil：

```go
var erk *errors.Error = nil
var err error = erk
fmt.Println(erk == nil)  // true
fmt.Println(err == nil)  // false (!!)
```

这在错误处理中会导致问题。errkratos 通过在所有函数中进行智能适配来解决这个问题。

### 清晰的命名

`Erk` 类型别名避免了标准 `errors` 包和 Kratos `errors` 之间的导入冲突：

```go
// 不用这种混乱的方式：
import (
    stderrors "errors"
    "github.com/go-kratos/kratos/v2/errors"
)

// 只需使用：
import "github.com/orzkratos/errkratos"
// 然后使用 errkratos.Erk
```

## 相关项目

- [ebzkratos](https://github.com/orzkratos/ebzkratos) - 不实现 error 接口的错误包装器

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/errkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/errkratos)