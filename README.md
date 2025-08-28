# errkratos

Advanced Kratos error handling utilities with type-safe operations and nil interface trap prevention.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Key Features

🎯 **Type-Safe Error Handling**: Simplified API for Kratos error manipulation without naming conflicts  
⚡ **Safe Error Wrapping**: Solves Go's notorious (*T)(nil) != nil trap through intelligent adaptation  
🔄 **Testing Integration**: Complete testify/assert and testify/require wrappers for Kratos errors

## Install

```bash
go get github.com/orzkratos/errkratos
```

## Usage

### Basic Error Handling

```go
import "github.com/orzkratos/errkratos"

// Type-safe error conversion
err := someFunction()
if erk, ok := errkratos.As(err); ok {
    fmt.Printf("Kratos error: %s (code: %d)\n", erk.Reason, erk.Code)
}

// Error comparison
erk1 := errors.BadRequest("INVALID_INPUT", "missing field")
erk2 := errors.BadRequest("INVALID_INPUT", "wrong format")
if errkratos.Is(erk1, erk2) {
    // Same error type (reason and code match)
}

// Convert any error to Kratos error
erk := errkratos.From(err)
```

### Testing with Assert

```go
import "github.com/orzkratos/errkratos/must/erkassert"

func TestSomething(t *testing.T) {
    var erk *errors.Error
    
    // Assert no error (handles nil interface correctly)
    erkassert.NoError(t, erk)
    
    // Assert error exists
    erk = errors.InternalServer("SERVER_ERROR", "database failed")
    erkassert.Error(t, erk)
    
    // Assert error equality
    expected := errors.BadRequest("INVALID_INPUT", "test")
    erkassert.Is(t, expected, erk)
}
```

### Testing with Require

```go
import "github.com/orzkratos/errkratos/must/erkrequire"

func TestCritical(t *testing.T) {
    var erk *errors.Error
    
    // Require no error (fails immediately if error exists)
    erkrequire.NoError(t, erk)
    
    // Continue only if no error...
}
```

### Production Error Enforcement

```go
import "github.com/orzkratos/errkratos/must/erkmust"

func criticalOperation() {
    erk := doSomethingImportant()
    
    // Panic if error exists (with structured logging)
    erkmust.Done(erk)
    
    // Or use Must (same behavior, different name)
    erkmust.Must(erk)
}
```

## Package Structure

```
errkratos/
├── errors.go           # Core API (As, Is, From)
├── must/               # Testing and enforcement tools
│   ├── erkassert/      # testify/assert wrapper
│   ├── erkrequire/     # testify/require wrapper
│   └── erkmust/        # Production panic utilities
└── internal/
    └── utils/          # Nil interface adaptation
```

## Why errkratos?

### The Nil Interface Problem

Go has a well-known issue where a typed nil pointer doesn't equal nil when converted to an interface:

```go
var erk *errors.Error = nil
var err error = erk
fmt.Println(erk == nil)  // true
fmt.Println(err == nil)  // false (!!)
```

This causes serious issues in error handling. errkratos solves this through intelligent adaptation in all its functions.

### Clean Naming

The `Erk` type alias avoids import conflicts between standard `errors` package and Kratos `errors`:

```go
// Instead of this confusion:
import (
    stderrors "errors"
    "github.com/go-kratos/kratos/v2/errors"
)

// Just use:
import "github.com/orzkratos/errkratos"
// And work with errkratos.Erk
```

## Related Projects

- [ebzkratos](https://github.com/orzkratos/ebzkratos) - Error wrapper that doesn't implement error interface

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-08-28 08:33:43.829511 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share your use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize by reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo for new releases and features
- 🌟 **Success stories?** Share how this package improved your workflow
- 💬 **General feedback?** All suggestions and comments are welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage interface).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement your changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation for user-facing changes and use meaningful commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/errkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/errkratos)