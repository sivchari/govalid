# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-10  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124492195	         9.640 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2612836	       462.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11145356	       107.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   53762	     22840 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	639392539	         1.859 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296323605	         4.116 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296114209	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350165710	         3.436 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21779529	        56.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1092 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1318052	       911.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   67694	     17583 ns/op	   15875 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275255562	         4.357 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481503277	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11319391	       106.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13684224	        88.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481375324	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11136050	       108.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30696880	        39.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9252871	       130.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13954783	        85.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6866833	       175.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160551850	         7.472 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10002847	       119.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4829133	       248.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73278	     16310 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481532192	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11402556	       105.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428553212	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11123853	       107.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183752450	         6.533 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8245473	       145.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36071544	        33.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8594156	       139.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4211404	       285.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72169	     16536 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	203041146	         5.912 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8488041	       141.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51721411	        23.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10075611	       118.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4197193	       284.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157903936	         7.600 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13892756	        86.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9380478	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71013	     16840 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321079572	         3.736 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7917746	       150.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643329258	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76394	     15674 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20669840	        57.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2494453	       480.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102745	     12190 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   69610	     17216 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24766198	        48.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2632053	       460.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3506182	       349.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70410	     17171 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.801 / 0 allocs | 107.5 / 0 allocs | **38.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.357 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 56.34 / 0 allocs | 1092 / 89 B / 5 allocs | **19.4x** | 911.5 / 0 allocs | **16.2x** | 17583 / 15875 B / 76 allocs | **312.1x** |
| GTE | 2.491 / 0 allocs | 108.7 / 0 allocs | **43.6x** | N/A | N/A | N/A | N/A |
| MinLength | 23.07 / 0 allocs | 118.7 / 0 allocs | **5.1x** | 284.8 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.57 / 0 allocs | 460.7 / 0 allocs | **9.5x** | 349.6 / 0 allocs | **7.2x** | 17171 / 15542 B / 76 allocs | **353.5x** |
| MaxItems | 6.533 / 0 allocs | 145.8 / 0 allocs | **22.3x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.34 / 0 allocs | 139.7 / 0 allocs | **4.2x** | 285.9 / 32 B / 2 allocs | **8.6x** | 16536 / 15648 B / 81 allocs | **496.0x** |
| LT | 2.489 / 0 allocs | 105.5 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinItems | 5.912 / 0 allocs | 141.9 / 0 allocs | **24.0x** | N/A | N/A | N/A | N/A |
| Alpha | 9.640 / 0 allocs | 462.4 / 0 allocs | **48.0x** | 107.8 / 0 allocs | **11.2x** | 22840 / 16937 B / 101 allocs | **2369.3x** |
| Required | 3.736 / 0 allocs | 150.2 / 0 allocs | **40.2x** | 1.867 / 0 allocs | **0.5x** | 15674 / 15488 B / 73 allocs | **4195.4x** |
| IPV4 | 39.13 / 0 allocs | 130.7 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 7.472 / 0 allocs | 119.9 / 0 allocs | **16.0x** | 248.9 / 32 B / 2 allocs | **33.3x** | 16310 / 15616 B / 79 allocs | **2182.8x** |
| IPV6 | 85.94 / 0 allocs | 175.3 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.83 / 0 allocs | 480.8 / 144 B / 1 allocs | **8.3x** | 12190 / 147 B / 1 allocs | **210.8x** | 17216 / 15681 B / 77 allocs | **297.7x** |
| Numeric | 7.600 / 0 allocs | 86.57 / 0 allocs | **11.4x** | 128.3 / 0 allocs | **16.9x** | 16840 / 15574 B / 78 allocs | **2215.8x** |
| GT | 2.491 / 0 allocs | 106.8 / 0 allocs | **42.9x** | 88.23 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.859 | 0 allocs |
| CELMultipleExpressions | 4.116 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.436 | 0 allocs |
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
