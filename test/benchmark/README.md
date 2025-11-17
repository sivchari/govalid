# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-17  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124125379	         9.665 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2581695	       467.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11140642	       109.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56292	     21311 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644632189	         1.860 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296177067	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296042186	         4.059 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	348804260	         3.434 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21020248	        59.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1105 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1328180	       906.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   67412	     17359 ns/op	   15877 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275312101	         4.360 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481342123	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11322378	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13772592	        88.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	479849065	         2.495 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11161747	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30780019	        39.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9031662	       130.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14192851	        85.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6879094	       174.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160754150	         7.464 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10054795	       119.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4866060	       250.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74301	     16063 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481446469	         2.499 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11353246	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427019120	         2.811 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11210246	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183398702	         6.545 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8245578	       145.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35972426	        33.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8647716	       139.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4241796	       282.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74304	     16208 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202518316	         5.998 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8463344	       142.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52068092	        23.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10099827	       118.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4220518	       283.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157261695	         7.650 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13824384	        86.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9450085	       127.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   70737	     16662 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321150451	         3.737 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7923193	       151.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642394806	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76758	     15510 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20564610	        57.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2519265	       475.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103184	     11701 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   72002	     16695 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24544072	        48.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2690095	       451.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3576217	       345.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71818	     16756 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.811 / 0 allocs | 107.4 / 0 allocs | **38.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.360 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.30 / 0 allocs | 1105 / 89 B / 5 allocs | **18.6x** | 906.7 / 0 allocs | **15.3x** | 17359 / 15877 B / 76 allocs | **292.7x** |
| GTE | 2.495 / 0 allocs | 107.7 / 0 allocs | **43.2x** | N/A | N/A | N/A | N/A |
| MinLength | 23.05 / 0 allocs | 118.5 / 0 allocs | **5.1x** | 283.0 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.65 / 0 allocs | 451.3 / 0 allocs | **9.3x** | 345.9 / 0 allocs | **7.1x** | 16756 / 15542 B / 76 allocs | **344.4x** |
| MaxItems | 6.545 / 0 allocs | 145.3 / 0 allocs | **22.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.37 / 0 allocs | 139.1 / 0 allocs | **4.2x** | 282.8 / 32 B / 2 allocs | **8.5x** | 16208 / 15648 B / 81 allocs | **485.7x** |
| LT | 2.499 / 0 allocs | 105.7 / 0 allocs | **42.3x** | N/A | N/A | N/A | N/A |
| MinItems | 5.998 / 0 allocs | 142.4 / 0 allocs | **23.7x** | N/A | N/A | N/A | N/A |
| Alpha | 9.665 / 0 allocs | 467.4 / 0 allocs | **48.4x** | 109.9 / 0 allocs | **11.4x** | 21311 / 16937 B / 101 allocs | **2205.0x** |
| Required | 3.737 / 0 allocs | 151.5 / 0 allocs | **40.5x** | 1.867 / 0 allocs | **0.5x** | 15510 / 15488 B / 73 allocs | **4150.4x** |
| IPV4 | 39.10 / 0 allocs | 130.1 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 7.464 / 0 allocs | 119.1 / 0 allocs | **16.0x** | 250.3 / 32 B / 2 allocs | **33.5x** | 16063 / 15616 B / 79 allocs | **2152.1x** |
| IPV6 | 85.51 / 0 allocs | 174.4 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.78 / 0 allocs | 475.8 / 144 B / 1 allocs | **8.2x** | 11701 / 145 B / 1 allocs | **202.5x** | 16695 / 15681 B / 77 allocs | **288.9x** |
| Numeric | 7.650 / 0 allocs | 86.87 / 0 allocs | **11.4x** | 127.9 / 0 allocs | **16.7x** | 16662 / 15574 B / 78 allocs | **2178.0x** |
| GT | 2.492 / 0 allocs | 106.6 / 0 allocs | **42.8x** | 88.28 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.860 | 0 allocs |
| CELMultipleExpressions | 4.053 | 0 allocs |
| CELBasic | 4.059 | 0 allocs |
| CELCrossField | 3.434 | 0 allocs |
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
