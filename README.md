# AronaUnflatd - FlatBuffer Schema Recovery Tool

[![Go Version](https://img.shields.io/github/go-mod/go-version/arisu-archive/arona-unflatd)](go.mod)
[![Build Status](https://img.shields.io/github/actions/workflow/status/arisu-archive/arona-unflatd/ci.yml)](https://github.com/arisu-archive/arona-unflatd/actions)
[![Coverage Status](https://codecov.io/gh/arisu-archive/arona-unflatd/branch/master/graph/badge.svg)](https://codecov.io/gh/arisu-archive/arona-unflatd)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

AronaUnflatd is a specialized tool that reconstructs FlatBuffer schema files (.fbs) from decompiled C# code using advanced static analysis techniques.

## Features

- 🛠️ Converts decompiled C# classes to FlatBuffer schemas
- 🧩 Preserves table structures, fields, and attributes
- 🔄 Reconstructs enum definitions with values
- 🌐 Handles nested tables and union types
- 📂 Processes entire directories recursively
- 📊 Maintains field ordering and type information
- ✅ Validation of generated schemas
- 📈 Verbose logging for debugging

## Installation

### From Source

```bash
git clone https://github.com/arisu-archive/arona-unflatd
cd arona-unflatd
make build
```

### Go Install

```bash
go install github.com/arisu-archive/arona-unflatd@latest
```

## Usage

### Basic Conversion

```bash
arona-unflatd decompile \
  -i ./decompiled_csharp \
  -o ./schema_output
```

## Technical Overview

### Parsing Pipeline

1. **Lexical Analysis** - Tree-sitter tokenizes C# source
2. **Syntax Tree Construction** - Builds AST with 50+ node types
3. **Pattern Matching** - Executes S-expression queries.
4. **Semantic Analysis** - Resolves type dependencies
5. **IDL Generation** - Outputs validated .fbs files

### Key Components

| Component          | Responsibility                          | Tech Stack        |
|---------------------|-----------------------------------------|-------------------|
| Parser Engine       | Syntax tree construction                | Tree-sitter C#    |
| AST Walker          | Pattern matching and node traversal     | Go + C bindings   |
| Type Resolver       | C# to FBS type conversion               | Custom rule engine|
| Schema Generator    | IDL formatting and validation           | Template engine   |

## Development

### Build System
```bash
make tidy    # Format code and clean deps
make audit   # Run security checks (govulncheck)
make mocks   # Generate mocks
make test    # Run tests with coverage
make bench   # Performance benchmarking
```

## Contributing

We welcome contributions! Please see our [Contribution Guidelines](CONTRIBUTING.md) and:
- Follow Go style guidelines
- Include Ginkgo test coverage
- Update documentation accordingly

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Disclaimer**: This tool requires properly decompiled C# code. Schema recovery accuracy depends on original binary structure preservation during decompilation.
