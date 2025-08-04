# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-08-04  
**Platform:** Linux 6.15.8-200.fc42.x86_64 x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-16                     	141274897	         8.547 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-16                	 2462288	       487.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-16       	11147374	       107.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-16              	561660699	         2.076 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.7505 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	358447099	         3.372 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	342246789	         3.419 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	378614107	         2.975 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	1000000000	         1.038 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	24549442	        48.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	  823983	      1508 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 1353092	       885.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-16              	   34730	     35025 ns/op	   15852 B/op	      76 allocs/op
BenchmarkGoValidEnum-16                      	339558291	         3.440 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	442100852	         2.470 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	 9040794	       131.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	13248624	        90.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	544875714	         2.193 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	 7976847	       134.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-16                    	162685860	         7.420 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-16               	 8675176	       131.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-16                	 3211560	       417.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-16             	   33007	     35430 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-16                        	536548378	         2.186 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	 7533094	       138.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	505882672	         2.226 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	 7872116	       132.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	238008087	         4.998 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	 6599076	       166.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	40753048	        26.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	 7382617	       144.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 2959848	       430.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-16          	   37117	     34153 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-16                  	209500981	         5.658 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	 6745190	       166.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	59803000	        19.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	 7750194	       149.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 2618836	       454.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-16                   	183533002	         6.551 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-16              	11707054	       100.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-16               	 8068819	       127.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-16            	   34864	     35200 ns/op	   15588 B/op	      78 allocs/op
BenchmarkGoValidRequired-16                  	388635582	         2.974 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	 6182602	       175.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	558555505	         1.968 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-16           	   40160	     31421 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-16                       	18690115	        64.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 1080440	      1075 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	   97249	     12938 ns/op	     148 B/op	       1 allocs/op
BenchmarkGookitValidateURL-16                	   36932	     33867 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-16                      	21317632	        57.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 2254420	       507.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 3348686	       363.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-16               	   34512	     32870 ns/op	   15556 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Alpha | 8.547ns | 487.9 | **57.1x faster** | 0 allocs/op | 0 allocs/op |
| Email | 48.28ns | 1508 | **31.2x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 2.470ns | 131.9 | **53.4x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 2.193ns | 134.0 | **61.1x faster** | 0 allocs/op | 0 allocs/op |
| Length | 7.420ns | 131.7 | **17.7x faster** | 0 allocs/op | 0 allocs/op |
| LT | 2.186ns | 138.8 | **63.5x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 2.226ns | 132.6 | **59.6x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 4.998ns | 166.4 | **33.3x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 26.88ns | 144.8 | **5.4x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 5.658ns | 166.9 | **29.5x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 19.67ns | 149.8 | **7.6x faster** | 0 allocs/op | 0 allocs/op |
| Numeric | 6.551ns | 100.4 | **15.3x faster** | 0 allocs/op | 0 allocs/op |
| Required | 2.974ns | 175.2 | **58.9x faster** | 0 allocs/op | 0 allocs/op |
| URL | 64.01ns | 1075 | **16.8x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 57.09ns | 507.0 | **8.9x faster** | 0 allocs/op | 0 allocs/op |
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
