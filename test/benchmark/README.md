# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-22  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124439245	         9.643 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2537866	       475.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11071233	       108.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55576	     21418 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	643592526	         1.862 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295452238	         4.058 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295638944	         4.056 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349424558	         3.429 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21847718	        58.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1106 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1284031	       934.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70340	     17331 ns/op	   15871 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275266534	         4.358 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481788882	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11249016	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13792201	        88.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482163122	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11031483	       108.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30875678	        39.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9262201	       131.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14059663	        85.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6913384	       174.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160406367	         7.478 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9997219	       120.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4836261	       246.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74200	     16338 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481798329	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11041548	       108.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428097992	         2.813 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11245161	       107.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183240148	         6.549 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8294997	       144.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35957056	        33.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8620804	       138.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4182181	       286.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73422	     17153 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202913139	         5.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8483042	       142.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52036246	        23.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10027963	       119.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4216725	       283.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	156828403	         7.652 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13144395	        91.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9389008	       129.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71563	     16814 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321621559	         3.733 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8008011	       150.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641179804	         1.952 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75616	     15830 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20614362	        57.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2503036	       478.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102889	     11703 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70339	     16890 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23607932	        50.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2688720	       446.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3509102	       350.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70718	     17023 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.813 / 0 allocs | 107.6 / 0 allocs | **38.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.358 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.55 / 0 allocs | 1106 / 88 B / 5 allocs | **18.9x** | 934.4 / 0 allocs | **16.0x** | 17331 / 15871 B / 76 allocs | **296.0x** |
| GTE | 2.491 / 0 allocs | 108.7 / 0 allocs | **43.6x** | N/A | N/A | N/A | N/A |
| MinLength | 23.06 / 0 allocs | 119.1 / 0 allocs | **5.2x** | 283.9 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 50.79 / 0 allocs | 446.9 / 0 allocs | **8.8x** | 350.9 / 0 allocs | **6.9x** | 17023 / 15542 B / 76 allocs | **335.2x** |
| MaxItems | 6.549 / 0 allocs | 144.2 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.33 / 0 allocs | 138.8 / 0 allocs | **4.2x** | 286.6 / 32 B / 2 allocs | **8.6x** | 17153 / 15648 B / 81 allocs | **514.6x** |
| LT | 2.491 / 0 allocs | 108.7 / 0 allocs | **43.6x** | N/A | N/A | N/A | N/A |
| MinItems | 5.920 / 0 allocs | 142.2 / 0 allocs | **24.0x** | N/A | N/A | N/A | N/A |
| Alpha | 9.643 / 0 allocs | 475.7 / 0 allocs | **49.3x** | 108.4 / 0 allocs | **11.2x** | 21418 / 16937 B / 101 allocs | **2221.1x** |
| Required | 3.733 / 0 allocs | 150.2 / 0 allocs | **40.2x** | 1.952 / 0 allocs | **0.5x** | 15830 / 15488 B / 73 allocs | **4240.6x** |
| IPV4 | 39.08 / 0 allocs | 131.2 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 7.478 / 0 allocs | 120.3 / 0 allocs | **16.1x** | 246.8 / 32 B / 2 allocs | **33.0x** | 16338 / 15616 B / 79 allocs | **2184.8x** |
| IPV6 | 85.70 / 0 allocs | 174.0 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.79 / 0 allocs | 478.0 / 144 B / 1 allocs | **8.3x** | 11703 / 146 B / 1 allocs | **202.5x** | 16890 / 15681 B / 77 allocs | **292.3x** |
| Numeric | 7.652 / 0 allocs | 91.45 / 0 allocs | **12.0x** | 129.4 / 0 allocs | **16.9x** | 16814 / 15574 B / 78 allocs | **2197.3x** |
| GT | 2.492 / 0 allocs | 106.6 / 0 allocs | **42.8x** | 88.63 / 0 allocs | **35.6x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.862 | 0 allocs |
| CELMultipleExpressions | 4.058 | 0 allocs |
| CELBasic | 4.056 | 0 allocs |
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
