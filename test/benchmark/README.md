# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-12  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	75474237	        15.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 3062133	       393.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11529799	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   59293	     20073 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	258454536	         4.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	97348264	        11.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	94042515	        12.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	143551230	         8.339 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.059 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19116692	        62.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1025 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1321178	       907.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   73483	     16328 ns/op	   15872 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	94978208	        12.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	212852595	         5.638 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	12317836	        98.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13535427	        90.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	209975864	         5.677 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	12309472	        97.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30105600	        40.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	10980526	       109.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	12496429	       106.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6881506	       177.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        11.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	12807886	        94.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 5318367	       225.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   77022	     15291 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	212395142	         5.648 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	12454416	        96.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	200158466	         6.046 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	12099310	        99.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	100000000	        12.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 9795680	       122.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	48362325	        26.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9410739	       127.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4545109	       263.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   77004	     15497 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	95549952	        12.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 9743126	       123.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	41601799	        28.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10951621	       110.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4591096	       261.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	91142158	        13.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	15639687	        77.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9196941	       134.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   75384	     15796 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8970232	       133.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	566243751	         2.114 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   80404	     14785 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	18847059	        63.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2839633	       421.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  108081	     10959 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   74326	     15925 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	22563330	        63.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2894515	       415.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3682654	       329.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   74167	     15958 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 6.046 / 0 allocs | 99.11 / 0 allocs | **16.4x** | N/A | N/A | N/A | N/A |
| Enum | 12.34 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 62.56 / 0 allocs | 1025 / 89 B / 5 allocs | **16.4x** | 907.0 / 0 allocs | **14.5x** | 16328 / 15872 B / 76 allocs | **261.0x** |
| GTE | 5.677 / 0 allocs | 97.54 / 0 allocs | **17.2x** | N/A | N/A | N/A | N/A |
| MinLength | 28.87 / 0 allocs | 110.6 / 0 allocs | **3.8x** | 261.0 / 32 B / 2 allocs | **9.0x** | N/A | N/A |
| UUID | 63.50 / 0 allocs | 415.3 / 0 allocs | **6.5x** | 329.4 / 0 allocs | **5.2x** | 15958 / 15542 B / 76 allocs | **251.3x** |
| MaxItems | 12.19 / 0 allocs | 122.6 / 0 allocs | **10.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 26.58 / 0 allocs | 127.6 / 0 allocs | **4.8x** | 263.4 / 32 B / 2 allocs | **9.9x** | 15497 / 15648 B / 81 allocs | **583.0x** |
| LT | 5.648 / 0 allocs | 96.07 / 0 allocs | **17.0x** | N/A | N/A | N/A | N/A |
| MinItems | 12.38 / 0 allocs | 123.1 / 0 allocs | **9.9x** | N/A | N/A | N/A | N/A |
| Alpha | 15.86 / 0 allocs | 393.6 / 0 allocs | **24.8x** | 106.3 / 0 allocs | **6.7x** | 20073 / 16937 B / 101 allocs | **1265.6x** |
| Required | 10.57 / 0 allocs | 133.7 / 0 allocs | **12.6x** | 2.114 / 0 allocs | **0.2x** | 14785 / 15488 B / 73 allocs | **1398.8x** |
| IPV4 | 40.62 / 0 allocs | 109.7 / 0 allocs | **2.7x** | N/A | N/A | N/A | N/A |
| Length | 11.32 / 0 allocs | 94.10 / 0 allocs | **8.3x** | 225.2 / 32 B / 2 allocs | **19.9x** | 15291 / 15616 B / 79 allocs | **1350.8x** |
| IPV6 | 106.1 / 0 allocs | 177.1 / 0 allocs | **1.7x** | N/A | N/A | N/A | N/A |
| URL | 63.36 / 0 allocs | 421.8 / 144 B / 1 allocs | **6.7x** | 10959 / 146 B / 1 allocs | **173.0x** | 15925 / 15681 B / 77 allocs | **251.3x** |
| Numeric | 13.07 / 0 allocs | 77.28 / 0 allocs | **5.9x** | 134.4 / 0 allocs | **10.3x** | 15796 / 15574 B / 78 allocs | **1208.6x** |
| GT | 5.638 / 0 allocs | 98.09 / 0 allocs | **17.4x** | 90.36 / 0 allocs | **16.0x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.489 | 0 allocs |
| CELMultipleExpressions | 11.98 | 0 allocs |
| CELBasic | 12.69 | 0 allocs |
| CELCrossField | 8.339 | 0 allocs |
| CELStringLength | 1.000 | 0 allocs |
| CELNumericComparison | 1.059 | 0 allocs |

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
