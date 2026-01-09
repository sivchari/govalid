# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-01-09  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	123761997	         9.694 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2561756	       471.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10891249	       110.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   52249	     22807 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	641396228	         1.879 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	290174700	         4.146 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	289367736	         4.139 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	315084907	         3.797 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19908411	        58.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1151 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1292506	       928.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   66044	     18071 ns/op	   15858 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	253466436	         4.734 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	423670444	         2.828 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11156032	       108.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13156959	        90.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	423612400	         2.827 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11011617	       110.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30181244	        39.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8950033	       131.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13652617	        86.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6759025	       176.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	146186458	         8.217 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9953905	       120.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4512735	       262.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   70028	     16707 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	421553473	         2.842 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11010984	       107.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	379655746	         3.144 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11117298	       108.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	199135240	         6.029 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8093410	       148.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36056220	        31.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8463554	       141.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 3964316	       300.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   70410	     17001 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	173362833	         6.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8455474	       142.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	47862757	        24.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10057380	       119.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4004508	       295.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	123171157	         9.738 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13297662	        88.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9214779	       129.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   67425	     17531 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	291667814	         4.099 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7981005	       151.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	637517072	         1.877 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   73604	     16735 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19149795	        61.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2385639	       504.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   98431	     12212 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   69651	     17293 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23404633	        56.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2551706	       469.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3420409	       349.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   67903	     17289 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.144 / 0 allocs | 108.8 / 0 allocs | **34.6x** | N/A | N/A | N/A | N/A |
| Enum | 4.734 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.12 / 0 allocs | 1151 / 89 B / 5 allocs | **19.8x** | 928.5 / 0 allocs | **16.0x** | 18071 / 15858 B / 76 allocs | **310.9x** |
| GTE | 2.827 / 0 allocs | 110.3 / 0 allocs | **39.0x** | N/A | N/A | N/A | N/A |
| MinLength | 24.82 / 0 allocs | 119.6 / 0 allocs | **4.8x** | 295.9 / 32 B / 2 allocs | **11.9x** | N/A | N/A |
| UUID | 56.82 / 0 allocs | 469.6 / 0 allocs | **8.3x** | 349.9 / 0 allocs | **6.2x** | 17289 / 15542 B / 76 allocs | **304.3x** |
| MaxItems | 6.029 / 0 allocs | 148.0 / 0 allocs | **24.5x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.81 / 0 allocs | 141.5 / 0 allocs | **4.4x** | 300.4 / 32 B / 2 allocs | **9.4x** | 17001 / 15648 B / 81 allocs | **534.5x** |
| LT | 2.842 / 0 allocs | 107.8 / 0 allocs | **37.9x** | N/A | N/A | N/A | N/A |
| MinItems | 6.925 / 0 allocs | 142.3 / 0 allocs | **20.5x** | N/A | N/A | N/A | N/A |
| Alpha | 9.694 / 0 allocs | 471.6 / 0 allocs | **48.6x** | 110.8 / 0 allocs | **11.4x** | 22807 / 16937 B / 101 allocs | **2352.7x** |
| Required | 4.099 / 0 allocs | 151.1 / 0 allocs | **36.9x** | 1.877 / 0 allocs | **0.5x** | 16735 / 15488 B / 73 allocs | **4082.7x** |
| IPV4 | 39.48 / 0 allocs | 131.5 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.217 / 0 allocs | 120.9 / 0 allocs | **14.7x** | 262.2 / 32 B / 2 allocs | **31.9x** | 16707 / 15616 B / 79 allocs | **2033.2x** |
| IPV6 | 86.77 / 0 allocs | 176.1 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 61.02 / 0 allocs | 504.5 / 144 B / 1 allocs | **8.3x** | 12212 / 146 B / 1 allocs | **200.1x** | 17293 / 15681 B / 77 allocs | **283.4x** |
| Numeric | 9.738 / 0 allocs | 88.41 / 0 allocs | **9.1x** | 129.5 / 0 allocs | **13.3x** | 17531 / 15574 B / 78 allocs | **1800.3x** |
| GT | 2.828 / 0 allocs | 108.2 / 0 allocs | **38.3x** | 90.10 / 0 allocs | **31.9x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.879 | 0 allocs |
| CELMultipleExpressions | 4.146 | 0 allocs |
| CELBasic | 4.139 | 0 allocs |
| CELCrossField | 3.797 | 0 allocs |
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
