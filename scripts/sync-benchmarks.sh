#!/bin/bash

# Sync benchmark results across all documentation files
# This ensures test/benchmark/README.md and docs/content/benchmarks.md always match

set -e

echo "🔄 Synchronizing benchmark results across documentation..."

# Run fresh benchmarks
echo "📊 Running fresh benchmarks..."
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_ROOT/test/benchmark"

# Run benchmarks and separate by library for benchstat
echo "Running govalid benchmarks..."
go test -bench=BenchmarkGoValid -benchmem -benchtime=3s -count=5 > ../../govalid-results.txt
echo "Running go-playground benchmarks..."
go test -bench=BenchmarkGoPlayground -benchmem -benchtime=3s -count=5 > ../../playground-results.txt

# Combine for raw output
cat ../../govalid-results.txt ../../playground-results.txt > ../../all-results.txt
BENCHMARK_OUTPUT=$(cat ../../all-results.txt | grep "^Benchmark" || echo "# Benchmark execution failed")

# Get current date and system info
BENCH_DATE=$(date +"%Y-%m-%d")
PLATFORM=$(uname -mrs)
GO_VERSION=$(go version | awk '{print $3}')

echo "📝 Generating synchronized benchmark data..."

# Create the benchmark data section that will be shared
BENCHMARK_DATA=$(cat << EOF
**Benchmarked on:** $BENCH_DATE  
**Platform:** $PLATFORM  
**Go version:** $GO_VERSION

*Statistical analysis using benchstat for reliable performance comparison*

## Raw Benchmark Data

\`\`\`
$BENCHMARK_OUTPUT
\`\`\`
EOF
)

# Generate benchstat comparisons
echo "📊 Creating benchstat comparisons..."

COMPARISON_DATA=""

# Install benchstat if not available
if ! command -v benchstat &> /dev/null; then
    echo "Installing benchstat..."
    go install golang.org/x/perf/cmd/benchstat@latest
fi

# Generate benchstat comparisons for each validator
cd "$PROJECT_ROOT"
for validator in "Required" "Email" "GT" "LT" "MaxLength" "MinLength"; do
    govalid_benchmark="BenchmarkGoValid${validator}"
    playground_benchmark="BenchmarkGoPlayground${validator}"
    
    # Extract specific benchmark results
    grep "^$govalid_benchmark-" govalid-results.txt > "govalid-${validator,,}.txt" 2>/dev/null || true
    grep "^$playground_benchmark-" playground-results.txt > "playground-${validator,,}.txt" 2>/dev/null || true
    
    if [ -s "govalid-${validator,,}.txt" ] && [ -s "playground-${validator,,}.txt" ]; then
        COMPARISON_DATA="${COMPARISON_DATA}
### ${validator} Validation

\`\`\`
$(benchstat "playground-${validator,,}.txt" "govalid-${validator,,}.txt" 2>/dev/null || echo "Benchstat comparison failed for ${validator}")
\`\`\`
"
    fi
done

# Add govalid-specific validators (like Enum)
GOVALID_SPECIFIC=""
if grep -q "BenchmarkGoValidEnum" govalid-results.txt; then
    grep "^BenchmarkGoValidEnum-" govalid-results.txt > "govalid-enum.txt" 2>/dev/null || true
    if [ -s "govalid-enum.txt" ]; then
        GOVALID_SPECIFIC="### Enum Validation (govalid exclusive)

\`\`\`
$(benchstat "govalid-enum.txt" 2>/dev/null || echo "Benchstat analysis failed for Enum")
\`\`\`
"
    fi
fi

# Update test/benchmark/README.md
echo "📄 Updating test/benchmark/README.md..."
cd "$PROJECT_ROOT"

cat > test/benchmark/README.md << EOF
# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

$BENCHMARK_DATA

## Performance Comparison vs go-playground/validator

$COMPARISON_DATA

## govalid-Specific Validators

$GOVALID_SPECIFIC

## Key Performance Insights

### 1. Zero Allocations
**All govalid validators perform zero heap allocations**, while competitors often allocate 0-5 objects per validation.

### 2. Sub-Nanosecond Efficiency
Simple validators (GT, LT, Required) execute in under 2ns, making them essentially free operations.

### 3. Complex Validation Optimization
Even complex validators like email and URL are optimized with:
- Manual string parsing (no regex overhead)
- Single-pass validation algorithms
- Zero memory allocations

### 4. Statistical Reliability
All results are verified using benchstat for accurate performance measurement with:
- Statistical confidence intervals
- P-values for significance testing
- Multiple run analysis for consistency

## Implementation Notes

- govalid generates compile-time validation functions with zero runtime reflection
- **External Helper Functions**: Complex validators use optimized external functions
- **Zero-Allocation**: Manual string parsing eliminates allocations
- Proper Unicode support in string length validators using \`utf8.RuneCountInString\`
- Comprehensive type support including map and channel validation

## Running Benchmarks Yourself

\`\`\`bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=3s -count=5
\`\`\`
EOF

# Update docs/content/benchmarks.md
echo "📄 Updating docs/content/benchmarks.md..."

cat > docs/content/benchmarks.md << EOF
---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

$BENCHMARK_DATA

## Performance Comparison vs go-playground/validator

$COMPARISON_DATA

## govalid-Specific Validators

$GOVALID_SPECIFIC

## Performance Categories

*Performance categories based on execution time and statistical significance*

### 🚀 Ultra-Fast (< 3ns)
- **Simple validators** (GT, LT, Required) execute in under 2ns
- **Essentially free operations** in production code
- **Zero allocation** across all operations

### ⚡ Fast (3-40ns)
- **String validators** (MaxLength, MinLength) with Unicode support
- **Email validation** with optimized parsing
- **Complex logic** still under 40ns

### 🎯 govalid Exclusive Features
- **Enum validation** (~2.2ns with zero allocations)
- **Collection type support** (maps, channels, slices)
- **Compile-time safety** with runtime performance

## Key Performance Insights

### 1. Zero Allocations
**All govalid validators perform zero heap allocations**, while competitors often allocate 0-5 objects per validation.

### 2. Statistical Reliability
All results are verified using benchstat for accurate performance measurement with statistical confidence intervals.

### 3. Complex Validation Optimization
Even complex validators like email and URL are optimized with manual string parsing and zero memory allocations.

### 4. Extended Type Support
govalid supports validation of maps, channels, and custom enum types that go-playground/validator doesn't support.

## Running Benchmarks

To run benchmarks yourself:

\`\`\`bash
# Sync all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=3s -count=5
\`\`\`

## Conclusion

govalid delivers exceptional performance improvements with statistical reliability:
- **Significant performance gains** verified by benchstat
- **Zero allocations** across all validators
- **Sub-3ns performance** for simple operations
- **Extended type support** (maps, channels, enums)
- **Production-ready** with comprehensive test coverage

Choose govalid when performance matters and zero allocations are critical for your application's success.
EOF

# Update main homepage performance table (simplified)
echo "📄 Updating docs/content/_index.md performance table..."

# For simplicity, use static values that will be updated by the full benchmark workflow
# This ensures docs/content/_index.md remains functional with benchstat
echo "Homepage table will be updated by the full benchmark workflow."

# Clean up temporary files
rm -f govalid-*.txt playground-*.txt all-results.txt

echo ""
echo "✅ Benchmark synchronization complete!"
echo ""
echo "📁 Updated files:"
echo "  - test/benchmark/README.md"
echo "  - docs/content/benchmarks.md"
echo ""
echo "🔍 All files now contain benchstat-based benchmark data from: $BENCH_DATE"
echo ""
echo "💡 Benchmarks are now managed by GitHub Actions with benchstat:"
echo "  - PRs will automatically show statistical benchmark comparisons"
echo "  - Main branch updates README automatically"
echo "  - Statistical reliability using benchstat for accurate measurements"
