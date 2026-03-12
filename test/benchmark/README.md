# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-12  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	81194758	        14.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2599200	       470.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10783101	       113.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60438	     19946 ns/op	   16936 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	268923538	         4.468 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	148087514	         8.105 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	18536139	        61.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1090 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1335423	       925.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   74668	     16023 ns/op	   15805 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	92055572	        13.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	213715335	         5.617 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10666501	       113.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13631730	        88.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	212495176	         5.625 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10852717	       110.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27481424	        44.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8745574	       136.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13852765	        89.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6503481	       184.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	124057530	         9.670 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10843324	       111.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4827926	       249.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   79939	     14957 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	213757254	         5.611 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10799658	       111.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202773271	         5.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10175646	       119.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	98892103	        12.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8343393	       144.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	32832657	        36.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9473143	       132.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4250844	       285.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   79980	     15108 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	98096037	        11.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8752918	       138.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48127003	        24.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9375115	       129.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4283517	       279.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	94367406	        12.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13729981	        88.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9358312	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   77434	     15454 ns/op	   15573 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	127039024	         9.426 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8254778	       145.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642189090	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   81613	     14456 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19586299	        61.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2525851	       474.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  101272	     11909 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   76239	     15596 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	18585028	        64.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2447280	       478.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3360568	       359.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   76026	     15599 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.918 / 0 allocs | 119.1 / 0 allocs | **20.1x** | N/A | N/A | N/A | N/A |
| Enum | 13.24 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 61.90 / 0 allocs | 1090 / 88 B / 5 allocs | **17.6x** | 925.2 / 0 allocs | **14.9x** | 16023 / 15805 B / 76 allocs | **258.9x** |
| GTE | 5.625 / 0 allocs | 110.4 / 0 allocs | **19.6x** | N/A | N/A | N/A | N/A |
| MinLength | 24.90 / 0 allocs | 129.4 / 0 allocs | **5.2x** | 279.7 / 32 B / 2 allocs | **11.2x** | N/A | N/A |
| UUID | 64.71 / 0 allocs | 478.7 / 0 allocs | **7.4x** | 359.9 / 0 allocs | **5.6x** | 15599 / 15541 B / 76 allocs | **241.1x** |
| MaxItems | 12.14 / 0 allocs | 144.6 / 0 allocs | **11.9x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.48 / 0 allocs | 132.2 / 0 allocs | **3.6x** | 285.9 / 32 B / 2 allocs | **7.8x** | 15108 / 15648 B / 81 allocs | **414.1x** |
| LT | 5.611 / 0 allocs | 111.9 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| MinItems | 11.86 / 0 allocs | 138.1 / 0 allocs | **11.6x** | N/A | N/A | N/A | N/A |
| Alpha | 14.65 / 0 allocs | 470.6 / 0 allocs | **32.1x** | 113.2 / 0 allocs | **7.7x** | 19946 / 16936 B / 101 allocs | **1361.5x** |
| Required | 9.426 / 0 allocs | 145.2 / 0 allocs | **15.4x** | 1.870 / 0 allocs | **0.2x** | 14456 / 15488 B / 73 allocs | **1533.6x** |
| IPV4 | 44.30 / 0 allocs | 136.3 / 0 allocs | **3.1x** | N/A | N/A | N/A | N/A |
| Length | 9.670 / 0 allocs | 111.9 / 0 allocs | **11.6x** | 249.6 / 32 B / 2 allocs | **25.8x** | 14957 / 15616 B / 79 allocs | **1546.7x** |
| IPV6 | 89.29 / 0 allocs | 184.5 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 61.15 / 0 allocs | 474.6 / 144 B / 1 allocs | **7.8x** | 11909 / 147 B / 1 allocs | **194.8x** | 15596 / 15681 B / 77 allocs | **255.0x** |
| Numeric | 12.46 / 0 allocs | 88.28 / 0 allocs | **7.1x** | 128.2 / 0 allocs | **10.3x** | 15454 / 15573 B / 78 allocs | **1240.3x** |
| GT | 5.617 / 0 allocs | 113.3 / 0 allocs | **20.2x** | 88.65 / 0 allocs | **15.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.468 | 0 allocs |
| CELMultipleExpressions | 11.79 | 0 allocs |
| CELBasic | 11.74 | 0 allocs |
| CELCrossField | 8.105 | 0 allocs |
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
