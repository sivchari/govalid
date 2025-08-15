# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-15  
**Platform:** Darwin 24.6.0 arm64  
**Go version:** go1.26-devel_6fbad4be75

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-16                     	168089846	         7.135 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-16                	 4900998	       244.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-16       	19893816	        61.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-16              	  103581	     11924 ns/op	   16834 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.1996 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	461211843	         2.540 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	465882668	         2.526 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	643609215	         1.853 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	681164754	         1.760 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	687036831	         1.790 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	33266722	        33.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1992195	       603.1 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 2076205	       580.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-16              	  119406	     10223 ns/op	   16087 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                      	450452634	         2.655 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	652665480	         1.815 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	20244009	        59.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	22657219	        52.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	655356458	         1.832 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	20097934	        59.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-16                    	270973284	         4.398 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-16               	20548970	        57.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-16                	 9129194	       132.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-16             	  128559	      9504 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-16                        	643424017	         1.852 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	20324672	        59.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	643610655	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	20015511	        59.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	375386139	         3.191 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15959719	        75.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	84025034	        14.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16737024	        71.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7842697	       149.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-16          	  131274	      9708 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16                  	346216347	         3.470 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15633795	        76.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	18389876	        65.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 7824910	       151.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-16                   	207439261	         5.764 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-16              	26552955	        45.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-16               	16665750	        72.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-16            	  127222	      9913 ns/op	   15550 B/op	      76 allocs/op
BenchmarkGoValidRequired-16                  	559292913	         2.128 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14746354	        81.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	652257514	         1.788 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-16           	  132348	      9426 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                       	28517024	        41.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4476746	       271.4 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  161191	      7497 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-16                	  123782	      9836 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                      	34359586	        35.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4867993	       245.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6211080	       190.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-16               	  122546	      9865 ns/op	   15518 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| UUID | 35.47 / 0 allocs | 245.8 / 0 allocs | **6.9x** | 190.2 / 0 allocs | **5.4x** | 9865 / 15518 B / 74 allocs | **278.1x** |
| URL | 41.67 / 0 allocs | 271.4 / 144 B / 1 allocs | **6.5x** | 7497 / 146 B / 1 allocs | **179.9x** | 9836 / 15641 B / 75 allocs | **236.0x** |
| GTE | 1.832 / 0 allocs | 59.59 / 0 allocs | **32.5x** | N/A | N/A | N/A | N/A |
| MinLength | 11.06 / 0 allocs | 65.47 / 0 allocs | **5.9x** | 151.1 / 32 B / 2 allocs | **13.7x** | N/A | N/A |
| LTE | 1.866 / 0 allocs | 59.78 / 0 allocs | **32.0x** | N/A | N/A | N/A | N/A |
| MinItems | 3.470 / 0 allocs | 76.57 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxItems | 3.191 / 0 allocs | 75.47 / 0 allocs | **23.7x** | N/A | N/A | N/A | N/A |
| Enum | 2.655 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Required | 2.128 / 0 allocs | 81.19 / 0 allocs | **38.2x** | 1.788 / 0 allocs | **0.8x** | 9426 / 15472 B / 72 allocs | **4429.5x** |
| MaxLength | 14.36 / 0 allocs | 71.74 / 0 allocs | **5.0x** | 149.3 / 32 B / 2 allocs | **10.4x** | 9708 / 15632 B / 80 allocs | **676.0x** |
| GT | 1.815 / 0 allocs | 59.91 / 0 allocs | **33.0x** | 52.82 / 0 allocs | **29.1x** | N/A | N/A |
| Numeric | 5.764 / 0 allocs | 45.50 / 0 allocs | **7.9x** | 72.45 / 0 allocs | **12.6x** | 9913 / 15550 B / 76 allocs | **1719.8x** |
| LT | 1.852 / 0 allocs | 59.34 / 0 allocs | **32.0x** | N/A | N/A | N/A | N/A |
| Length | 4.398 / 0 allocs | 57.73 / 0 allocs | **13.1x** | 132.5 / 32 B / 2 allocs | **30.1x** | 9504 / 15600 B / 78 allocs | **2161.0x** |
| Alpha | 7.135 / 0 allocs | 244.4 / 0 allocs | **34.3x** | 61.63 / 0 allocs | **8.6x** | 11924 / 16834 B / 97 allocs | **1671.2x** |
| Email | 33.82 / 0 allocs | 603.1 / 89 B / 5 allocs | **17.8x** | 580.3 / 0 allocs | **17.2x** | 10223 / 16087 B / 74 allocs | **302.3x** |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 0.1996 | 0 allocs |
| CELMultipleExpressions | 2.540 | 0 allocs |
| CELBasic | 2.526 | 0 allocs |
| CELCrossField | 1.853 | 0 allocs |
| CELStringLength | 1.760 | 0 allocs |
| CELNumericComparison | 1.790 | 0 allocs |

CEL (Common Expression Language) support allows complex runtime expressions with near-zero overhead.

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
