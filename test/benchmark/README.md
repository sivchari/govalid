# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-01-13  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124248940	         9.659 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2630712	       458.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11228809	       108.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56000	     22683 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	595479408	         1.841 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295993258	         4.056 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296310962	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	321070500	         3.740 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21702159	        55.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1128 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1356970	       887.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69928	     17155 ns/op	   15856 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256551091	         4.674 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427721958	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11094166	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13639563	        88.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	428370427	         2.808 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10918548	       110.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30697107	        39.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9163561	       131.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14082175	        86.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6905907	       173.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147985256	         8.109 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9996051	       120.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4805079	       248.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73382	     16192 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	428699635	         2.798 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11320972	       105.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385341141	         3.112 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10905042	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202725025	         5.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8244709	       145.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37689054	        31.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8641178	       138.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4212216	       284.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74535	     16301 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	175316181	         6.846 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8451951	       142.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48798702	        24.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9899364	       119.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4213281	       284.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124390310	         9.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13791218	        86.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9424147	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   70884	     16773 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296349445	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8044712	       151.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642339967	         1.871 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75886	     15810 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20663169	        58.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2490320	       484.4 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  101888	     11977 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70986	     16791 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24326504	        56.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2680092	       447.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3645500	       334.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71830	     16679 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.112 / 0 allocs | 108.3 / 0 allocs | **34.8x** | N/A | N/A | N/A | N/A |
| Enum | 4.674 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.47 / 0 allocs | 1128 / 89 B / 5 allocs | **20.3x** | 887.3 / 0 allocs | **16.0x** | 17155 / 15856 B / 76 allocs | **309.3x** |
| GTE | 2.808 / 0 allocs | 110.4 / 0 allocs | **39.3x** | N/A | N/A | N/A | N/A |
| MinLength | 24.58 / 0 allocs | 119.2 / 0 allocs | **4.8x** | 284.0 / 32 B / 2 allocs | **11.6x** | N/A | N/A |
| UUID | 56.74 / 0 allocs | 447.4 / 0 allocs | **7.9x** | 334.5 / 0 allocs | **5.9x** | 16679 / 15542 B / 76 allocs | **294.0x** |
| MaxItems | 5.917 / 0 allocs | 145.1 / 0 allocs | **24.5x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.73 / 0 allocs | 138.9 / 0 allocs | **4.4x** | 284.9 / 32 B / 2 allocs | **9.0x** | 16301 / 15648 B / 81 allocs | **513.7x** |
| LT | 2.798 / 0 allocs | 105.2 / 0 allocs | **37.6x** | N/A | N/A | N/A | N/A |
| MinItems | 6.846 / 0 allocs | 142.3 / 0 allocs | **20.8x** | N/A | N/A | N/A | N/A |
| Alpha | 9.659 / 0 allocs | 458.5 / 0 allocs | **47.5x** | 108.1 / 0 allocs | **11.2x** | 22683 / 16937 B / 101 allocs | **2348.4x** |
| Required | 4.053 / 0 allocs | 151.1 / 0 allocs | **37.3x** | 1.871 / 0 allocs | **0.5x** | 15810 / 15488 B / 73 allocs | **3900.8x** |
| IPV4 | 39.16 / 0 allocs | 131.1 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.109 / 0 allocs | 120.3 / 0 allocs | **14.8x** | 248.0 / 32 B / 2 allocs | **30.6x** | 16192 / 15616 B / 79 allocs | **1996.8x** |
| IPV6 | 86.47 / 0 allocs | 173.9 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 58.01 / 0 allocs | 484.4 / 144 B / 1 allocs | **8.4x** | 11977 / 146 B / 1 allocs | **206.5x** | 16791 / 15681 B / 77 allocs | **289.5x** |
| Numeric | 9.645 / 0 allocs | 86.90 / 0 allocs | **9.0x** | 128.3 / 0 allocs | **13.3x** | 16773 / 15574 B / 78 allocs | **1739.0x** |
| GT | 2.801 / 0 allocs | 106.3 / 0 allocs | **38.0x** | 88.22 / 0 allocs | **31.5x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.841 | 0 allocs |
| CELMultipleExpressions | 4.056 | 0 allocs |
| CELBasic | 4.049 | 0 allocs |
| CELCrossField | 3.740 | 0 allocs |
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
