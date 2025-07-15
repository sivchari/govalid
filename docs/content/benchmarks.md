---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-15  
**Platform:** Darwin 24.5.0 arm64  
**Go version:** go1.25-devel_a8e99ab19c

## Raw Benchmark Data

```
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.09323 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8409 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.1015 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	618011162	         1.946 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	617148498	         1.953 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	605757344	         1.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	581620783	         2.036 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	31846465	        38.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1840472	       649.4 ns/op	      90 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 1976113	       606.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16              	27997833	        44.04 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateEmail-16              	  121108	     10481 ns/op	   16065 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                      	539168131	         2.231 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	609764490	         1.936 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19207836	        63.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	21764314	        54.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	610084184	         1.946 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19815618	        60.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	626116629	         1.916 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19840449	        60.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	625070131	         1.969 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19450260	        61.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	479873872	         2.520 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	14759504	        80.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	77328928	        15.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16308603	        73.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7367277	       161.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6660802	       183.4 ns/op	     432 B/op	       4 allocs/op
BenchmarkGookitValidateMaxLength-16          	  115408	     10276 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16                  	428160852	         2.785 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	14833287	        79.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17671692	        67.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 7323378	       165.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-16                  	623661589	         1.935 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	13981681	        85.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	625855254	         1.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16           	33811228	        34.91 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateRequired-16           	  117372	     10089 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                       	27992418	        42.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4223242	       287.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  149181	      7887 ns/op	     144 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16                	  153288	      7825 ns/op	     171 B/op	       2 allocs/op
BenchmarkGookitValidateURL-16                	  113994	      9850 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                      	32124428	        37.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4753952	       254.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6066202	       197.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16               	 5034348	       238.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateUUID-16               	  113701	     10603 ns/op	   15519 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 38.24ns | 649.4 | **17.0x faster** | 0 allocs/op | 5 allocs + 90 B/op |
| GT | 1.936ns | 63.00 | **32.5x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.946ns | 60.30 | **31.0x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.916ns | 60.77 | **31.7x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.969ns | 61.91 | **31.4x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.520ns | 80.30 | **31.9x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.67ns | 73.50 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.785ns | 79.81 | **28.7x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.60ns | 67.51 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.935ns | 85.51 | **44.2x faster** | 0 allocs/op | 0 allocs/op |
| URL | 42.54ns | 287.9 | **6.8x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 37.69ns | 254.0 | **6.7x faster** | 0 allocs/op | 0 allocs/op |
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
