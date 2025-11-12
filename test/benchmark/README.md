# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-12  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124455841	         9.642 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2585251	       470.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10946227	       110.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56869	     20998 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644252737	         1.856 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296430279	         4.047 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296014846	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350366830	         3.426 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21741991	        57.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1084 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1331311	       898.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69554	     17041 ns/op	   15851 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275179399	         4.366 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481487469	         2.494 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11316384	       106.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13632044	        88.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482031350	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11171710	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30784900	        39.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9163455	       130.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14165006	        85.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6892460	       176.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160500240	         7.474 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10057432	       119.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4858594	       246.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74605	     16061 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481493626	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11308801	       106.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428021547	         2.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11220768	       107.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183699692	         6.531 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8311438	       144.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36099948	        33.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8651029	       139.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4258196	       281.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74410	     16269 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	203150757	         5.908 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8498094	       141.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52104174	        23.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10013631	       120.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4228650	       282.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157887442	         7.606 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13896844	        86.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9382813	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71234	     16640 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321477176	         3.734 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7917554	       150.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642200263	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   77318	     15529 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20770219	        57.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2506255	       477.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103458	     11656 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71626	     16880 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24506185	        48.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2683366	       447.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3452661	       352.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70848	     16677 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.800 / 0 allocs | 107.1 / 0 allocs | **38.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.366 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.86 / 0 allocs | 1084 / 89 B / 5 allocs | **18.7x** | 898.2 / 0 allocs | **15.5x** | 17041 / 15851 B / 76 allocs | **294.5x** |
| GTE | 2.491 / 0 allocs | 108.3 / 0 allocs | **43.5x** | N/A | N/A | N/A | N/A |
| MinLength | 23.02 / 0 allocs | 120.8 / 0 allocs | **5.2x** | 282.6 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.55 / 0 allocs | 447.2 / 0 allocs | **9.2x** | 352.9 / 0 allocs | **7.3x** | 16677 / 15542 B / 76 allocs | **343.5x** |
| MaxItems | 6.531 / 0 allocs | 144.3 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.29 / 0 allocs | 139.1 / 0 allocs | **4.2x** | 281.9 / 32 B / 2 allocs | **8.5x** | 16269 / 15648 B / 81 allocs | **488.7x** |
| LT | 2.491 / 0 allocs | 106.2 / 0 allocs | **42.6x** | N/A | N/A | N/A | N/A |
| MinItems | 5.908 / 0 allocs | 141.6 / 0 allocs | **24.0x** | N/A | N/A | N/A | N/A |
| Alpha | 9.642 / 0 allocs | 470.1 / 0 allocs | **48.8x** | 110.4 / 0 allocs | **11.4x** | 20998 / 16937 B / 101 allocs | **2177.8x** |
| Required | 3.734 / 0 allocs | 150.8 / 0 allocs | **40.4x** | 1.869 / 0 allocs | **0.5x** | 15529 / 15488 B / 73 allocs | **4158.8x** |
| IPV4 | 39.32 / 0 allocs | 130.7 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 7.474 / 0 allocs | 119.5 / 0 allocs | **16.0x** | 246.2 / 32 B / 2 allocs | **32.9x** | 16061 / 15616 B / 79 allocs | **2148.9x** |
| IPV6 | 85.39 / 0 allocs | 176.3 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 57.65 / 0 allocs | 477.6 / 144 B / 1 allocs | **8.3x** | 11656 / 145 B / 1 allocs | **202.2x** | 16880 / 15681 B / 77 allocs | **292.8x** |
| Numeric | 7.606 / 0 allocs | 86.54 / 0 allocs | **11.4x** | 128.3 / 0 allocs | **16.9x** | 16640 / 15574 B / 78 allocs | **2187.7x** |
| GT | 2.494 / 0 allocs | 106.4 / 0 allocs | **42.7x** | 88.32 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.856 | 0 allocs |
| CELMultipleExpressions | 4.047 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.426 | 0 allocs |
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
