---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-11  
**Platform:** Darwin 24.5.0 arm64  
**Go version:** go1.25-devel_a8e99ab19c

## Raw Benchmark Data

```
BenchmarkGoValidEmail-16               	32555283	        36.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16          	 1936972	       622.2 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16           	 2032665	       589.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16        	31072719	        39.77 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateEmail-16        	  125805	      9594 ns/op	   15937 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                	540628920	         2.241 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                  	645607456	         1.840 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16             	19408393	        62.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16              	21097849	        54.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                 	615522447	         1.847 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16            	19576056	        61.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                  	654769564	         1.855 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16             	19214833	        62.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                 	624618152	         1.878 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16            	17564595	        63.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16            	488668489	         2.473 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16       	15210948	        79.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16           	76613064	        15.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16      	16315606	        73.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16       	 7689249	       155.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16    	 7680783	       160.7 ns/op	     432 B/op	       4 allocs/op
BenchmarkGookitValidateMaxLength-16    	  135312	      9132 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16            	434481410	         2.759 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16       	15326298	        79.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16           	100000000	        11.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16      	16812727	        70.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16       	 7675908	       156.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-16            	630665175	         1.886 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16       	14016931	        84.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16        	614499763	         1.940 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16     	35003366	        33.33 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateRequired-16     	  140805	      8671 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                 	28968247	        41.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16            	 4370895	       281.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16             	  154573	      7818 ns/op	     145 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16          	  152456	      7844 ns/op	     171 B/op	       2 allocs/op
BenchmarkGookitValidateURL-16          	  138303	      9011 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                	32390448	        36.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16           	 4753230	       251.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16            	 6205647	       195.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16         	 5079465	       235.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateUUID-16         	  127963	      9241 ns/op	   15514 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 36.80ns | 622.2 | **16.9x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 1.840ns | 62.71 | **34.1x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.847ns | 61.32 | **33.2x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.855ns | 62.41 | **33.6x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.878ns | 63.26 | **33.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.473ns | 79.47 | **32.1x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.51ns | 73.98 | **4.8x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.759ns | 79.28 | **28.7x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.59ns | 70.98 | **6.1x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.886ns | 84.82 | **45.0x faster** | 0 allocs/op | 0 allocs/op |
| URL | 41.90ns | 281.8 | **6.7x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 36.82ns | 251.5 | **6.8x faster** | 0 allocs/op | 0 allocs/op |
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
