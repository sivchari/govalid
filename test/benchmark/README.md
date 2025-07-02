# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

Benchmarked on: 2024-06-30
Platform: Apple M3 Max (16 cores)
Go version: go1.23

```
BenchmarkGoValidGT-16                     	1000000000	         0.5251 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                	20238277	        59.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                     	1000000000	         0.5187 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                	20031939	        59.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16              	88987215	        13.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16         	17163094	        70.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16               	1000000000	         1.013 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16          	14326383	        81.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16               	574773432	         1.907 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16          	20432972	        58.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16              	224338778	         5.256 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16         	19335181	        61.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                    	565656739	         2.127 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16               	18250533	        66.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                    	574773432	         2.063 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16               	18777334	        65.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLengthError-16         	48645196	        24.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLengthError-16    	 6990411	       182.0 ns/op	     216 B/op	       5 allocs/op
```

## GoValid-Exclusive Validators

Some validators are unique to GoValid and don't have direct equivalents in go-playground/validator:

```
BenchmarkGoValidEnum-16                   	587095758	         2.063 ns/op	       0 B/op	       0 allocs/op
```

## Performance Summary

| Validator | GoValid (ns/op) | go-playground/validator (ns/op) | Improvement |
|-----------|-----------------|--------------------------------|-------------|
| GT        | 0.53            | 59.99                         | **113x faster** |
| LT        | 0.52            | 59.64                         | **115x faster** |
| MaxLength | 13.48           | 70.16                         | **5.2x faster** |
| Required  | 1.01            | 81.89                         | **81x faster** |
| MaxItems  | 1.91            | 58.37                         | **31x faster** |
| MinLength | 5.26            | 61.87                         | **12x faster** |
| GTE       | 2.13            | 66.01                         | **31x faster** |
| LTE       | 2.06            | 65.50                         | **32x faster** |
| MaxLength (Error) | 24.79     | 182.0                      | **7.3x faster** |

### GoValid-Exclusive Performance

| Validator | GoValid (ns/op) | Description |
|-----------|-----------------|-------------|
| Enum      | 2.06            | **Enumeration validation** |

## Key Findings

1. **Exceptional Performance**: GoValid shows 5x to 115x performance improvements across all validators
2. **Unique Features**: GoValid provides enum validation with ~2ns performance
3. **Sub-nanosecond Execution**: GT/LT validators execute in ~0.5ns, Required in ~1ns
4. **Zero Allocations**: All GoValid validators perform zero heap allocations (vs 5 allocs for playground errors)
5. **Unicode Efficiency**: MaxLength with Unicode support still 5x faster than playground
6. **Error Handling**: Even error cases are 7x faster with zero allocations

## Implementation Notes

- GoValid generates compile-time validation functions with zero runtime reflection
- Proper Unicode support in MaxLength validator using rune counting
- Interface-based import system for scalable package management
- Kubernetes-compatible required field validation (empty slice vs nil distinction)