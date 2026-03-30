# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-30  
**Platform:** Linux 6.17.0-1008-azure x86_64  
**Go version:** go1.25.0

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	74641412	        15.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2605677	       463.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11048797	       109.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60978	     19661 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	251765155	         4.765 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	153539148	         7.816 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	20512742	        60.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1070 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1322721	       907.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   76012	     15739 ns/op	   15707 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	99173905	        11.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	202402563	         5.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10478199	       114.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13333551	        91.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	202271432	         5.926 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10887858	       110.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	29397487	        40.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8931746	       133.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	12582652	        95.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6644905	       181.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        10.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11438659	       105.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4625588	       258.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   78554	     15019 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	213883897	         5.610 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11639101	       103.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202336524	         5.929 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11108692	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	100000000	        11.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8498382	       141.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	48462582	        23.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9137211	       131.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4257320	       280.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   77806	     15322 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	95414042	        12.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8402922	       143.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	44049511	        27.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9580960	       125.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4153130	       285.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	94135854	        12.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12781462	        93.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9195247	       130.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   77910	     15319 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7950030	       151.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642167092	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   81661	     14439 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	19510218	        61.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2325808	       512.7 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103299	     12400 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   76656	     15463 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	22923930	        57.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2505012	       481.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3397558	       354.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   77312	     15497 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.929 / 0 allocs | 107.7 / 0 allocs | **18.2x** | N/A | N/A | N/A | N/A |
| Enum | 11.89 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 60.49 / 0 allocs | 1070 / 88 B / 5 allocs | **17.7x** | 907.1 / 0 allocs | **15.0x** | 15739 / 15707 B / 74 allocs | **260.2x** |
| GTE | 5.926 / 0 allocs | 110.5 / 0 allocs | **18.6x** | N/A | N/A | N/A | N/A |
| MinLength | 27.16 / 0 allocs | 125.7 / 0 allocs | **4.6x** | 285.8 / 32 B / 2 allocs | **10.5x** | N/A | N/A |
| UUID | 57.26 / 0 allocs | 481.0 / 0 allocs | **8.4x** | 354.0 / 0 allocs | **6.2x** | 15497 / 15501 B / 74 allocs | **270.6x** |
| MaxItems | 11.53 / 0 allocs | 141.7 / 0 allocs | **12.3x** | N/A | N/A | N/A | N/A |
| MaxLength | 23.71 / 0 allocs | 131.6 / 0 allocs | **5.6x** | 280.6 / 32 B / 2 allocs | **11.8x** | 15322 / 15632 B / 80 allocs | **646.2x** |
| LT | 5.610 / 0 allocs | 103.7 / 0 allocs | **18.5x** | N/A | N/A | N/A | N/A |
| MinItems | 12.47 / 0 allocs | 143.2 / 0 allocs | **11.5x** | N/A | N/A | N/A | N/A |
| Alpha | 15.92 / 0 allocs | 463.5 / 0 allocs | **29.1x** | 109.9 / 0 allocs | **6.9x** | 19661 / 16815 B / 97 allocs | **1235.0x** |
| Required | 10.60 / 0 allocs | 151.4 / 0 allocs | **14.3x** | 1.869 / 0 allocs | **0.2x** | 14439 / 15472 B / 72 allocs | **1362.2x** |
| IPV4 | 40.65 / 0 allocs | 133.9 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 10.61 / 0 allocs | 105.4 / 0 allocs | **9.9x** | 258.1 / 32 B / 2 allocs | **24.3x** | 15019 / 15600 B / 78 allocs | **1415.6x** |
| IPV6 | 95.99 / 0 allocs | 181.5 / 0 allocs | **1.9x** | N/A | N/A | N/A | N/A |
| URL | 61.37 / 0 allocs | 512.7 / 144 B / 1 allocs | **8.4x** | 12400 / 147 B / 1 allocs | **202.1x** | 15463 / 15641 B / 75 allocs | **252.0x** |
| Numeric | 12.49 / 0 allocs | 93.14 / 0 allocs | **7.5x** | 130.7 / 0 allocs | **10.5x** | 15319 / 15533 B / 76 allocs | **1226.5x** |
| GT | 5.925 / 0 allocs | 114.4 / 0 allocs | **19.3x** | 91.15 / 0 allocs | **15.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.765 | 0 allocs |
| CELMultipleExpressions | 11.88 | 0 allocs |
| CELBasic | 11.88 | 0 allocs |
| CELCrossField | 7.816 | 0 allocs |
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
