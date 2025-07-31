# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-30  
**Platform:** Linux 6.15.8-200.fc42.x86_64 x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.09323 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8409 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.1015 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	618011162	         1.946 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	617148498	         1.953 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	605757344	         1.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	581620783	         2.036 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	31846465	        38.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1840472	       649.4 ns/op	      90 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 1976113	       606.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16              	27997833	        44.04 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateEmail-16              	  121108	     10481 ns/op	   16065 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                      	539168131	         2.231 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	609764490	         1.936 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19207836	        63.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	21764314	        54.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	610084184	         1.946 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19815618	        60.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	626116629	         1.916 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19840449	        60.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	625070131	         1.969 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19450260	        61.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	479873872	         2.520 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	14759504	        80.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	77328928	        15.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16308603	        73.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7367277	       161.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6660802	       183.4 ns/op	     432 B/op	       4 allocs/op
BenchmarkGookitValidateMaxLength-16          	  115408	     10276 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16                  	428160852	         2.785 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	14833287	        79.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17671692	        67.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 7323378	       165.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-16                  	623661589	         1.935 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	13981681	        85.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	625855254	         1.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16           	33811228	        34.91 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateRequired-16           	  117372	     10089 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                       	27992418	        42.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4223242	       287.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  149181	      7887 ns/op	     144 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16                	  153288	      7825 ns/op	     171 B/op	       2 allocs/op
BenchmarkGookitValidateURL-16                	  113994	      9850 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                      	32124428	        37.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4753952	       254.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6066202	       197.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16               	 5034348	       238.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateUUID-16               	  113701	     10603 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGoValidNumeric-12                      351653084        3.438 ns/op           0 B/op          0 allocs/op
BenchmarkGoPlaygroundNumeric-12                 18947991         61.60 ns/op           0 B/op          0 allocs/op
BenchmarkGoValidatorNumeric-12                  13556292         87.01 ns/op           0 B/op          0 allocs/op
BenchmarkOzzoValidationNumeric-12                7795245         149.2 ns/op          40 B/op          2 allocs/op
BenchmarkGookitValidateNumeric-12                 107311         11131 ns/op       15585 B/op         78 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 38.24ns | 649.4 | **17.0x faster** | 0 allocs/op | 5 allocs + 90 B/op |
| GT | 1.936ns | 63.00 | **32.5x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.946ns | 60.30 | **31.0x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.916ns | 60.77 | **31.7x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.969ns | 61.91 | **31.4x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.520ns | 80.30 | **31.9x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.67ns | 73.50 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.785ns | 79.81 | **28.7x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.60ns | 67.51 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.935ns | 85.51 | **44.2x faster** | 0 allocs/op | 0 allocs/op |
| URL | 42.54ns | 287.9 | **6.8x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 37.69ns | 254.0 | **6.7x faster** | 0 allocs/op | 0 allocs/op |
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
