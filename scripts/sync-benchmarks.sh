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

# Update main homepage performance table
echo "📄 Updating docs/content/_index.md performance table..."

# Extract key metrics for the homepage table (simplified showcase)
HOMEPAGE_TABLE=$(echo "$BENCHMARK_OUTPUT" | awk '
{
    # Store benchmark results
    if ($1 ~ /BenchmarkGoValidRequired/) govalid_required = $3
    if ($1 ~ /BenchmarkGoPlaygroundRequired/) playground_required = $3
    if ($1 ~ /BenchmarkGoValidatorRequired/) asaskevich_required = $3
    
    if ($1 ~ /BenchmarkGoValidEmail/) govalid_email = $3
    if ($1 ~ /BenchmarkGoPlaygroundEmail/) playground_email = $3
    if ($1 ~ /BenchmarkGoValidatorEmail/) asaskevich_email = $3
    if ($1 ~ /BenchmarkGookitValidateEmail/) gookit_email = $3
    
    if ($1 ~ /BenchmarkGoValidGT/) govalid_gt = $3
    if ($1 ~ /BenchmarkGoPlaygroundGT/) playground_gt = $3
    if ($1 ~ /BenchmarkGoValidatorGT/) asaskevich_gt = $3
    
    if ($1 ~ /BenchmarkGoValidMaxLength/) govalid_maxlength = $3
    if ($1 ~ /BenchmarkGoPlaygroundMaxLength/) playground_maxlength = $3
    if ($1 ~ /BenchmarkGoValidatorMaxLength/) asaskevich_maxlength = $3
    if ($1 ~ /BenchmarkGookitValidateMaxLength/) gookit_maxlength = $3
    
    if ($1 ~ /BenchmarkGoValidEnum/) govalid_enum = $3
}
END {
    # Calculate best competitor for each validator
    print "| Validator | govalid | Best Competitor | Improvement |"
    print "|-----------|---------|-----------------|-------------|"
    
    # Required
    best_req = playground_required
    best_req_name = "go-playground"
    if (asaskevich_required != "" && asaskevich_required < best_req) {
        best_req = asaskevich_required
        best_req_name = "asaskevich"
    }
    req_imp = int(best_req / govalid_required * 10) / 10
    print "| Required  | " govalid_required " | " best_req " (" best_req_name ") | **" req_imp "x faster** |"
    
    # Email
    best_email = 999999
    best_email_name = ""
    if (playground_email != "") {
        best_email = playground_email
        best_email_name = "go-playground"
    }
    if (asaskevich_email != "" && asaskevich_email < best_email) {
        best_email = asaskevich_email
        best_email_name = "asaskevich"
    }
    if (gookit_email != "" && gookit_email < best_email) {
        best_email = gookit_email
        best_email_name = "gookit"
    }
    email_imp = int(best_email / govalid_email * 10) / 10
    print "| Email     | " govalid_email " | " best_email " (" best_email_name ") | **" email_imp "x faster** |"
    
    # GT
    best_gt = playground_gt
    best_gt_name = "go-playground"
    if (asaskevich_gt != "" && asaskevich_gt < best_gt) {
        best_gt = asaskevich_gt
        best_gt_name = "asaskevich"
    }
    gt_imp = int(best_gt / govalid_gt * 10) / 10
    print "| GT/LT     | " govalid_gt " | " best_gt " (" best_gt_name ") | **" gt_imp "x faster** |"
    
    # MaxLength
    best_maxlen = 999999
    best_maxlen_name = ""
    if (playground_maxlength != "") {
        best_maxlen = playground_maxlength
        best_maxlen_name = "go-playground"
    }
    if (asaskevich_maxlength != "" && asaskevich_maxlength < best_maxlen) {
        best_maxlen = asaskevich_maxlength
        best_maxlen_name = "asaskevich"
    }
    if (gookit_maxlength != "" && gookit_maxlength < best_maxlen) {
        best_maxlen = gookit_maxlength
        best_maxlen_name = "gookit"
    }
    max_imp = int(best_maxlen / govalid_maxlength * 10) / 10
    print "| MaxLength | " govalid_maxlength " | " best_maxlen " (" best_maxlen_name ") | **" max_imp "x faster** |"
    
    # Enum (govalid exclusive)
    print "| Enum      | " govalid_enum " | N/A | **govalid exclusive** |"
}')

# Update the performance table in _index.md
# Create a temporary file with the new table
echo "$HOMEPAGE_TABLE" > /tmp/homepage_table.txt

# Replace the table section in _index.md
awk '
BEGIN { in_table = 0 }
/\| Validator.*\| govalid/ { in_table = 1; next }
/\| Enum/ && in_table { 
    # Print the new table content
    while ((getline line < "/tmp/homepage_table.txt") > 0) {
        print line
    }
    close("/tmp/homepage_table.txt")
    in_table = 0
    next
}
in_table { next }
{ print }
' docs/content/_index.md > /tmp/_index_new.md && mv /tmp/_index_new.md docs/content/_index.md

rm -f /tmp/homepage_table.txt

echo ""
echo "✅ Benchmark synchronization complete!"
echo ""
echo "📁 Updated files:"
echo "  - test/benchmark/README.md"
echo "  - docs/content/_index.md (performance table)"
echo ""
echo "🔍 Benchmark data updated from: $BENCH_DATE"
echo ""
echo "📊 Full benchmark results: test/benchmark/README.md"
