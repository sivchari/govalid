# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

Benchmarked on: 2025-07-09
Platform: Apple M3 Max (16 cores)
Go version: go1.24.3

```
BenchmarkGoValidEmail-16              	26353546	        45.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16         	 1840674	       651.8 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidEnum-16               	534860061	         2.244 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                 	607874124	         1.977 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16            	18497275	        64.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                	609910148	         1.970 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16           	19079733	        62.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                 	613839652	         1.969 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16            	17733880	        63.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                	610599888	         1.966 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16           	18782808	        63.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16           	477820183	         2.524 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16      	14864668	        81.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16          	75906330	        15.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16     	16083867	        74.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16           	427818543	         2.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16      	14967289	        81.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16          	96739086	        12.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16     	17179126	        69.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16           	605284765	         1.966 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16      	13964752	        86.74 ns/op	       0 B/op	       0 allocs/op
```

## govalid-Specific Validators

These validators are unique to govalid and don't have direct equivalents in go-playground/validator:

```
BenchmarkGoValidEnum-16               	539464647	         2.164 ns/op	       0 B/op	       0 allocs/op
```

## Performance Summary

| Validator | GoValid (ns/op) | go-playground/validator (ns/op) | Improvement |
|-----------|-----------------|--------------------------------|-------------|
| Email     | 45.08           | 651.8                         | **14.5x faster** |
| GT        | 1.98            | 64.33                         | **32.5x faster** |
| GTE       | 1.97            | 62.81                         | **31.9x faster** |
| LT        | 1.97            | 63.23                         | **32.1x faster** |
| LTE       | 1.97            | 63.67                         | **32.3x faster** |
| MaxItems  | 2.52            | 81.28                         | **32.3x faster** |
| MaxLength | 15.90           | 74.75                         | **4.7x faster** |
| MinItems  | 2.80            | 81.11                         | **29.0x faster** |
| MinLength | 12.11           | 69.28                         | **5.7x faster** |
| Required  | 1.97            | 86.74                         | **44.0x faster** |

## govalid-Exclusive Features

### Enum Validation
- **Enum**: Comprehensive enum validation for string, numeric, and custom types (~2.16ns)
  - Supports multiple value comparisons in a single operation
  - Works with custom type definitions (e.g., `type Status string`)
  - Zero-allocation enum checking with compile-time safety

### Collection Type Extension
These validators support map and channel types, which go-playground/validator doesn't support:

- **MaxItems**: slice, array, map, channel length ≤ limit  
- **MinItems**: slice, array, map, channel length ≥ limit

## Key Findings

1. **Exceptional Performance**: GoValid shows 4.7x to 44x performance improvements across all validators
2. **Sub-3ns Execution**: All collection, numeric, and enum validators execute in <3ns  
3. **Zero Allocations**: All GoValid validators perform zero heap allocations vs 0-5 allocs for playground
4. **Unicode Efficiency**: String length validators with Unicode support still 4.7-5.7x faster
5. **Email Optimization**: Email validation **14.5x faster** with zero allocations vs 88B/5 allocs (major improvement from 3x to 14.5x)
6. **Extended Type Support**: Collection validators work with map/channel types unsupported by playground
7. **Enum Innovation**: First-class enum validation for multiple types with compile-time safety
8. **Helper Function Architecture**: Complex validators use optimized external helper functions for better performance

## Implementation Notes

- GoValid generates compile-time validation functions with zero runtime reflection
- **External Helper Functions**: Complex validators (email) use optimized external functions in `validation/validationhelper`
- **Zero-Allocation Email Validation**: Manual string parsing eliminates `strings.Split` and `strings.Contains` allocations
- Proper Unicode support in string length validators using `utf8.RuneCountInString`
- Interface-based import system for scalable package management
- Comprehensive type support including map and channel validation
- Kubernetes-compatible required field validation (empty slice vs nil distinction)

### Email Validation Optimizations

The email validator underwent significant optimization using external helper functions:

- **Before**: 204.8 ns/op with inline function generation
- **After**: 45.08 ns/op with optimized external helper (4.5x improvement)
- **Key optimizations**: 
  - Manual domain label parsing instead of `strings.Split` (eliminates 32B allocation)
  - Direct character-by-character dot detection instead of `strings.Contains`
  - Single-pass validation algorithm
  - Zero memory allocations