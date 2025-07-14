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
BenchmarkGoValidEmail-16               	29599147	        36.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16          	 1888803	       630.4 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16           	 2063482	       584.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16        	30213571	        39.35 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateEmail-16        	  124010	      9559 ns/op	   15952 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                	538697865	         2.223 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                  	660667740	         1.825 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16             	18637350	        63.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16              	22722162	        53.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                 	624790245	         1.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16            	19602438	        61.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                  	655241349	         1.832 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16             	19689493	        61.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                 	621192230	         1.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16            	18723085	        63.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16            	484081939	         2.469 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16       	15240030	        79.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16           	75544552	        15.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16      	15868310	        74.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16       	 7721468	       155.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16    	 7747087	       159.3 ns/op	     432 B/op	       4 allocs/op
BenchmarkGookitValidateMaxLength-16    	  129084	      9125 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16            	425949052	         2.779 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16       	15269174	        79.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16           	100000000	        11.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16      	17384396	        69.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16       	 7493881	       159.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-16            	628318717	         1.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16       	13567683	        85.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16        	619585749	         1.929 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16     	34890398	        33.69 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateRequired-16     	  139293	      8698 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                 	29324109	        41.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16            	 4392483	       277.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16             	  155212	      7776 ns/op	     146 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16          	  155686	      7739 ns/op	     170 B/op	       2 allocs/op
BenchmarkGookitValidateURL-16          	  132864	      9216 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                	33677632	        36.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16           	 4768158	       253.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16            	 6204364	       193.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16         	 5191681	       231.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateUUID-16         	  136436	      9256 ns/op	   15514 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 36.80ns | 630.4 | **17.1x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 1.825ns | 63.35 | **34.7x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.914ns | 61.48 | **32.1x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.832ns | 61.59 | **33.6x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.925ns | 63.47 | **33.0x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.469ns | 79.18 | **32.1x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.58ns | 74.90 | **4.8x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.779ns | 79.15 | **28.5x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.46ns | 69.43 | **6.1x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.914ns | 85.82 | **44.8x faster** | 0 allocs/op | 0 allocs/op |
| URL | 41.68ns | 277.8 | **6.7x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 36.21ns | 253.0 | **7.0x faster** | 0 allocs/op | 0 allocs/op |
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
