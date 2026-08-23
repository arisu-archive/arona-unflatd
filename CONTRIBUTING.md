# Contributing Guide

We welcome contributions to AronaUnflatd! Please follow these guidelines to ensure smooth collaboration.

## Development Setup

### Prerequisites
- Go 1.25+
- GNU Make
- A POSIX-compatible shell (Git Bash works on Windows)

```bash
# Clone and install pinned development tools
git clone https://github.com/arisu-archive/arona-unflatd.git
cd arona-unflatd
make prepare
```

## Coding Standards

### Style Guidelines
1. **Formatting**:
   ```bash
   make tidy  # Runs gofmt + go mod tidy
   ```
2. **Linting**:
   ```bash
   make audit  # Runs 59+ linters via golangci-lint
   ```
3. **Error Handling**:
   ```go
   // Use wrapped errors
   if err := process(); err != nil {
       return fmt.Errorf("process failed: %w", err)
   }
   ```
4. **Documentation**:
   - Document all exported functions
   - Use GoDoc style comments
   - Update README.md for user-facing changes

## Testing Practices

### Unit Tests

```bash
make test  # Runs all tests with coverage
```

**Requirements**:
- Add a regression test for every behavior change and bug fix
- BDD-style tests using Ginkgo/Gomega
- Table-driven tests for complex logic

**Example Test**:
```go
Describe("Schema Conversion", func() {
    When("processing valid C# files", func() {
        It("should generate correct FBS", func() {
            // Test implementation
        })
    })
})
```

### Integration Tests

- Use `testdata/` directory for golden files
- Compare generated schemas against expected output
- Add parser fixtures for malformed or unusual decompiler output

## Pull Request Process

1. **Branch Naming**:
   ```bash
   git checkout -b feat/parser-enhancements
   # or
   git checkout -b fix/issue-123
   ```

2. **Commit Messages**:
   ```bash
   git commit -m "feat(parser): add union type support"
   git commit -m "fix(codegen): handle nested tables"
   ```
   Follow [Conventional Commits](https://www.conventionalcommits.org)

3. **PR Guidelines**:
   - Reference related GitHub issues
   - Include test coverage report
   - Update documentation if needed
   - Keep changes focused (+/- 500 lines)

## Review Process

**Maintainers Will Check**:
- [ ] Code quality meets project standards
- [ ] Tests cover all new functionality
- [ ] Documentation is updated
- [ ] No security vulnerabilities (govulncheck clean)
- [ ] Backward compatibility maintained

## Security Practices

1. **Vulnerability Reporting**:
   - Disclose via GitHub Security Advisories

2. **Dependencies**:
   ```bash
   go mod verify  # Verify dependency integrity
   ```

## License

By contributing, you agree to license your work under the [MIT License](LICENSE).

## Need Help?

- Join our [Discussions](https://github.com/arisu-archive/arona-unflatd/discussions)
- File a [GitHub Issue](https://github.com/arisu-archive/arona-unflatd/issues)
