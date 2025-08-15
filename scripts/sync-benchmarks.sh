#!/bin/bash

# Sync benchmark results across all documentation files
# This ensures test/benchmark/README.md and docs/content/benchmarks.md always match

set -e

echo "🔄 Synchronizing benchmark results across documentation..."

# Run fresh benchmarks
echo "📊 Running fresh benchmarks..."
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_ROOT/test"
BENCHMARK_OUTPUT=$(go test ./benchmark/ -bench=. -benchmem -benchtime=1s -timeout=5m | grep "^Benchmark" || echo "# Benchmark execution failed")

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

## Raw Benchmark Data

\`\`\`
$BENCHMARK_OUTPUT
\`\`\`
EOF
)

# Parse benchmark results into comparison table
echo "📊 Creating performance comparison table..."

COMPARISON_TABLE=$(echo "$BENCHMARK_OUTPUT" | awk '
BEGIN { 
    print "| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |"
    print "|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|"
}
{
    # Store all benchmark results by validator name
    # Important: Check more specific patterns first to avoid false matches
    
    if ($1 ~ /BenchmarkGoValidCEL/) {
        # Skip CEL benchmarks in main table
        next
    }
    else if ($1 ~ /BenchmarkGoValidator/) {
        # This is asaskevich/govalidator (BenchmarkGoValidator*)
        validator = $1
        gsub(/BenchmarkGoValidator/, "", validator)
        gsub(/-.*/, "", validator)
        
        time = $3
        allocs = $7
        bytes = $5
        
        if (bytes == "0" || allocs == "0") {
            result = time " / 0 allocs"
        } else {
            result = time " / " bytes " B / " allocs " allocs"
        }
        
        asaskevich_results[validator] = result
        asaskevich_time[validator] = time
    }
    else if ($1 ~ /BenchmarkGoValid/) {
        # This is govalid (BenchmarkGoValid*)
        validator = $1
        gsub(/BenchmarkGoValid/, "", validator)
        gsub(/-.*/, "", validator)
        
        time = $3
        allocs = $7
        bytes = $5
        
        # Format result string
        if (bytes == "0" || allocs == "0") {
            result = time " / 0 allocs"
        } else {
            result = time " / " bytes " B / " allocs " allocs"
        }
        
        govalid_results[validator] = result
        govalid_time[validator] = time
    }
    else if ($1 ~ /BenchmarkGoPlayground/) {
        validator = $1
        gsub(/BenchmarkGoPlayground/, "", validator)
        gsub(/-.*/, "", validator)
        
        time = $3
        allocs = $7
        bytes = $5
        
        if (bytes == "0" || allocs == "0") {
            result = time " / 0 allocs"
        } else {
            result = time " / " bytes " B / " allocs " allocs"
        }
        
        playground_results[validator] = result
        playground_time[validator] = time
    }
    else if ($1 ~ /BenchmarkAsaskevichGovalidator/) {
        validator = $1
        gsub(/BenchmarkAsaskevichGovalidator/, "", validator)
        gsub(/-.*/, "", validator)
        
        time = $3
        allocs = $7
        bytes = $5
        
        if (bytes == "0" || allocs == "0") {
            result = time " / 0 allocs"
        } else {
            result = time " / " bytes " B / " allocs " allocs"
        }
        
        asaskevich_results[validator] = result
        asaskevich_time[validator] = time
    }
    else if ($1 ~ /BenchmarkGookitValidate/) {
        validator = $1
        gsub(/BenchmarkGookitValidate/, "", validator)
        gsub(/-.*/, "", validator)
        
        time = $3
        allocs = $7
        bytes = $5
        
        if (bytes == "0" || allocs == "0") {
            result = time " / 0 allocs"
        } else {
            result = time " / " bytes " B / " allocs " allocs"
        }
        
        gookit_results[validator] = result
        gookit_time[validator] = time
    }
}
END {
    # Print results for each validator
    for (validator in govalid_results) {
        # Skip CEL-related benchmarks in the main table
        if (validator ~ /CEL/) continue
        
        govalid_col = govalid_results[validator]
        playground_col = (validator in playground_results) ? playground_results[validator] : "N/A"
        asaskevich_col = (validator in asaskevich_results) ? asaskevich_results[validator] : "N/A"
        gookit_col = (validator in gookit_results) ? gookit_results[validator] : "N/A"
        
        # Calculate improvement for each library
        govalid_ns = govalid_time[validator]
        gsub(/ns\/op/, "", govalid_ns)
        
        # vs go-playground
        playground_improvement = "N/A"
        if (validator in playground_time) {
            pg_ns = playground_time[validator]
            gsub(/ns\/op/, "", pg_ns)
            if (govalid_ns > 0) {
                improvement = pg_ns / govalid_ns
                playground_improvement = sprintf("**%.1fx**", improvement)
            }
        }
        
        # vs asaskevich
        asaskevich_improvement = "N/A"
        if (validator in asaskevich_time) {
            as_ns = asaskevich_time[validator]
            gsub(/ns\/op/, "", as_ns)
            if (govalid_ns > 0) {
                improvement = as_ns / govalid_ns
                asaskevich_improvement = sprintf("**%.1fx**", improvement)
            }
        }
        
        # vs gookit
        gookit_improvement = "N/A"
        if (validator in gookit_time) {
            gk_ns = gookit_time[validator]
            gsub(/ns\/op/, "", gk_ns)
            if (govalid_ns > 0) {
                improvement = gk_ns / govalid_ns
                gookit_improvement = sprintf("**%.1fx**", improvement)
            }
        }
        
        printf "| %s | %s | %s | %s | %s | %s | %s | %s |\n", 
               validator, govalid_col, playground_col, playground_improvement, asaskevich_col, asaskevich_improvement, gookit_col, gookit_improvement
    }
}')

# Generate CEL benchmarks table separately
echo "📊 Creating CEL benchmarks table..."

CEL_TABLE=$(echo "$BENCHMARK_OUTPUT" | awk '
BEGIN { 
    print "| CEL Validator | govalid (ns/op) | Allocations |"
    print "|---------------|-----------------|-------------|"
}
{
    if ($1 ~ /BenchmarkGoValidCEL/) {
        validator = $1
        gsub(/BenchmarkGoValid/, "", validator)
        gsub(/-.*/, "", validator)
        
        time = $3
        allocs = $7
        bytes = $5
        
        if (bytes == "0" || allocs == "0") {
            alloc_str = "0 allocs"
        } else {
            alloc_str = bytes " B / " allocs " allocs"
        }
        
        printf "| %s | %s | %s |\n", validator, time, alloc_str
    }
}')

# Update test/benchmark/README.md
echo "📄 Updating test/benchmark/README.md..."
cd "$PROJECT_ROOT"

cat > test/benchmark/README.md << EOF
# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

$BENCHMARK_DATA

## Performance Comparison

$COMPARISON_TABLE

## CEL Expression Validation (govalid Exclusive)

$CEL_TABLE

CEL (Common Expression Language) support allows complex runtime expressions with near-zero overhead.

## Running Benchmarks

\`\`\`bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=10s

# Run specific validator benchmarks
go test ./benchmark/ -bench=BenchmarkGoValid{ValidatorName} -benchmem
\`\`\`
EOF

# docs/content/benchmarks.md is now deprecated - we only maintain test/benchmark/README.md
echo "📄 Skipping docs/content/benchmarks.md (deprecated - using test/benchmark/README.md instead)"

# Skip updating _index.md - keep it static
echo "📄 Skipping docs/content/_index.md (performance table is now static)"

echo ""
echo "✅ Benchmark synchronization complete!"
echo ""
echo "📁 Updated files:"
echo "  - test/benchmark/README.md"
echo ""
echo "🔍 Benchmark data updated from: $BENCH_DATE"
echo ""
echo "📊 Full benchmark results: test/benchmark/README.md"
