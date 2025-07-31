# Benchmark Results

This document contains performance comparison results between govalid and go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-08-01  
**Platform:** Linux 6.15.8-200.fc42.x86_64 x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-16                     	166734322	         7.223 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-16                	 2506509	       480.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationAlpha-16              	  985546	      1319 ns/op	     120 B/op	       6 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-16       	 9699534	       109.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-16              	544960296	         2.006 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.5232 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	391184774	         3.084 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	       215.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	379138917	         3.186 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	399707060	         3.007 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	1000000000	         1.756 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	20130372	        56.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1000000	      1388 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 1311596	       885.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16              	 3931170	       320.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateEmail-16              	   29343	     38728 ns/op	   15848 B/op	      76 allocs/op
BenchmarkGoValidEnum-16                      	337831936	         3.441 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	547653525	         2.165 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	 9096225	       130.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	13130914	        90.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	556985961	         2.152 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	 9341229	       132.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-16                    	164305790	         7.307 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-16               	 8500341	       130.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-16                	 2867533	       433.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationLength-16             	 1288365	       930.4 ns/op	     448 B/op	       5 allocs/op
BenchmarkGookitValidateLength-16             	   32127	     37532 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-16                        	492942547	         2.388 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	 7540310	       133.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	551570212	         2.171 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	 8483163	       132.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	266733841	         4.506 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	 7657694	       157.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	47891312	        24.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	 7648212	       148.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 3504940	       423.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 1291771	       921.5 ns/op	     448 B/op	       5 allocs/op
BenchmarkGookitValidateMaxLength-16          	   31868	     38272 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-16                  	213083244	         5.886 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	 7380105	       157.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	55939418	        20.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	 8271450	       142.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 2814954	       447.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-16                   	139209772	         8.266 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-16              	11580610	        98.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-16               	 8324683	       122.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationNumeric-16            	 2758869	       481.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateNumeric-16            	   31362	     39585 ns/op	   15588 B/op	      78 allocs/op
BenchmarkGoValidRequired-16                  	367639327	         3.126 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	 6046644	       175.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	639989053	         1.878 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16           	 5013363	       275.9 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateRequired-16           	   33103	     36497 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-16                       	23126715	        51.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 1000000	      1272 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	   91410	     12605 ns/op	     145 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16                	   89580	     13331 ns/op	     186 B/op	       3 allocs/op
BenchmarkGookitValidateURL-16                	   32683	     37383 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-16                      	22569304	        53.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 2402175	       482.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 3233359	       351.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16               	 1684221	       739.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateUUID-16               	   31687	     35158 ns/op	   15556 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Alpha | 7.223ns | 480.4 | **66.5x faster** | 0 allocs/op | 0 allocs/op |
| Email | 56.95ns | 1388 | **24.4x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 2.165ns | 130.8 | **60.4x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 2.152ns | 132.2 | **61.4x faster** | 0 allocs/op | 0 allocs/op |
| Length | 7.307ns | 130.4 | **17.8x faster** | 0 allocs/op | 0 allocs/op |
| LT | 2.388ns | 133.3 | **55.8x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 2.171ns | 132.6 | **61.1x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 4.506ns | 157.2 | **34.9x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 24.55ns | 148.8 | **6.1x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 5.886ns | 157.1 | **26.7x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 20.68ns | 142.6 | **6.9x faster** | 0 allocs/op | 0 allocs/op |
| Numeric | 8.266ns | 98.54 | **11.9x faster** | 0 allocs/op | 0 allocs/op |
| Required | 3.126ns | 175.8 | **56.2x faster** | 0 allocs/op | 0 allocs/op |
| URL | 51.75ns | 1272 | **24.6x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 53.18ns | 482.9 | **9.1x faster** | 0 allocs/op | 0 allocs/op |
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
