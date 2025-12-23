# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-23  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124318562	         9.657 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2636830	       460.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11264620	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56450	     21061 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	653212812	         1.845 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296008987	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296415496	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320758332	         3.744 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21954350	        57.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1113 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1352814	       885.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70240	     17006 ns/op	   15854 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256703253	         4.671 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	428468708	         2.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11334403	       106.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13777251	        88.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	427993986	         2.872 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10941897	       110.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30508236	        39.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9245883	       131.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14114186	        85.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6922209	       172.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147885306	         8.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9967434	       120.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4829000	       247.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74630	     16147 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	428033017	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11334510	       106.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385555042	         3.112 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11160780	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202767606	         5.916 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8211666	       144.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37070136	        31.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8657838	       139.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4172248	       287.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73038	     16295 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	175312939	         6.845 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8474239	       141.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48798110	        24.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10178030	       118.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4234707	       283.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124177588	         9.661 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13690887	        87.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9410319	       128.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71136	     16668 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296790992	         4.045 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8017930	       148.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641854782	         1.919 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76899	     15592 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20762107	        57.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2497398	       478.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102992	     11673 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70636	     16732 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	21522344	        54.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2678853	       448.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3567164	       336.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70225	     16749 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.112 / 0 allocs | 107.7 / 0 allocs | **34.6x** | N/A | N/A | N/A | N/A |
| Enum | 4.671 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.01 / 0 allocs | 1113 / 89 B / 5 allocs | **19.5x** | 885.6 / 0 allocs | **15.5x** | 17006 / 15854 B / 76 allocs | **298.3x** |
| GTE | 2.872 / 0 allocs | 110.2 / 0 allocs | **38.4x** | N/A | N/A | N/A | N/A |
| MinLength | 24.57 / 0 allocs | 118.7 / 0 allocs | **4.8x** | 283.9 / 32 B / 2 allocs | **11.6x** | N/A | N/A |
| UUID | 54.32 / 0 allocs | 448.4 / 0 allocs | **8.3x** | 336.8 / 0 allocs | **6.2x** | 16749 / 15542 B / 76 allocs | **308.3x** |
| MaxItems | 5.916 / 0 allocs | 144.1 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.76 / 0 allocs | 139.0 / 0 allocs | **4.4x** | 287.2 / 32 B / 2 allocs | **9.0x** | 16295 / 15648 B / 81 allocs | **513.1x** |
| LT | 2.802 / 0 allocs | 106.1 / 0 allocs | **37.9x** | N/A | N/A | N/A | N/A |
| MinItems | 6.845 / 0 allocs | 141.2 / 0 allocs | **20.6x** | N/A | N/A | N/A | N/A |
| Alpha | 9.657 / 0 allocs | 460.7 / 0 allocs | **47.7x** | 107.4 / 0 allocs | **11.1x** | 21061 / 16937 B / 101 allocs | **2180.9x** |
| Required | 4.045 / 0 allocs | 148.9 / 0 allocs | **36.8x** | 1.919 / 0 allocs | **0.5x** | 15592 / 15488 B / 73 allocs | **3854.6x** |
| IPV4 | 39.13 / 0 allocs | 131.3 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 8.113 / 0 allocs | 120.5 / 0 allocs | **14.9x** | 247.3 / 32 B / 2 allocs | **30.5x** | 16147 / 15616 B / 79 allocs | **1990.3x** |
| IPV6 | 85.03 / 0 allocs | 172.9 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.77 / 0 allocs | 478.6 / 144 B / 1 allocs | **8.3x** | 11673 / 145 B / 1 allocs | **202.1x** | 16732 / 15681 B / 77 allocs | **289.6x** |
| Numeric | 9.661 / 0 allocs | 87.42 / 0 allocs | **9.0x** | 128.5 / 0 allocs | **13.3x** | 16668 / 15574 B / 78 allocs | **1725.3x** |
| GT | 2.800 / 0 allocs | 106.2 / 0 allocs | **37.9x** | 88.10 / 0 allocs | **31.5x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.845 | 0 allocs |
| CELMultipleExpressions | 4.054 | 0 allocs |
| CELBasic | 4.048 | 0 allocs |
| CELCrossField | 3.744 | 0 allocs |
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
