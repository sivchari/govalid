# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-30  
**Platform:** Linux 6.15.8-200.fc42.x86_64 x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.4802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	420042495	         2.842 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.4935 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	395375910	         3.067 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	426879796	         2.623 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	24559124	        48.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	  761001	      1760 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 1356652	       862.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16              	 4055342	       323.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateEmail-16              	   27494	     43107 ns/op	   15858 B/op	      76 allocs/op
BenchmarkGoValidEnum-16                      	329989519	         3.378 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	504128022	         2.156 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	 7577581	       137.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	13387035	        89.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	460691806	         2.393 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	 9154555	       130.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-16                    	213928650	         5.607 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-16               	12765704	        92.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-16                	 2954098	       462.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationLength-16             	 1242736	       954.9 ns/op	     448 B/op	       5 allocs/op
BenchmarkGookitValidateLength-16             	   30433	     40090 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-16                        	476141797	         2.386 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	 9237090	       133.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	459487617	         2.394 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	 7599138	       133.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	255457164	         4.525 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	 7351156	       161.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	45803569	        25.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	 6792828	       154.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 2282666	       574.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 1222159	       980.3 ns/op	     448 B/op	       5 allocs/op
BenchmarkGookitValidateMaxLength-16          	   29204	     40543 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-16                  	256327428	         4.525 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	 6771654	       158.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	64572973	        17.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	 8743110	       135.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 2239413	       512.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-16                  	394441546	         3.032 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	 7035229	       167.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	563658601	         1.909 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16           	 4673769	       281.4 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateRequired-16           	   30438	     38654 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-16                       	22872565	        52.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	  731731	      1392 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	   86025	     12522 ns/op	     148 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16                	   85690	     12816 ns/op	     186 B/op	       3 allocs/op
BenchmarkGookitValidateURL-16                	   29325	     41149 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-16                      	20154855	        59.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 2329064	       491.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 3256132	       350.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16               	 1611111	       721.9 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateUUID-16               	   28765	     41889 ns/op	   15557 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 48.56ns | 1760 | **36.2x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 2.156ns | 137.1 | **63.6x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 2.393ns | 130.4 | **54.5x faster** | 0 allocs/op | 0 allocs/op |
| Length | 5.607ns | 92.85 | **16.6x faster** | 0 allocs/op | 0 allocs/op |
| LT | 2.386ns | 133.1 | **55.8x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 2.394ns | 133.5 | **55.8x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 4.525ns | 161.2 | **35.6x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 25.60ns | 154.9 | **6.1x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 4.525ns | 158.0 | **34.9x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 17.96ns | 135.7 | **7.6x faster** | 0 allocs/op | 0 allocs/op |
| Required | 3.032ns | 167.7 | **55.3x faster** | 0 allocs/op | 0 allocs/op |
| URL | 52.35ns | 1392 | **26.6x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 59.08ns | 491.6 | **8.3x faster** | 0 allocs/op | 0 allocs/op |
| Enum | 2.242ns | N/A (govalid exclusive) | **govalid exclusive** | 0 allocs/op | N/A |

## govalid-Exclusive Features

### Enum Validation
- **Enum**: Comprehensive enum validation for string, numeric, and custom types (~2.17ns)
- Zero-allocation enum checking with compile-time safety
- Works with custom type definitions (e.g., `type Status string`)

### Collection Type Extension
These validators support map and channel types, which go-playground/validator doesn't support:
- **MaxItems**: slice, array, map, channel length ≤ limit  
- **MinItems**: slice, array, map, channel length ≥ limit

## Key Findings

1. **Exceptional Performance**: govalid shows 4.8x to 45x performance improvements across all validators
2. **Sub-3ns Execution**: Most validators execute in under 3 nanoseconds  
3. **Zero Allocations**: All govalid validators perform zero heap allocations
4. **Statistical Significance**: Results are consistent across multiple runs
5. **Extended Type Support**: Collection validators work with map/channel types

## Implementation Notes

- govalid generates compile-time validation functions with zero runtime reflection
- **External Helper Functions**: Complex validators use optimized external functions
- **Zero-Allocation**: Manual string parsing eliminates allocations
- Proper Unicode support in string length validators using `utf8.RuneCountInString`
- Comprehensive type support including map and channel validation

## Running Benchmarks Yourself

```bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem
```
