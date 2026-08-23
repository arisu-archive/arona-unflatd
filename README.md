# AronaUnflatd - FlatBuffer Schema Recovery Tool

![Global Version](https://img.shields.io/badge/dynamic/yaml?url=https%3A%2F%2Fba.pokeguy.dev%2Fcom.nexon.bluearchive%2Fversion.txt&query=%24&prefix=v&style=for-the-badge&logo=nexon&label=Global&color=0099ff)![Japan Version](https://img.shields.io/badge/dynamic/yaml?url=https%3A%2F%2Fba.pokeguy.dev%2Fcom.YostarJP.BlueArchive%2Fversion.txt&query=%24&prefix=v&style=for-the-badge&logo=googleplay&label=Yostar&color=7d3cc8)

[![Go Version](https://img.shields.io/github/go-mod/go-version/arisu-archive/arona-unflatd?style=for-the-badge)](go.mod)
[![Build Status](https://img.shields.io/github/actions/workflow/status/arisu-archive/arona-unflatd/ci.yml?style=for-the-badge)](https://github.com/arisu-archive/arona-unflatd/actions)
[![Coverage Status](https://img.shields.io/codecov/c/github/arisu-archive/arona-unflatd?style=for-the-badge
)](https://codecov.io/gh/arisu-archive/arona-unflatd)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)

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
2. **Concrete Syntax Tree Construction** - Builds a tolerant Tree-sitter C# tree
3. **Pattern Matching** - Extracts relevant declarations with S-expression queries
4. **Schema Recovery** - Applies FlatBuffers-specific inference rules
5. **IDL Generation** - Deterministically renders `.fbs` files

### Key Components

| Component          | Responsibility                          | Tech Stack        |
|---------------------|-----------------------------------------|-------------------|
| Parser Engine       | Syntax tree construction                | Tree-sitter C#    |
| Declaration Extractor | Query capture and node traversal      | Go + C bindings   |
| Recovery Rules      | C# pattern to FBS conversion             | Go rules          |
| Schema Generator    | Deterministic IDL formatting             | Go renderer       |

## Development

### Build System
```bash
make tidy    # Format code and clean deps
make audit   # Run security checks (govulncheck)
make test    # Run tests with coverage
make build   # Build the CLI into ./tmp/bin
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
