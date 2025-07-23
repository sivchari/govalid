---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-23  
**Platform:** Linux 6.14.0-24-generic x86_64  
**Go version:** go1.24.5

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-2                    	164679214	         7.255 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-2               	 3490214	       331.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationAlpha-2             	 2516132	       477.3 ns/op	     120 B/op	       6 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-2      	14315827	        79.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-2             	743583081	         1.707 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-2            	1000000000	         0.9764 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-2   	1000000000	         0.9613 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-2    	1000000000	         0.7782 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-2                 	901600953	         1.611 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-2            	1000000000	         1.616 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-2          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-2     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-2                    	26108822	        47.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-2               	 1254596	       930.4 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-2                	 1850910	       638.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-2             	13999159	        87.90 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateEmail-2             	   73690	     15720 ns/op	   15586 B/op	      76 allocs/op
BenchmarkGoValidEnum-2                     	542174090	         2.234 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-2                       	1000000000	         1.199 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-2                  	13928852	        83.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-2                   	19900410	        62.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-2                      	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-2                 	14896790	        84.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-2                       	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-2                  	14765770	        81.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-2                      	1000000000	         1.031 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-2                 	14025458	        84.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-2                 	429054247	         2.941 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-2            	10227320	       102.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-2                	55182356	        22.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-2           	12008785	       101.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-2            	 5319528	       226.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-2         	 3160599	       377.3 ns/op	     448 B/op	       5 allocs/op
BenchmarkGookitValidateMaxLength-2         	   80406	     15361 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-2                 	385985637	         3.094 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-2            	11784582	        99.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-2                	84465284	        15.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-2           	12688219	        96.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-2            	 4933124	       231.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-2                 	854728023	         1.403 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-2            	11203710	       111.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-2             	996787096	         1.174 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-2          	16377784	        75.77 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateRequired-2          	   81187	     14546 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-2                      	23008521	        52.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-2                 	 2862736	       397.3 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-2                  	  124737	      9494 ns/op	     144 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-2               	  123141	      9101 ns/op	     184 B/op	       3 allocs/op
BenchmarkGookitValidateURL-2               	   77670	     15216 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-2                     	37147503	        31.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-2                	 3499276	       356.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-2                 	 4639436	       259.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-2              	 3319221	       356.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateUUID-2              	   74491	     15303 ns/op	   15539 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Alpha | 7.255ns | 331.2 | **45.7x faster** | 0 allocs/op | 0 allocs/op |
| Email | 47.49ns | 930.4 | **19.6x faster** | 0 allocs/op | 5 allocs + 88 B/op |
| GT | 1.199ns | 83.82 | **69.9x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.000ns | 84.35 | **84.3x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.000ns | 81.06 | **81.1x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.031ns | 84.62 | **82.1x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.941ns | 102.9 | **35.0x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 22.77ns | 101.0 | **4.4x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 3.094ns | 99.92 | **32.3x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 15.52ns | 96.09 | **6.2x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.403ns | 111.6 | **79.5x faster** | 0 allocs/op | 0 allocs/op |
| URL | 52.09ns | 397.3 | **7.6x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 31.19ns | 356.5 | **11.4x faster** | 0 allocs/op | 0 allocs/op |
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
