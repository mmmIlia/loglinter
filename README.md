# loglinter

[![Go Report Card](https://goreportcard.com/badge/github.com/mmmIlia/loglinter)](https://goreportcard.com/report/github.com/mmmIlia/loglinter)
[![Build Status](https://github.com/mmmIlia/loglinter/workflows/build/badge.svg)](https://github.com/mmmIlia/loglinter/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A smart Go linter that enforces logging best practices. Designed to work as a plugin for [golangci-lint](https://golangci-lint.run/).

It checks your `slog` and `zap` calls for:
- 🔡 **Style:** Ensures messages start with lowercase.
- 🇬🇧 **Language:** Enforces English language.
- 🧹 **Cleanliness:** Forbids emojis, noisy punctuation (`!`, `?`), and trailing dots.
- 🔒 **Security:** Detects potential sensitive data leaks (passwords, tokens, keys).

✨ **Bonus:** Auto-fix support for style issues!

## Installation

### With golangci-lint (Recommended)

1. Create a file named `.custom-gcl.yml` in your project root:

```yaml
version: v1.57.2
plugins:
  - module: 'github.com/mmmIlia/loglinter'
    import: 'github.com/mmmIlia/loglinter/plugin'
    version: v1.0.0
```

2. Configure `.golangci.yml`:

```yaml
linters-settings:
  custom:
    loglinter:
      path: .custom-gcl.yml
      description: "Enforce logging standards"
      original-url: github.com/mmmIlia/loglinter
      settings:
        # Optional configuration
        disable-lowercase: false
        disable-english: false
        disable-special-chars: false
        disable-sensitive: false
        sensitive-patterns: "ssn,credit_card" # Comma-separated custom patterns

linters:
  enable:
    - loglinter
```

3. Run the linter:

```bash
golangci-lint custom run
```

### Standalone Binary

You can also build and run it directly:

```bash
go install github.com/mmmIlia/loglinter/cmd/loglinter@latest
loglinter ./...
```

## Rules & Examples

### 1. Lowercase Start
Messages should start with a lowercase letter (unless it's an acronym like HTTP).

❌ `slog.Info("Starting server")`  
✅ `slog.Info("starting server")`

### 2. English Only
Messages must use ASCII characters only.

❌ `slog.Info("Сервер запущен")`  
✅ `slog.Info("server started")`

### 3. No Special Characters
Emojis and noisy punctuation (`!`, `?`, trailing `.`) are forbidden.

❌ `slog.Info("server started! 🚀")`  
✅ `slog.Info("server started")`

### 4. Sensitive Data
Prevents logging of secrets like passwords, tokens, and keys.

❌ `slog.Info("password: " + password)`  
❌ `slog.Info("token=" + token)`  
✅ `slog.Info("user authenticated")`

## License

MIT