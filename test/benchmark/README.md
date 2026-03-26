# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-26  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.25.0

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	74664996	        15.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2624628	       462.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10805110	       109.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   61569	     20285 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	253221326	         4.757 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	154437519	         7.741 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21056260	        57.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1097 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1282168	       934.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   76714	     15640 ns/op	   15725 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	99423055	        11.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	203059203	         5.901 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	 9999061	       116.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13871023	        89.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	201914776	         5.933 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11081175	       108.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	29515324	        40.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8844055	       135.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13242154	        92.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6661339	       180.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        10.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11488341	       105.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4595395	       258.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80541	     14884 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	208851296	         5.678 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11685810	       103.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202819285	         5.929 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11042853	       108.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	100000000	        11.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8581564	       139.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	51940215	        23.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9130726	       131.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4285038	       278.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   79695	     14920 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	96032127	        12.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8415391	       142.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	44255066	        27.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9570379	       125.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4180099	       285.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	95912703	        12.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12829462	        93.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 8907876	       135.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   76191	     15419 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7993785	       149.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	640466121	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   82975	     14359 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	19490730	        61.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2369017	       506.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103504	     11685 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   78864	     15224 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	23758890	        60.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2558934	       476.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3390856	       353.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   78342	     15237 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.929 / 0 allocs | 108.6 / 0 allocs | **18.3x** | N/A | N/A | N/A | N/A |
| Enum | 11.80 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.91 / 0 allocs | 1097 / 88 B / 5 allocs | **18.9x** | 934.7 / 0 allocs | **16.1x** | 15640 / 15725 B / 74 allocs | **270.1x** |
| GTE | 5.933 / 0 allocs | 108.5 / 0 allocs | **18.3x** | N/A | N/A | N/A | N/A |
| MinLength | 27.11 / 0 allocs | 125.7 / 0 allocs | **4.6x** | 285.0 / 32 B / 2 allocs | **10.5x** | N/A | N/A |
| UUID | 60.85 / 0 allocs | 476.7 / 0 allocs | **7.8x** | 353.8 / 0 allocs | **5.8x** | 15237 / 15501 B / 74 allocs | **250.4x** |
| MaxItems | 11.52 / 0 allocs | 139.5 / 0 allocs | **12.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 23.15 / 0 allocs | 131.9 / 0 allocs | **5.7x** | 278.3 / 32 B / 2 allocs | **12.0x** | 14920 / 15632 B / 80 allocs | **644.5x** |
| LT | 5.678 / 0 allocs | 103.1 / 0 allocs | **18.2x** | N/A | N/A | N/A | N/A |
| MinItems | 12.62 / 0 allocs | 142.8 / 0 allocs | **11.3x** | N/A | N/A | N/A | N/A |
| Alpha | 15.88 / 0 allocs | 462.5 / 0 allocs | **29.1x** | 109.9 / 0 allocs | **6.9x** | 20285 / 16815 B / 97 allocs | **1277.4x** |
| Required | 10.63 / 0 allocs | 149.9 / 0 allocs | **14.1x** | 1.870 / 0 allocs | **0.2x** | 14359 / 15472 B / 72 allocs | **1350.8x** |
| IPV4 | 40.57 / 0 allocs | 135.9 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 10.68 / 0 allocs | 105.2 / 0 allocs | **9.9x** | 258.8 / 32 B / 2 allocs | **24.2x** | 14884 / 15600 B / 78 allocs | **1393.6x** |
| IPV6 | 92.71 / 0 allocs | 180.7 / 0 allocs | **1.9x** | N/A | N/A | N/A | N/A |
| URL | 61.42 / 0 allocs | 506.6 / 144 B / 1 allocs | **8.2x** | 11685 / 147 B / 1 allocs | **190.2x** | 15224 / 15641 B / 75 allocs | **247.9x** |
| Numeric | 12.47 / 0 allocs | 93.16 / 0 allocs | **7.5x** | 135.2 / 0 allocs | **10.8x** | 15419 / 15533 B / 76 allocs | **1236.5x** |
| GT | 5.901 / 0 allocs | 116.1 / 0 allocs | **19.7x** | 89.51 / 0 allocs | **15.2x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.757 | 0 allocs |
| CELMultipleExpressions | 11.87 | 0 allocs |
| CELBasic | 11.84 | 0 allocs |
| CELCrossField | 7.741 | 0 allocs |
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
