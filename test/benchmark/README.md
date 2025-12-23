# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-23  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124355292	         9.650 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2584806	       468.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11265663	       111.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55808	     21265 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	652392970	         1.839 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295890828	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296448591	         4.051 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320347861	         3.746 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21594559	        55.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1097 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1358547	       883.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68935	     17353 ns/op	   15841 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	255974263	         4.689 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	426230949	         2.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11298028	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13475900	        89.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	428471244	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11030271	       109.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30686299	        39.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9243440	       130.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14079445	        85.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6608238	       175.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	148025068	         8.108 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9877158	       120.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4814850	       249.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73995	     16521 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	408944334	         2.845 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11398695	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	384084350	         3.122 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11231752	       107.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202633338	         5.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8225138	       145.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37641312	        31.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8616808	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4199289	       286.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72247	     16466 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	174827569	         6.858 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8462271	       142.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48127006	        24.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10063514	       118.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4194944	       289.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	122080249	         9.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13735078	        87.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9430269	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71512	     16902 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296658727	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8020395	       151.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642923924	         1.887 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75424	     15846 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20744756	        59.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2494419	       478.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103516	     12001 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70494	     16979 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	22926351	        51.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2682464	       456.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3426110	       352.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70258	     17709 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.122 / 0 allocs | 107.5 / 0 allocs | **34.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.689 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.53 / 0 allocs | 1097 / 89 B / 5 allocs | **19.8x** | 883.9 / 0 allocs | **15.9x** | 17353 / 15841 B / 76 allocs | **312.5x** |
| GTE | 2.802 / 0 allocs | 109.0 / 0 allocs | **38.9x** | N/A | N/A | N/A | N/A |
| MinLength | 24.95 / 0 allocs | 118.8 / 0 allocs | **4.8x** | 289.2 / 32 B / 2 allocs | **11.6x** | N/A | N/A |
| UUID | 51.47 / 0 allocs | 456.0 / 0 allocs | **8.9x** | 352.8 / 0 allocs | **6.9x** | 17709 / 15542 B / 76 allocs | **344.1x** |
| MaxItems | 5.917 / 0 allocs | 145.8 / 0 allocs | **24.6x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.74 / 0 allocs | 139.3 / 0 allocs | **4.4x** | 286.4 / 32 B / 2 allocs | **9.0x** | 16466 / 15648 B / 81 allocs | **518.8x** |
| LT | 2.845 / 0 allocs | 105.6 / 0 allocs | **37.1x** | N/A | N/A | N/A | N/A |
| MinItems | 6.858 / 0 allocs | 142.6 / 0 allocs | **20.8x** | N/A | N/A | N/A | N/A |
| Alpha | 9.650 / 0 allocs | 468.0 / 0 allocs | **48.5x** | 111.4 / 0 allocs | **11.5x** | 21265 / 16937 B / 101 allocs | **2203.6x** |
| Required | 4.048 / 0 allocs | 151.2 / 0 allocs | **37.4x** | 1.887 / 0 allocs | **0.5x** | 15846 / 15488 B / 73 allocs | **3914.5x** |
| IPV4 | 39.10 / 0 allocs | 130.3 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.108 / 0 allocs | 120.8 / 0 allocs | **14.9x** | 249.6 / 32 B / 2 allocs | **30.8x** | 16521 / 15616 B / 79 allocs | **2037.6x** |
| IPV6 | 85.20 / 0 allocs | 175.3 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 59.10 / 0 allocs | 478.8 / 144 B / 1 allocs | **8.1x** | 12001 / 146 B / 1 allocs | **203.1x** | 16979 / 15681 B / 77 allocs | **287.3x** |
| Numeric | 9.800 / 0 allocs | 87.57 / 0 allocs | **8.9x** | 128.3 / 0 allocs | **13.1x** | 16902 / 15574 B / 78 allocs | **1724.7x** |
| GT | 2.804 / 0 allocs | 106.6 / 0 allocs | **38.0x** | 89.10 / 0 allocs | **31.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.839 | 0 allocs |
| CELMultipleExpressions | 4.054 | 0 allocs |
| CELBasic | 4.051 | 0 allocs |
| CELCrossField | 3.746 | 0 allocs |
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
