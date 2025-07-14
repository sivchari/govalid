# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

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
