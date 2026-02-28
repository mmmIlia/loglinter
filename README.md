# loglinter

[![Go Report Card](https://goreportcard.com/badge/github.com/mmmIlia/loglinter?)](https://goreportcard.com/report/github.com/mmmIlia/loglinter)
[![Build Status](https://github.com/mmmIlia/loglinter/workflows/build/badge.svg)](https://github.com/mmmIlia/loglinter/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A smart Go linter that enforces logging best practices. Designed to work as a plugin for [golangci-lint](https://golangci-lint.run/).

It analyzes your `slog` and `zap` calls to ensure your logs are consistent, clean, and secure.

## Features

- 🔡 **Style:** Ensures messages start with lowercase.
- 🇬🇧 **Language:** Enforces English language for consistency.
- 🧹 **Cleanliness:** Forbids emojis, noisy punctuation (`!`, `?`), and trailing dots.
- 🔒 **Security:** Detects potential sensitive data leaks (passwords, tokens, keys) in both messages and structured fields.
- ✨ **Auto-Fix:** Provides one-click fixes for all style issues in your IDE.

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
      # --- Advanced Configuration Section ---
      settings:
        disable-english: true # Example: disable a specific rule
        sensitive-patterns: "ssn,credit_card" # Example: add custom sensitive patterns

linters:
  enable:
    - loglinter
```

3. Run the linter:

```bash
golangci-lint custom run
```

### Standalone Binary

You can also build and run it directly to use its configuration flags:

```bash
go install github.com/mmmIlia/loglinter/cmd/loglinter@latest
loglinter -disable-english -sensitive-patterns="ssn,credit_card" ./...
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

## How It Works (Architecture)

This linter uses a "Pipeline" pattern for style checks. A log message passes through a series of rules (lowercase -> english -> special chars), each of which cleans up the string. This allows the linter to fix multiple issues in a single pass, providing a superior developer experience.

Security checks (`sensitive-data`) are performed separately on all arguments of a log call, inspecting both string literals and variable names for potential leaks.

## Contributing

Want to add a new rule?
1.  **TextRule (for style):** Implement the `rules.TextRule` interface. It's a pure function that takes a string and returns a fixed string and a list of violations.
2.  **NodeRule (for security/logic):** Implement the `rules.NodeRule` interface. It inspects `ast.Node` and reports diagnostics directly.
3.  Add your rule to the pipeline in `pkg/analyzer/analyzer.go`.
4.  Add tests.

Pull requests are welcome!

## License

MIT