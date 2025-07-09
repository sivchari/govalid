# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

Benchmarked on: 2025-07-09
Platform: Apple M3 Max (16 cores)
Go version: go1.24.3

```
BenchmarkGoValidEmail-16              	32228103	        38.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16         	 1868752	       644.1 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidEnum-16               	532029958	         2.242 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                 	611265494	         1.965 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16            	18765514	        64.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                	610501644	         1.972 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16           	18728052	        62.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                 	612178263	         1.966 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16            	18604290	        62.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                	611989772	         1.965 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16           	18566031	        63.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16           	473255866	         2.537 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16      	14922596	        80.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16          	75382410	        15.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16     	15913501	        74.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16           	386819032	         3.088 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16      	14729302	        81.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16          	96840576	        11.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16     	17009824	        70.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16           	612683832	         1.959 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16      	13983676	        86.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16               	36773295	        32.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16          	 4684831	       258.2 ns/op	       0 B/op	       0 allocs/op
```

## govalid-Specific Validators

These validators are unique to govalid and don't have direct equivalents in go-playground/validator:

```
BenchmarkGoValidEnum-16               	539464647	         2.164 ns/op	       0 B/op	       0 allocs/op
```

## Performance Summary

| Validator | GoValid (ns/op) | go-playground/validator (ns/op) | Improvement |
|-----------|-----------------|--------------------------------|-------------|
| Email     | 38.15           | 644.1                         | **16.9x faster** |
| GT        | 1.97            | 64.42                         | **32.7x faster** |
| GTE       | 1.97            | 62.61                         | **31.8x faster** |
| LT        | 1.97            | 62.53                         | **31.7x faster** |
| LTE       | 1.97            | 63.91                         | **32.5x faster** |
| MaxItems  | 2.54            | 80.74                         | **31.8x faster** |
| MaxLength | 15.75           | 74.93                         | **4.8x faster** |
| MinItems  | 3.09            | 81.00                         | **26.2x faster** |
| MinLength | 11.77           | 70.59                         | **6.0x faster** |
| Required  | 1.96            | 86.59                         | **44.2x faster** |
| UUID      | 32.73           | 258.2                         | **7.9x faster** |

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

1. **Exceptional Performance**: GoValid shows 4.8x to 44.2x performance improvements across all validators
2. **Sub-3ns Execution**: All collection, numeric, and enum validators execute in <3ns  
3. **Zero Allocations**: All GoValid validators perform zero heap allocations vs 0-5 allocs for playground
4. **Unicode Efficiency**: String length validators with Unicode support still 4.8-6.0x faster
5. **Email Optimization**: Email validation **16.9x faster** with zero allocations vs 89B/5 allocs (major improvement from 3x to 16.9x)
6. **UUID Validation**: UUID validation **7.9x faster** with zero allocations and optimal string parsing
7. **Extended Type Support**: Collection validators work with map/channel types unsupported by playground
8. **Enum Innovation**: First-class enum validation for multiple types with compile-time safety
9. **Helper Function Architecture**: Complex validators use optimized external helper functions for better performance

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
- **After**: 38.15 ns/op with optimized external helper (5.4x improvement)
- **Key optimizations**: 
  - Manual domain label parsing instead of `strings.Split` (eliminates 32B allocation)
  - Direct character-by-character dot detection instead of `strings.Contains`
  - Single-pass validation algorithm
  - Zero memory allocations
  - **Function decomposition**: Refactored into smaller, optimized functions for better compiler optimization

### UUID Validation Optimizations

The UUID validator is implemented using external helper functions for optimal performance:

- **Performance**: 32.73 ns/op with zero allocations
- **vs go-playground/validator**: 7.9x faster (258.2 ns/op)
- **Key features**:
  - RFC 4122 compliant validation (8-4-4-4-12 hex digits with hyphens)
  - Direct string indexing for hyphen position validation
  - Efficient hex character validation without regex
  - Version (1-5) and variant (8,9,A,B) field validation
  - **Function decomposition**: Modular design for better lint compliance and compiler optimization