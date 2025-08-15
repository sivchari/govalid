# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-15  
**Platform:** Darwin 24.6.0 arm64  
**Go version:** go1.26-devel_6fbad4be75

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-16                     	168865384	         7.193 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-16                	 4978832	       242.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-16       	19626148	        61.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-16              	  102858	     11927 ns/op	   16834 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.2022 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	454332843	         2.563 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	453914157	         2.573 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	650905266	         1.851 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	684814404	         1.766 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	664789868	         1.821 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	36152426	        34.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	 2009169	       597.8 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 2049222	       586.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-16              	  118150	     10282 ns/op	   16091 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                      	447943938	         2.661 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	668224464	         1.808 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	21037544	        56.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	23517595	        51.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	669875866	         1.872 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19673769	        58.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-16                    	276112791	         4.391 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-16               	21888705	        55.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-16                	 9073251	       132.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-16             	  130785	      9630 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-16                        	658684084	         1.820 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	21202818	        56.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	654405499	         1.839 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	21277569	        56.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	374376817	         3.212 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15977452	        75.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	84178035	        14.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	17257351	        68.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 8072121	       148.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-16          	  131133	      9603 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16                  	347334909	         3.462 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15847144	        75.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        10.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	19626697	        61.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 8083320	       149.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-16                   	214583640	         5.559 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-16              	27490568	        44.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-16               	16518866	        72.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-16            	  124155	      9841 ns/op	   15550 B/op	      76 allocs/op
BenchmarkGoValidRequired-16                  	565022959	         2.112 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14834302	        81.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	688114657	         1.748 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-16           	  133226	      9390 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                       	27790670	        41.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4499622	       267.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  159542	      7569 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-16                	  122894	      9873 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                      	35398273	        33.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4964876	       242.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6328350	       187.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-16               	  127159	      9953 ns/op	   15518 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| UUID | 33.76 / 0 allocs | 242.0 / 0 allocs | **7.2x** | 187.9 / 0 allocs | **5.6x** | 9953 / 15518 B / 74 allocs | **294.8x** |
| URL | 41.72 / 0 allocs | 267.5 / 144 B / 1 allocs | **6.4x** | 7569 / 147 B / 1 allocs | **181.4x** | 9873 / 15641 B / 75 allocs | **236.6x** |
| GTE | 1.872 / 0 allocs | 58.50 / 0 allocs | **31.2x** | N/A | N/A | N/A | N/A |
| MinLength | 10.84 / 0 allocs | 61.61 / 0 allocs | **5.7x** | 149.6 / 32 B / 2 allocs | **13.8x** | N/A | N/A |
| LTE | 1.839 / 0 allocs | 56.52 / 0 allocs | **30.7x** | N/A | N/A | N/A | N/A |
| MinItems | 3.462 / 0 allocs | 75.80 / 0 allocs | **21.9x** | N/A | N/A | N/A | N/A |
| MaxItems | 3.212 / 0 allocs | 75.40 / 0 allocs | **23.5x** | N/A | N/A | N/A | N/A |
| Enum | 2.661 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Required | 2.112 / 0 allocs | 81.11 / 0 allocs | **38.4x** | 1.748 / 0 allocs | **0.8x** | 9390 / 15472 B / 72 allocs | **4446.0x** |
| MaxLength | 14.29 / 0 allocs | 68.96 / 0 allocs | **4.8x** | 148.5 / 32 B / 2 allocs | **10.4x** | 9603 / 15632 B / 80 allocs | **672.0x** |
| GT | 1.808 / 0 allocs | 56.93 / 0 allocs | **31.5x** | 51.58 / 0 allocs | **28.5x** | N/A | N/A |
| Numeric | 5.559 / 0 allocs | 44.15 / 0 allocs | **7.9x** | 72.80 / 0 allocs | **13.1x** | 9841 / 15550 B / 76 allocs | **1770.3x** |
| LT | 1.820 / 0 allocs | 56.70 / 0 allocs | **31.2x** | N/A | N/A | N/A | N/A |
| Length | 4.391 / 0 allocs | 55.23 / 0 allocs | **12.6x** | 132.7 / 32 B / 2 allocs | **30.2x** | 9630 / 15600 B / 78 allocs | **2193.1x** |
| Alpha | 7.193 / 0 allocs | 242.8 / 0 allocs | **33.8x** | 61.75 / 0 allocs | **8.6x** | 11927 / 16834 B / 97 allocs | **1658.1x** |
| Email | 34.10 / 0 allocs | 597.8 / 89 B / 5 allocs | **17.5x** | 586.0 / 0 allocs | **17.2x** | 10282 / 16091 B / 74 allocs | **301.5x** |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 0.2022 | 0 allocs |
| CELMultipleExpressions | 2.563 | 0 allocs |
| CELBasic | 2.573 | 0 allocs |
| CELCrossField | 1.851 | 0 allocs |
| CELStringLength | 1.766 | 0 allocs |
| CELNumericComparison | 1.821 | 0 allocs |

CEL (Common Expression Language) support allows complex runtime expressions with near-zero overhead.

## govalid Exclusive Features

### Enum Validation
- Comprehensive enum validation for string, numeric, and custom types
- Zero-allocation enum checking with compile-time safety
- Works with custom type definitions (e.g., `type Status string`)

### Collection Type Extension
These validators support map and channel types beyond standard slices:
- **MaxItems**: slice, array, map, channel length validation
- **MinItems**: slice, array, map, channel length validation

### CEL Expression Support
- Complex runtime validation expressions
- Cross-field validation
- Custom business logic without code generation

## Implementation Notes

### Zero-Allocation Design
- All govalid validators perform **zero heap allocations**
- Manual string parsing eliminates memory overhead
- Stack-only operations for maximum cache efficiency

### Performance Optimizations
- Compile-time code generation (no runtime reflection)
- Inlined validation functions for simple checks
- External helper functions for complex validators
- Single-pass algorithms for string operations

### Unicode Support
- Proper UTF-8 handling with `utf8.RuneCountInString`
- Consistent behavior with go-playground/validator for international text

## Running Benchmarks

```bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=10s

# Run specific validator benchmarks
go test ./benchmark/ -bench=BenchmarkGoValid{ValidatorName} -benchmem
```
