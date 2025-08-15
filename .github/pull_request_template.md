## Description

<!-- Brief description of what this PR accomplishes -->

## Type of Change

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Refactoring (no functional changes)

## Quality Checklist

Before submitting this pull request, ensure your contribution meets these quality standards:

### Code Quality
- [ ] Code passes `make golangci-lint` without errors
- [ ] All tests pass: `make test`
- [ ] Code follows existing patterns and conventions
- [ ] No hard-coded values (use constants or configuration)
- [ ] Error messages are clear and actionable

### Testing Requirements
- [ ] Golden tests created and passing (`internal/analyzers/govalid/testdata/`)
- [ ] Unit tests implemented with boundary value testing (`test/unit/`)
- [ ] Benchmark tests comparing against popular validation libraries (`test/benchmark/`)
  - [ ] [go-playground/validator](https://github.com/go-playground/validator) (BenchmarkGoPlayground*)
  - [ ] [asaskevich/govalidator](https://github.com/asaskevich/govalidator) (BenchmarkGoValidator*)
  - [ ] [gookit/validate](https://github.com/gookit/validate) (BenchmarkGookitValidate*)
- [ ] All test files generated: `cd test && go generate`

### Performance Standards
- [ ] Zero allocations in validation logic (verified via benchmarks)
- [ ] Performance improvement over existing validation libraries (minimum 2x faster)
- [ ] Benchmark results added to `test/benchmark/README.md`

### Documentation
- [ ] README.md updated with new marker documentation
- [ ] Code includes appropriate comments and examples
- [ ] Validator usage examples provided

### Integration
- [ ] Validator scaffold generated with `make generate-validator MARKER=yourmarker`
- [ ] Registry files automatically updated
- [ ] Binary rebuilt and installed: `go install ./cmd/govalid/`

### Final Verification
- [ ] All GitHub Actions CI checks pass
- [ ] No breaking changes to existing functionality
- [ ] Contribution follows project license requirements

## Additional Notes

<!-- Any additional information about the implementation, design decisions, or areas that need special attention -->

## Related Issues

<!-- Link to any related issues: Fixes #123, Closes #456 -->
