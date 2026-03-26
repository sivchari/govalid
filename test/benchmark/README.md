# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-26  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.25.0

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	80359317	        15.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 3037429	       395.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	12176450	        98.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   69470	     17130 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	217722310	         5.729 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	114737794	         9.737 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	135149913	         8.871 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	195711608	         6.131 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.138 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	654745725	         1.841 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21002610	        56.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1004 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1502206	       798.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   90734	     13357 ns/op	   15734 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	123651740	         9.699 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	291804354	         4.114 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	12241327	        98.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	15381888	        78.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	292154749	         4.105 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11972371	       100.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	36147583	        33.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	10027143	       119.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14231358	        84.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 7239336	       165.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	129902440	         9.244 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	12634267	        94.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 5082960	       235.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   94424	     12722 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	292621120	         4.120 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11753985	       101.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	292791421	         4.102 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11605704	       103.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	137501107	         8.725 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 9560820	       127.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	48736573	        24.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	10331326	       116.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4676186	       254.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   94562	     12782 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	100000000	        10.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 9616671	       124.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	40089560	        27.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10759467	       111.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4611982	       259.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	100000000	        11.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	16104025	        74.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	10616049	       112.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   91992	     13046 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	150092530	         7.994 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 9420748	       127.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	1000000000	         1.160 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   96542	     12199 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	18558751	        64.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2689418	       446.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  115303	     10517 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   92259	     13026 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	27872449	        42.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2823292	       425.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3766088	       319.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   91278	     13084 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 4.102 / 0 allocs | 103.0 / 0 allocs | **25.1x** | N/A | N/A | N/A | N/A |
| Enum | 9.699 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 56.82 / 0 allocs | 1004 / 89 B / 5 allocs | **17.7x** | 798.6 / 0 allocs | **14.1x** | 13357 / 15734 B / 74 allocs | **235.1x** |
| GTE | 4.105 / 0 allocs | 100.1 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| MinLength | 27.80 / 0 allocs | 111.4 / 0 allocs | **4.0x** | 259.2 / 32 B / 2 allocs | **9.3x** | N/A | N/A |
| UUID | 42.63 / 0 allocs | 425.0 / 0 allocs | **10.0x** | 319.1 / 0 allocs | **7.5x** | 13084 / 15501 B / 74 allocs | **306.9x** |
| MaxItems | 8.725 / 0 allocs | 127.3 / 0 allocs | **14.6x** | N/A | N/A | N/A | N/A |
| MaxLength | 24.60 / 0 allocs | 116.1 / 0 allocs | **4.7x** | 254.8 / 32 B / 2 allocs | **10.4x** | 12782 / 15632 B / 80 allocs | **519.6x** |
| LT | 4.120 / 0 allocs | 101.7 / 0 allocs | **24.7x** | N/A | N/A | N/A | N/A |
| MinItems | 10.79 / 0 allocs | 124.4 / 0 allocs | **11.5x** | N/A | N/A | N/A | N/A |
| Alpha | 15.00 / 0 allocs | 395.3 / 0 allocs | **26.4x** | 98.28 / 0 allocs | **6.6x** | 17130 / 16815 B / 97 allocs | **1142.0x** |
| Required | 7.994 / 0 allocs | 127.4 / 0 allocs | **15.9x** | 1.160 / 0 allocs | **0.1x** | 12199 / 15472 B / 72 allocs | **1526.0x** |
| IPV4 | 33.26 / 0 allocs | 119.5 / 0 allocs | **3.6x** | N/A | N/A | N/A | N/A |
| Length | 9.244 / 0 allocs | 94.54 / 0 allocs | **10.2x** | 235.8 / 32 B / 2 allocs | **25.5x** | 12722 / 15600 B / 78 allocs | **1376.2x** |
| IPV6 | 84.43 / 0 allocs | 165.5 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 64.21 / 0 allocs | 446.5 / 144 B / 1 allocs | **7.0x** | 10517 / 146 B / 1 allocs | **163.8x** | 13026 / 15641 B / 75 allocs | **202.9x** |
| Numeric | 11.17 / 0 allocs | 74.32 / 0 allocs | **6.7x** | 112.9 / 0 allocs | **10.1x** | 13046 / 15533 B / 76 allocs | **1167.9x** |
| GT | 4.114 / 0 allocs | 98.00 / 0 allocs | **23.8x** | 78.22 / 0 allocs | **19.0x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 5.729 | 0 allocs |
| CELMultipleExpressions | 9.737 | 0 allocs |
| CELBasic | 8.871 | 0 allocs |
| CELCrossField | 6.131 | 0 allocs |
| CELStringLength | 1.138 | 0 allocs |
| CELNumericComparison | 1.841 | 0 allocs |

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
