#!/bin/bash

# Check if benchmark results are synchronized across all documentation files
# This script is used in CI to detect when benchmarks are out of sync

set -e

echo "🔍 Checking benchmark synchronization..."

# Files to check
BENCHMARK_README="test/benchmark/README.md"
DOCS_BENCHMARK="docs/content/benchmarks.md"
DOCS_INDEX="docs/content/_index.md"

# Check if files exist
if [ ! -f "$BENCHMARK_README" ]; then
    echo "❌ $BENCHMARK_README not found"
    exit 1
fi

if [ ! -f "$DOCS_BENCHMARK" ]; then
    echo "❌ $DOCS_BENCHMARK not found"
    exit 1
fi

if [ ! -f "$DOCS_INDEX" ]; then
    echo "❌ $DOCS_INDEX not found"
    exit 1
fi

# Extract benchmark dates from files
echo "📅 Checking benchmark dates..."

BENCHMARK_DATE=$(grep "Benchmarked on:" "$BENCHMARK_README" | head -1 | sed 's/.*Benchmarked on[*: ]*//; s/\*\*//g; s/  .*//')
DOCS_BENCH_DATE=$(grep "Benchmarked on:" "$DOCS_BENCHMARK" | head -1 | sed 's/.*Benchmarked on[*: ]*//; s/\*\*//g; s/  .*//')

echo "  test/benchmark/README.md: $BENCHMARK_DATE"
echo "  docs/content/benchmarks.md: $DOCS_BENCH_DATE"

# Check if dates match
if [ "$BENCHMARK_DATE" != "$DOCS_BENCH_DATE" ]; then
    echo ""
    echo "❌ BENCHMARK SYNC ERROR: Dates don't match!"
    echo "   test/benchmark/README.md: $BENCHMARK_DATE"
    echo "   docs/content/benchmarks.md: $DOCS_BENCH_DATE"
    echo ""
    echo "🔧 To fix this, run:"
    echo "   make sync-benchmarks"
    echo ""
    exit 1
fi

# Extract key performance numbers and compare
echo "🔢 Checking performance numbers..."

# Extract Required validator performance from both files
README_REQUIRED=$(grep "| Required" "$BENCHMARK_README" | head -1 | awk -F'|' '{print $2}' | tr -d ' ')
DOCS_REQUIRED=$(grep "| Required" "$DOCS_BENCHMARK" | head -1 | awk -F'|' '{print $2}' | tr -d ' ')

echo "  Required validator (README): $README_REQUIRED"
echo "  Required validator (docs):   $DOCS_REQUIRED"

if [ "$README_REQUIRED" != "$DOCS_REQUIRED" ]; then
    echo ""
    echo "❌ BENCHMARK SYNC ERROR: Performance numbers don't match!"
    echo "   test/benchmark/README.md: $README_REQUIRED"
    echo "   docs/content/benchmarks.md: $DOCS_REQUIRED"
    echo ""
    echo "🔧 To fix this, run:"
    echo "   make sync-benchmarks"
    echo ""
    exit 1
fi

# Check if _index.md has recent performance data
INDEX_REQUIRED=$(grep "| Required" "$DOCS_INDEX" | head -1 | awk -F'|' '{print $2}' | tr -d ' ')
echo "  Required validator (index):  $INDEX_REQUIRED"

if [ "$INDEX_REQUIRED" != "$README_REQUIRED" ]; then
    echo ""
    echo "❌ BENCHMARK SYNC ERROR: Homepage performance table is out of sync!"
    echo "   Homepage: $INDEX_REQUIRED"
    echo "   Expected: $README_REQUIRED"
    echo ""
    echo "🔧 To fix this, run:"
    echo "   make sync-benchmarks"
    echo ""
    exit 1
fi

# Check for placeholder values that might indicate outdated data
echo "🔍 Checking for outdated placeholder values..."

if grep -q "1.2ns" "$BENCHMARK_README" "$DOCS_BENCHMARK" "$DOCS_INDEX"; then
    echo ""
    echo "⚠️  WARNING: Found placeholder values (1.2ns) in benchmark data"
    echo "   This might indicate outdated benchmark results"
    echo ""
    echo "🔧 Consider running:"
    echo "   make sync-benchmarks"
    echo ""
fi

# Check raw benchmark data sections
echo "📊 Checking raw benchmark data consistency..."

# Extract first few lines of raw benchmark data
README_RAW=$(sed -n '/```/,/```/p' "$BENCHMARK_README" | sed '1d;$d' | head -3)
DOCS_RAW=$(sed -n '/```/,/```/p' "$DOCS_BENCHMARK" | sed '1d;$d' | head -3)

if [ "$README_RAW" != "$DOCS_RAW" ]; then
    echo ""
    echo "❌ BENCHMARK SYNC ERROR: Raw benchmark data doesn't match!"
    echo ""
    echo "🔧 To fix this, run:"
    echo "   make sync-benchmarks"
    echo ""
    exit 1
fi

echo ""
echo "✅ All benchmark data is synchronized!"
echo ""
echo "📊 Summary:"
echo "  - Benchmark date: $BENCHMARK_DATE"
echo "  - Required validator: $README_REQUIRED"
echo "  - Files checked: 3"
echo "  - Status: ✅ Synchronized"
echo ""

# Additional checks for CI
if [ "${CI:-false}" = "true" ]; then
    echo "🤖 CI Mode: Additional checks..."
    
    # Check if benchmarks are recent (within last 30 days)
    if command -v date >/dev/null 2>&1; then
        BENCH_EPOCH=$(date -j -f "%Y-%m-%d" "$BENCHMARK_DATE" "+%s" 2>/dev/null || echo "0")
        CURRENT_EPOCH=$(date "+%s")
        DAYS_OLD=$(( (CURRENT_EPOCH - BENCH_EPOCH) / 86400 ))
        
        if [ $DAYS_OLD -gt 30 ]; then
            echo ""
            echo "⚠️  CI WARNING: Benchmark data is $DAYS_OLD days old"
            echo "   Consider updating benchmarks for release"
            echo ""
        fi
    fi
fi