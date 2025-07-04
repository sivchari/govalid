# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

Benchmarked on: 2025-07-02
Platform: Apple M3 Max (16 cores)
Go version: go1.24

```
BenchmarkGoValidGT-16                 	644065195	         1.882 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16            	19697977	        59.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                 	645251919	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16            	19474315	        60.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16           	485595762	         2.475 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16      	15206080	        77.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16          	78798972	        15.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16     	17434964	        69.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16           	434966996	         2.751 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16      	15503900	        76.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16          	100000000	        11.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16     	18134911	        65.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16           	638489334	         1.892 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16      	14232363	        83.66 ns/op	       0 B/op	       0 allocs/op
```

## Performance Summary

| Validator | GoValid (ns/op) | go-playground/validator (ns/op) | Improvement |
|-----------|-----------------|--------------------------------|-------------|
| GT        | 1.88            | 59.36                         | **32x faster** |
| LT        | 1.87            | 60.70                         | **32x faster** |
| MaxItems  | 2.48            | 77.82                         | **31x faster** |
| MaxLength | 15.24           | 69.01                         | **4.5x faster** |
| MinItems  | 2.75            | 76.50                         | **28x faster** |
| MinLength | 11.47           | 65.94                         | **5.7x faster** |
| Required  | 1.89            | 83.66                         | **44x faster** |

## Collection Validators (govalid-only)

These validators support map and channel types, which go-playground/validator doesn't support:

- **MaxItems**: slice, array, map, channel length ≤ limit
- **MinItems**: slice, array, map, channel length ≥ limit

## Key Findings

1. **Exceptional Performance**: GoValid shows 4.5x to 44x performance improvements across all validators
2. **Sub-3ns Execution**: All collection and numeric validators execute in <3ns
3. **Zero Allocations**: All GoValid validators perform zero heap allocations
4. **Unicode Efficiency**: String length validators with Unicode support still 4.5-5.7x faster
5. **Extended Type Support**: Collection validators work with map/channel types unsupported by playground

## Implementation Notes

- GoValid generates compile-time validation functions with zero runtime reflection
- Proper Unicode support in string length validators using `utf8.RuneCountInString`
- Interface-based import system for scalable package management
- Comprehensive type support including map and channel validation
- Kubernetes-compatible required field validation (empty slice vs nil distinction)