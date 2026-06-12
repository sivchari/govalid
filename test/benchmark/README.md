# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-06-12  
**Platform:** Linux 6.17.0-1018-azure x86_64  
**Go version:** go1.25.0

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	71514444	        15.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2630534	       460.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10991576	       109.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60444	     19810 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	251837133	         4.774 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	153523504	         7.813 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	20654326	        58.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1087 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1315880	       912.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   75105	     15865 ns/op	   15725 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	99216987	        11.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	202195380	         6.005 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10479457	       114.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13432944	        89.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	202243147	         5.931 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11063841	       109.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	29424877	        40.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8975406	       134.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	12997882	        92.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6674433	       180.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        10.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11392998	       105.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4610112	       259.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   81027	     14994 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	213601844	         5.613 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11639534	       103.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202314708	         5.941 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11113272	       109.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	100000000	        11.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8590350	       139.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	49165819	        23.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9326305	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4267443	       280.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   78894	     15202 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	94901766	        12.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8293568	       144.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	44009064	        27.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9528295	       126.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4166032	       287.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	94850112	        12.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12810391	        93.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9146454	       129.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   75423	     15550 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8026435	       149.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641648810	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   77443	     15217 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	19553230	        61.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2309800	       516.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102656	     11684 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   74845	     16150 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	22419398	        58.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2556877	       473.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3282164	       366.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   74322	     15987 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.941 / 0 allocs | 109.1 / 0 allocs | **18.4x** | N/A | N/A | N/A | N/A |
| Enum | 11.86 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.35 / 0 allocs | 1087 / 88 B / 5 allocs | **18.6x** | 912.2 / 0 allocs | **15.6x** | 15865 / 15725 B / 74 allocs | **271.9x** |
| GTE | 5.931 / 0 allocs | 109.3 / 0 allocs | **18.4x** | N/A | N/A | N/A | N/A |
| MinLength | 27.18 / 0 allocs | 126.4 / 0 allocs | **4.7x** | 287.0 / 32 B / 2 allocs | **10.6x** | N/A | N/A |
| UUID | 58.17 / 0 allocs | 473.4 / 0 allocs | **8.1x** | 366.3 / 0 allocs | **6.3x** | 15987 / 15501 B / 74 allocs | **274.8x** |
| MaxItems | 11.55 / 0 allocs | 139.5 / 0 allocs | **12.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 23.78 / 0 allocs | 128.3 / 0 allocs | **5.4x** | 280.9 / 32 B / 2 allocs | **11.8x** | 15202 / 15632 B / 80 allocs | **639.3x** |
| LT | 5.613 / 0 allocs | 103.3 / 0 allocs | **18.4x** | N/A | N/A | N/A | N/A |
| MinItems | 12.49 / 0 allocs | 144.8 / 0 allocs | **11.6x** | N/A | N/A | N/A | N/A |
| Alpha | 15.96 / 0 allocs | 460.6 / 0 allocs | **28.9x** | 109.3 / 0 allocs | **6.8x** | 19810 / 16815 B / 97 allocs | **1241.2x** |
| Required | 10.61 / 0 allocs | 149.6 / 0 allocs | **14.1x** | 1.870 / 0 allocs | **0.2x** | 15217 / 15472 B / 72 allocs | **1434.2x** |
| IPV4 | 40.63 / 0 allocs | 134.3 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 10.60 / 0 allocs | 105.0 / 0 allocs | **9.9x** | 259.0 / 32 B / 2 allocs | **24.4x** | 14994 / 15600 B / 78 allocs | **1414.5x** |
| IPV6 | 92.48 / 0 allocs | 180.8 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 61.27 / 0 allocs | 516.0 / 144 B / 1 allocs | **8.4x** | 11684 / 146 B / 1 allocs | **190.7x** | 16150 / 15641 B / 75 allocs | **263.6x** |
| Numeric | 12.50 / 0 allocs | 93.48 / 0 allocs | **7.5x** | 129.7 / 0 allocs | **10.4x** | 15550 / 15533 B / 76 allocs | **1244.0x** |
| GT | 6.005 / 0 allocs | 114.6 / 0 allocs | **19.1x** | 89.62 / 0 allocs | **14.9x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.774 | 0 allocs |
| CELMultipleExpressions | 11.87 | 0 allocs |
| CELBasic | 11.87 | 0 allocs |
| CELCrossField | 7.813 | 0 allocs |
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
