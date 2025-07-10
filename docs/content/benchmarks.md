---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-10  
**Platform:** Darwin 24.5.0 arm64  
**Go version:** go1.25-devel_a8e99ab19c

## Raw Benchmark Data

```
BenchmarkGoValidEmail-16             	31843813	        37.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16        	 1930374	       629.8 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidEnum-16              	548327136	         2.230 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                	614606246	         1.876 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16           	20117954	        59.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16               	649853220	         1.832 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16          	19230384	        61.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                	622414173	         1.913 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16           	20024850	        59.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16               	657234140	         1.838 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16          	19951810	        59.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16          	481109035	         2.497 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16     	15247614	        78.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16         	79215328	        15.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16    	16983074	        70.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16          	430797020	         2.787 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16     	15170478	        78.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16         	100000000	        11.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16    	18397171	        65.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16          	628808877	         1.935 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16     	14102072	        85.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16              	29724859	        36.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16         	 4872068	       248.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16               	29471809	        41.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16          	 4240443	       276.2 ns/op	     144 B/op	       1 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 37.61ns | 629.8 | **16.7x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 1.876ns | 59.96 | **32.0x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.832ns | 61.01 | **33.3x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.913ns | 59.85 | **31.3x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.838ns | 59.21 | **32.2x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.497ns | 78.67 | **31.5x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.23ns | 70.79 | **4.6x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.787ns | 78.60 | **28.2x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.47ns | 65.70 | **5.7x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.935ns | 85.67 | **44.3x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.49ns | 248.7 | **6.8x faster** | 0 allocs/op | 0 allocs/op |
| URL | 41.73ns | 276.2 | **6.6x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| Enum | 2.242ns | N/A (govalid exclusive) | **govalid exclusive** | 0 allocs/op | N/A |

## Performance Categories

### 🚀 Ultra-Fast (< 3ns)
- **Required**: ~1.9ns - 45x faster
- **GT/GTE/LT/LTE**: ~1.9ns - 32x faster
- **Enum**: ~2.2ns - govalid exclusive
- **MaxItems**: ~2.5ns - 32x faster
- **MinItems**: ~2.8ns - 29x faster

### ⚡ Fast (3-40ns)
- **MinLength**: ~11ns - 6x faster
- **MaxLength**: ~15ns - 5x faster
- **UUID**: ~36ns - 7x faster
- **URL**: ~41ns - 7x faster
- **Email**: ~36ns - 17x faster

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

### 4. String Length Performance
Unicode-aware string validators are 4.8-6.0x faster despite proper UTF-8 handling.

## govalid-Exclusive Features

### Enum Validation
```go
// +govalid:enum=admin,user,guest
Role string
```
- **Performance**: ~2.2ns with 0 allocations
- **No equivalent** in go-playground/validator
- Supports string, numeric, and custom types

### Extended Collection Support
```go
// +govalid:maxitems=10
Items map[string]int  // Maps supported!

// +govalid:minitems=1
Channel chan string   // Channels supported!
```

## Optimization Techniques

### 1. Code Generation
- **Compile-time validation functions** (no runtime reflection)
- **Inlined simple operations** for maximum speed
- **Direct field access** with no interface overhead

### 2. External Helper Functions
Complex validators use optimized external functions for better performance.

### 3. Manual String Parsing
- **Character-by-character parsing** instead of `strings.Split`
- **Direct indexing** instead of `strings.Contains`
- **Single-pass algorithms** for complex validation

### 4. Memory Optimization
- **Zero heap allocations** across all validators
- **Stack-only operations** for maximum cache efficiency
- **Minimal memory footprint** in generated code

## Running Benchmarks

To run benchmarks yourself:

```bash
# Sync all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem
```

## Conclusion

govalid delivers exceptional performance improvements:
- **4.8x to 45x faster** than go-playground/validator
- **Zero allocations** across all validators
- **Sub-3ns performance** for simple operations
- **Extended type support** (maps, channels, enums)
- **Production-ready** with comprehensive test coverage

Choose govalid when performance matters and zero allocations are critical for your application's success.
