# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-24  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124334235	         9.652 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2639982	       465.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11173584	       108.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56586	     21118 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	645312404	         1.863 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296331912	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296420420	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350424008	         3.429 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21622951	        55.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1087 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1338292	       895.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70258	     17108 ns/op	   15852 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275277868	         4.361 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481681336	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11288293	       106.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13659861	        88.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481542375	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11019217	       109.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30565642	        39.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9245209	       129.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14053294	        85.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6937707	       172.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160748752	         7.464 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9973279	       120.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4893198	       245.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74259	     16242 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481469397	         2.494 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11110320	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428237650	         2.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11207625	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183483102	         6.537 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8286718	       144.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36059589	        33.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8623513	       139.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4191133	       287.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72561	     16534 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202427799	         5.926 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8510286	       141.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52085782	        23.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10109413	       119.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4202919	       283.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	156534842	         7.675 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13160829	        91.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9450040	       128.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   69853	     17287 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320909109	         3.737 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8041528	       149.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642616776	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75675	     15785 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20629993	        57.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2521821	       478.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   92636	     13011 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71084	     17000 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23427306	        52.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2686074	       447.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3571257	       336.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71540	     16896 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.804 / 0 allocs | 107.4 / 0 allocs | **38.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.361 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.43 / 0 allocs | 1087 / 89 B / 5 allocs | **19.6x** | 895.5 / 0 allocs | **16.2x** | 17108 / 15852 B / 76 allocs | **308.6x** |
| GTE | 2.492 / 0 allocs | 109.1 / 0 allocs | **43.8x** | N/A | N/A | N/A | N/A |
| MinLength | 23.06 / 0 allocs | 119.0 / 0 allocs | **5.2x** | 283.9 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 52.22 / 0 allocs | 447.1 / 0 allocs | **8.6x** | 336.6 / 0 allocs | **6.4x** | 16896 / 15542 B / 76 allocs | **323.6x** |
| MaxItems | 6.537 / 0 allocs | 144.9 / 0 allocs | **22.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.35 / 0 allocs | 139.7 / 0 allocs | **4.2x** | 287.1 / 32 B / 2 allocs | **8.6x** | 16534 / 15648 B / 81 allocs | **495.8x** |
| LT | 2.494 / 0 allocs | 108.3 / 0 allocs | **43.4x** | N/A | N/A | N/A | N/A |
| MinItems | 5.926 / 0 allocs | 141.4 / 0 allocs | **23.9x** | N/A | N/A | N/A | N/A |
| Alpha | 9.652 / 0 allocs | 465.3 / 0 allocs | **48.2x** | 108.5 / 0 allocs | **11.2x** | 21118 / 16937 B / 101 allocs | **2187.9x** |
| Required | 3.737 / 0 allocs | 149.0 / 0 allocs | **39.9x** | 1.867 / 0 allocs | **0.5x** | 15785 / 15488 B / 73 allocs | **4224.0x** |
| IPV4 | 39.06 / 0 allocs | 129.7 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 7.464 / 0 allocs | 120.1 / 0 allocs | **16.1x** | 245.4 / 32 B / 2 allocs | **32.9x** | 16242 / 15616 B / 79 allocs | **2176.0x** |
| IPV6 | 85.82 / 0 allocs | 172.9 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.78 / 0 allocs | 478.2 / 144 B / 1 allocs | **8.3x** | 13011 / 145 B / 1 allocs | **225.2x** | 17000 / 15681 B / 77 allocs | **294.2x** |
| Numeric | 7.675 / 0 allocs | 91.20 / 0 allocs | **11.9x** | 128.4 / 0 allocs | **16.7x** | 17287 / 15574 B / 78 allocs | **2252.4x** |
| GT | 2.491 / 0 allocs | 106.9 / 0 allocs | **42.9x** | 88.16 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.863 | 0 allocs |
| CELMultipleExpressions | 4.048 | 0 allocs |
| CELBasic | 4.048 | 0 allocs |
| CELCrossField | 3.429 | 0 allocs |
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
