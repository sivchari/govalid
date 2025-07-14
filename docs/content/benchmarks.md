---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-14  
**Platform:** Darwin 24.5.0 arm64  
**Go version:** go1.25-devel_a8e99ab19c

## Raw Benchmark Data

```
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.09043 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8271 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09822 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	628736943	         1.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	611306106	         1.946 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	616978609	         1.887 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	651276344	         1.852 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	30768770	        37.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1873390	       641.6 ns/op	      90 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 2009972	       597.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16              	28449108	        42.24 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateEmail-16              	  115292	     10513 ns/op	   16046 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                      	539276762	         2.234 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	623069402	         1.944 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	18876345	        62.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	21448790	        54.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	620349398	         1.945 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19405530	        61.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	632340391	         1.935 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19819094	        61.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	625917411	         1.951 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19396447	        61.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	481143999	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	14702533	        79.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	78670040	        15.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16025070	        74.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7206669	       165.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6425402	       183.9 ns/op	     432 B/op	       4 allocs/op
BenchmarkGookitValidateMaxLength-16          	  113613	     10316 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16                  	430591298	         2.789 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	14961015	        80.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17607387	        68.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 7375346	       164.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-16                  	609726277	         1.954 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14052097	        85.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	624336777	         1.926 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16           	35281578	        34.62 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateRequired-16           	  115696	      9785 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                       	27739598	        42.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4208511	       287.4 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  157510	      7810 ns/op	     144 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16                	  152617	      7919 ns/op	     170 B/op	       2 allocs/op
BenchmarkGookitValidateURL-16                	  113732	     10629 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                      	32107666	        37.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4689171	       254.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6082100	       198.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16               	 4979631	       238.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateUUID-16               	  115017	     10716 ns/op	   15520 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 37.08ns | 641.6 | **17.3x faster** | 0 allocs/op | 5 allocs + 90 B/op |
| GT | 1.944ns | 62.76 | **32.3x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.945ns | 61.64 | **31.7x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.935ns | 61.14 | **31.6x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.951ns | 61.86 | **31.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.491ns | 79.98 | **32.1x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.46ns | 74.83 | **4.8x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.789ns | 80.17 | **28.7x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.59ns | 68.10 | **5.9x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.954ns | 85.45 | **43.7x faster** | 0 allocs/op | 0 allocs/op |
| URL | 42.28ns | 287.4 | **6.8x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 37.26ns | 254.1 | **6.8x faster** | 0 allocs/op | 0 allocs/op |
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
