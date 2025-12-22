# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-22  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124344760	         9.649 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2586979	       487.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10153946	       124.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   54510	     21427 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644147716	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296082746	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295919352	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349659625	         3.430 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21726349	        55.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1110 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1000000	      1186 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69370	     17164 ns/op	   15878 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	274661487	         4.362 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481374060	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11270008	       107.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13608229	        88.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481846884	         2.502 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11030232	       109.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30708554	        39.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9205842	       130.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14052735	        85.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6936402	       172.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160552826	         7.471 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9941464	       120.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4844788	       247.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74584	     16116 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481268814	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10983736	       108.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	426944060	         2.806 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11198437	       106.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183682300	         6.531 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8340866	       144.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36002370	        33.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8609683	       139.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4186880	       287.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74245	     16243 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202820672	         5.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8496592	       141.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52096533	        23.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10171239	       118.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4215984	       284.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	158124330	         7.598 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13188805	        90.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9454156	       128.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72487	     16690 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321231757	         3.735 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8013486	       149.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642644934	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76946	     15626 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20851320	        57.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2492386	       480.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102892	     11728 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70498	     16775 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23737236	        52.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2683592	       447.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3586809	       336.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70992	     16740 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.806 / 0 allocs | 106.9 / 0 allocs | **38.1x** | N/A | N/A | N/A | N/A |
| Enum | 4.362 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.42 / 0 allocs | 1110 / 89 B / 5 allocs | **20.0x** | 1186 / 0 allocs | **21.4x** | 17164 / 15878 B / 76 allocs | **309.7x** |
| GTE | 2.502 / 0 allocs | 109.3 / 0 allocs | **43.7x** | N/A | N/A | N/A | N/A |
| MinLength | 23.05 / 0 allocs | 118.5 / 0 allocs | **5.1x** | 284.6 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 52.48 / 0 allocs | 447.4 / 0 allocs | **8.5x** | 336.0 / 0 allocs | **6.4x** | 16740 / 15542 B / 76 allocs | **319.0x** |
| MaxItems | 6.531 / 0 allocs | 144.2 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.40 / 0 allocs | 139.4 / 0 allocs | **4.2x** | 287.2 / 32 B / 2 allocs | **8.6x** | 16243 / 15648 B / 81 allocs | **486.3x** |
| LT | 2.491 / 0 allocs | 108.8 / 0 allocs | **43.7x** | N/A | N/A | N/A | N/A |
| MinItems | 5.917 / 0 allocs | 141.7 / 0 allocs | **23.9x** | N/A | N/A | N/A | N/A |
| Alpha | 9.649 / 0 allocs | 487.6 / 0 allocs | **50.5x** | 124.5 / 0 allocs | **12.9x** | 21427 / 16937 B / 101 allocs | **2220.6x** |
| Required | 3.735 / 0 allocs | 149.8 / 0 allocs | **40.1x** | 1.868 / 0 allocs | **0.5x** | 15626 / 15488 B / 73 allocs | **4183.7x** |
| IPV4 | 39.13 / 0 allocs | 130.3 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 7.471 / 0 allocs | 120.2 / 0 allocs | **16.1x** | 247.6 / 32 B / 2 allocs | **33.1x** | 16116 / 15616 B / 79 allocs | **2157.1x** |
| IPV6 | 85.66 / 0 allocs | 172.8 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.57 / 0 allocs | 480.1 / 144 B / 1 allocs | **8.3x** | 11728 / 145 B / 1 allocs | **203.7x** | 16775 / 15681 B / 77 allocs | **291.4x** |
| Numeric | 7.598 / 0 allocs | 90.80 / 0 allocs | **12.0x** | 128.1 / 0 allocs | **16.9x** | 16690 / 15574 B / 78 allocs | **2196.6x** |
| GT | 2.492 / 0 allocs | 107.0 / 0 allocs | **42.9x** | 88.17 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.868 | 0 allocs |
| CELMultipleExpressions | 4.052 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.430 | 0 allocs |
| CELStringLength | 1.000 | 0 allocs |
| CELNumericComparison | 1.000 | 0 allocs |

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
