# Benchmark Management Guide

This guide explains how to manage benchmarks and performance documentation in govalid.

## 🤖 Automated Benchmark System

Since implementing GitHub Actions, benchmarks are **fully automated**:

### For Pull Requests
- ✅ Benchmarks run automatically on every PR
- ✅ Results are posted as PR comments
- ✅ Comparison with base branch using `benchstat`
- ✅ Performance regressions are highlighted

### For Main Branch
- ✅ Benchmarks run automatically on merge
- ✅ `test/benchmark/README.md` is updated automatically
- ✅ Documentation files are synchronized automatically
- ✅ Performance categories are calculated automatically

## 📊 What Gets Updated Automatically

### 1. Performance Comparison Tables
All benchmark comparison tables are generated from live benchmark results:
```markdown
| Validator | govalid (ns/op) | go-playground (ns/op) | Improvement |
|-----------|-----------------|----------------------|-------------|
| Required  | 1.840ns        | 83.21ns              | **45.2x faster** |
```

### 2. Performance Categories
Performance categories are automatically calculated based on actual results:
- **🚀 Ultra-Fast (< 3ns)**: Validators under 3ns
- **⚡ Fast (3-40ns)**: Validators between 3-40ns  
- **🎯 govalid Exclusive**: Features unique to govalid

### 3. Platform and Date Information
- Benchmark date, platform, and Go version are automatically updated
- Raw benchmark data is included in full

## 🆕 Adding New Validators

When you implement a new validator, follow these steps to ensure it appears in benchmarks:

### 1. Implement Benchmark Tests
Create benchmark tests in `test/benchmark/benchmark_{validator}_test.go`:

```go
func BenchmarkGoValid{Validator}(b *testing.B) {
    instance := test.{Validator}{Field: "test_value"}
    b.ResetTimer()
    for b.Loop() {
        err := test.Validate{Validator}(&instance)
        if err != nil {
            b.Fatal("unexpected error:", err)
        }
    }
    b.StopTimer()
}

func BenchmarkGoPlayground{Validator}(b *testing.B) {
    validate := validator.New()
    instance := test.{Validator}{Field: "test_value"}
    b.ResetTimer()
    for b.Loop() {
        err := validate.Struct(&instance)
        if err != nil {
            b.Fatal("unexpected error:", err)
        }
    }
    b.StopTimer()
}
```

### 2. Update Benchmark Workflow (if needed)
The GitHub Actions workflow automatically detects new benchmark functions that follow the naming pattern:
- `BenchmarkGoValid{Validator}`
- `BenchmarkGoPlayground{Validator}`

### 3. Add to Comparison Parsing
If your validator name doesn't follow the standard pattern, update the validator list in `.github/workflows/benchmark.yaml`:

```yaml
for validator in Required GT LT GTE LTE MaxLength MinLength MaxItems MinItems Enum Email UUID URL YourNewValidator; do
```

### 4. govalid-Exclusive Features
If your validator has no go-playground equivalent, it will automatically be categorized as "govalid Exclusive".

For validators that extend beyond go-playground capabilities (like map/channel support), document this in the workflow parsing or add special handling.

## 🔧 Manual Overrides (Rarely Needed)

### Custom Performance Categories
If you need custom performance categorization, you can modify the AWK script in `scripts/sync-benchmarks.sh`:

```bash
# Example: Add a new category for validators > 40ns
if (govalid_ns >= 40) {
    slow = slow sprintf("- **%s**: ~%.0fns - %dx faster\\n", validator, govalid_ns, improvement)
}
```

### Special Documentation
For complex validators that need special explanation, add them to the "Key Performance Insights" section in `scripts/sync-benchmarks.sh`.

## 🚨 Important Notes

### DO NOT manually edit these files:
- ❌ `test/benchmark/README.md` (auto-generated)
- ❌ `docs/content/benchmarks.md` (auto-generated)
- ❌ Performance comparison tables (auto-generated)

### You CAN manually edit:
- ✅ `scripts/sync-benchmarks.sh` (for custom logic)
- ✅ Benchmark test files in `test/benchmark/`
- ✅ This guide (`BENCHMARK_GUIDE.md`)

## 🔍 Troubleshooting

### Benchmark Not Appearing
1. Check benchmark function naming: `BenchmarkGoValid{Validator}`
2. Ensure the validator name is in the parsing loop
3. Verify benchmark actually runs: `cd test && go test ./benchmark/ -bench=BenchmarkGoValid{Validator}`

### Performance Category Wrong
1. Check if the performance threshold logic in `sync-benchmarks.sh` needs adjustment
2. Consider if the validator should be in a different category

### Comparison Missing
1. Ensure both `BenchmarkGoValid{Validator}` and `BenchmarkGoPlayground{Validator}` exist
2. Check that both benchmarks actually run and produce results

## 🎯 Best Practices

1. **Follow naming conventions**: `BenchmarkGoValid{Validator}` format
2. **Include competitor benchmarks**: Always add `BenchmarkGoPlayground{Validator}` for comparison
3. **Test locally first**: Run `cd test && go test ./benchmark/ -bench=.` before committing
4. **Let automation work**: Don't manually edit auto-generated files
5. **Document unique features**: If your validator is govalid-exclusive, document why it's valuable

## 🔄 Migration from Manual Process

The old manual process involved:
- ❌ Running benchmarks locally
- ❌ Manually copying results to README
- ❌ Manually updating performance numbers
- ❌ Risk of inconsistency across documentation

The new automated process:
- ✅ Runs in consistent Ubuntu environment
- ✅ Updates all documentation automatically
- ✅ Provides PR-level performance comparison
- ✅ Eliminates manual errors and inconsistencies

This ensures all contributors can see accurate, up-to-date performance data regardless of their local environment.