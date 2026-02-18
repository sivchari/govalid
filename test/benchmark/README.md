# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-02-18  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	60890743	        17.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2474246	       477.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11203880	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55701	     21432 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	262921842	         4.540 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	98405324	        12.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	98708700	        12.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	150625123	         7.921 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	20585583	        59.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1095 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1334458	       897.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   67768	     17732 ns/op	   15854 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	99094752	        11.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	201925375	         5.931 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11168988	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13772110	        87.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	200686483	         5.953 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10786066	       111.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	29908232	        40.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9238453	       130.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13749277	        89.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6917114	       173.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        11.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9951350	       120.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4854049	       248.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73437	     16413 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	202596578	         5.924 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11061538	       108.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	213431214	         5.614 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11196612	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	100000000	        11.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8161902	       147.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35531071	        33.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8577930	       140.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4199226	       285.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72622	     16515 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	91689745	        12.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8380142	       141.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	42488205	        27.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10046664	       119.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4193359	       283.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	95929717	        12.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12893950	        92.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9438073	       128.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   69734	     17335 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8017228	       150.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	550301830	         2.189 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   74563	     16051 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19871079	        60.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2485050	       481.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102362	     11789 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   68935	     17218 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	19528813	        57.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2684007	       447.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3580250	       336.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70680	     17032 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.614 / 0 allocs | 108.0 / 0 allocs | **19.2x** | N/A | N/A | N/A | N/A |
| Enum | 11.95 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.20 / 0 allocs | 1095 / 89 B / 5 allocs | **18.5x** | 897.0 / 0 allocs | **15.2x** | 17732 / 15854 B / 76 allocs | **299.5x** |
| GTE | 5.953 / 0 allocs | 111.6 / 0 allocs | **18.7x** | N/A | N/A | N/A | N/A |
| MinLength | 27.80 / 0 allocs | 119.9 / 0 allocs | **4.3x** | 283.7 / 32 B / 2 allocs | **10.2x** | N/A | N/A |
| UUID | 57.16 / 0 allocs | 447.9 / 0 allocs | **7.8x** | 336.7 / 0 allocs | **5.9x** | 17032 / 15542 B / 76 allocs | **298.0x** |
| MaxItems | 11.85 / 0 allocs | 147.3 / 0 allocs | **12.4x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.69 / 0 allocs | 140.6 / 0 allocs | **4.2x** | 285.9 / 32 B / 2 allocs | **8.5x** | 16515 / 15648 B / 81 allocs | **490.2x** |
| LT | 5.924 / 0 allocs | 108.6 / 0 allocs | **18.3x** | N/A | N/A | N/A | N/A |
| MinItems | 12.78 / 0 allocs | 141.7 / 0 allocs | **11.1x** | N/A | N/A | N/A | N/A |
| Alpha | 17.14 / 0 allocs | 477.4 / 0 allocs | **27.9x** | 108.3 / 0 allocs | **6.3x** | 21432 / 16937 B / 101 allocs | **1250.4x** |
| Required | 10.73 / 0 allocs | 150.7 / 0 allocs | **14.0x** | 2.189 / 0 allocs | **0.2x** | 16051 / 15488 B / 73 allocs | **1495.9x** |
| IPV4 | 40.33 / 0 allocs | 130.8 / 0 allocs | **3.2x** | N/A | N/A | N/A | N/A |
| Length | 11.23 / 0 allocs | 120.7 / 0 allocs | **10.7x** | 248.3 / 32 B / 2 allocs | **22.1x** | 16413 / 15616 B / 79 allocs | **1461.5x** |
| IPV6 | 89.39 / 0 allocs | 173.2 / 0 allocs | **1.9x** | N/A | N/A | N/A | N/A |
| URL | 60.17 / 0 allocs | 481.8 / 144 B / 1 allocs | **8.0x** | 11789 / 145 B / 1 allocs | **195.9x** | 17218 / 15681 B / 77 allocs | **286.2x** |
| Numeric | 12.49 / 0 allocs | 92.85 / 0 allocs | **7.4x** | 128.0 / 0 allocs | **10.2x** | 17335 / 15574 B / 78 allocs | **1387.9x** |
| GT | 5.931 / 0 allocs | 107.4 / 0 allocs | **18.1x** | 87.99 / 0 allocs | **14.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.540 | 0 allocs |
| CELMultipleExpressions | 12.18 | 0 allocs |
| CELBasic | 12.27 | 0 allocs |
| CELCrossField | 7.921 | 0 allocs |
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
