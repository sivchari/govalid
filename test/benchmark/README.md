# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-30  
**Platform:** Linux 6.17.0-1008-azure x86_64  
**Go version:** go1.25.0

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	74132899	        15.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2592199	       466.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10954714	       111.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   61276	     19550 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	251953106	         4.764 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	153802768	         7.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	20522743	        60.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1079 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1319280	       912.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   74631	     15832 ns/op	   15714 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	100000000	        11.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	201806620	         5.941 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10452874	       114.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13492263	        89.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	199706391	         5.982 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11100649	       108.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	29490934	        40.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8845192	       136.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13033093	        92.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6679935	       181.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        10.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11483802	       105.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4601137	       258.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80630	     14911 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	214072790	         5.610 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11627176	       103.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	200797327	         5.948 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11051527	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	100000000	        11.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8513325	       141.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	50720934	        23.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9363927	       127.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4268869	       279.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   78555	     15035 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	95357245	        12.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8113696	       144.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	43971098	        27.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9882321	       121.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4187664	       285.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	95075448	        12.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12820597	        93.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9173442	       131.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   77716	     15425 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7927954	       150.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641733490	         1.871 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   82280	     14504 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	19545720	        61.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2366829	       507.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  100192	     11863 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   76741	     15455 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	25271698	        60.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2574976	       465.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3400773	       359.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   77481	     15372 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.948 / 0 allocs | 108.0 / 0 allocs | **18.2x** | N/A | N/A | N/A | N/A |
| Enum | 11.85 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 60.12 / 0 allocs | 1079 / 88 B / 5 allocs | **17.9x** | 912.7 / 0 allocs | **15.2x** | 15832 / 15714 B / 74 allocs | **263.3x** |
| GTE | 5.982 / 0 allocs | 108.4 / 0 allocs | **18.1x** | N/A | N/A | N/A | N/A |
| MinLength | 27.12 / 0 allocs | 121.4 / 0 allocs | **4.5x** | 285.0 / 32 B / 2 allocs | **10.5x** | N/A | N/A |
| UUID | 60.23 / 0 allocs | 465.7 / 0 allocs | **7.7x** | 359.5 / 0 allocs | **6.0x** | 15372 / 15501 B / 74 allocs | **255.2x** |
| MaxItems | 11.53 / 0 allocs | 141.7 / 0 allocs | **12.3x** | N/A | N/A | N/A | N/A |
| MaxLength | 23.83 / 0 allocs | 127.8 / 0 allocs | **5.4x** | 279.2 / 32 B / 2 allocs | **11.7x** | 15035 / 15632 B / 80 allocs | **630.9x** |
| LT | 5.610 / 0 allocs | 103.2 / 0 allocs | **18.4x** | N/A | N/A | N/A | N/A |
| MinItems | 12.47 / 0 allocs | 144.1 / 0 allocs | **11.6x** | N/A | N/A | N/A | N/A |
| Alpha | 15.92 / 0 allocs | 466.0 / 0 allocs | **29.3x** | 111.3 / 0 allocs | **7.0x** | 19550 / 16815 B / 97 allocs | **1228.0x** |
| Required | 10.60 / 0 allocs | 150.6 / 0 allocs | **14.2x** | 1.871 / 0 allocs | **0.2x** | 14504 / 15472 B / 72 allocs | **1368.3x** |
| IPV4 | 40.57 / 0 allocs | 136.3 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 10.60 / 0 allocs | 105.2 / 0 allocs | **9.9x** | 258.6 / 32 B / 2 allocs | **24.4x** | 14911 / 15600 B / 78 allocs | **1406.7x** |
| IPV6 | 92.74 / 0 allocs | 181.6 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 61.41 / 0 allocs | 507.1 / 144 B / 1 allocs | **8.3x** | 11863 / 146 B / 1 allocs | **193.2x** | 15455 / 15641 B / 75 allocs | **251.7x** |
| Numeric | 12.50 / 0 allocs | 93.76 / 0 allocs | **7.5x** | 131.0 / 0 allocs | **10.5x** | 15425 / 15533 B / 76 allocs | **1234.0x** |
| GT | 5.941 / 0 allocs | 114.6 / 0 allocs | **19.3x** | 89.68 / 0 allocs | **15.1x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.764 | 0 allocs |
| CELMultipleExpressions | 11.99 | 0 allocs |
| CELBasic | 11.87 | 0 allocs |
| CELCrossField | 7.804 | 0 allocs |
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
