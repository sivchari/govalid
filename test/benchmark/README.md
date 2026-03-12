# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-12  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	80885136	        14.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2653892	       453.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10427574	       111.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60088	     20274 ns/op	   16935 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	270346882	         4.425 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	148036628	         8.112 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19415353	        61.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1085 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1338508	       896.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   73604	     16208 ns/op	   15790 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	93210297	        12.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	209399163	         5.671 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10744606	       113.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13544278	        88.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	213757500	         5.607 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10781695	       113.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27162374	        43.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8835469	       135.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13970772	        87.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6369579	       188.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	124190707	         9.663 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10888310	       112.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4757002	       249.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80238	     14966 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	214100593	         5.606 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11021941	       111.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202717594	         5.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10277724	       118.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	98952349	        12.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8264679	       146.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	32843818	        36.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9546020	       126.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4266562	       283.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   78253	     15242 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	99836425	        11.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8596260	       140.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	47456164	        25.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9327722	       130.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4249488	       279.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	94324509	        12.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13548079	        88.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9411086	       127.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   76064	     15737 ns/op	   15573 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	128372523	         9.346 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8270394	       145.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641341792	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   81693	     14738 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	18931072	        62.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2482393	       483.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  100204	     11993 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   74856	     16199 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	18371877	        64.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2586094	       466.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3401115	       356.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   75328	     15761 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.920 / 0 allocs | 118.5 / 0 allocs | **20.0x** | N/A | N/A | N/A | N/A |
| Enum | 12.77 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 61.82 / 0 allocs | 1085 / 88 B / 5 allocs | **17.6x** | 896.5 / 0 allocs | **14.5x** | 16208 / 15790 B / 76 allocs | **262.2x** |
| GTE | 5.607 / 0 allocs | 113.0 / 0 allocs | **20.2x** | N/A | N/A | N/A | N/A |
| MinLength | 25.06 / 0 allocs | 130.2 / 0 allocs | **5.2x** | 279.9 / 32 B / 2 allocs | **11.2x** | N/A | N/A |
| UUID | 64.79 / 0 allocs | 466.1 / 0 allocs | **7.2x** | 356.7 / 0 allocs | **5.5x** | 15761 / 15541 B / 76 allocs | **243.3x** |
| MaxItems | 12.16 / 0 allocs | 146.0 / 0 allocs | **12.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.47 / 0 allocs | 126.8 / 0 allocs | **3.5x** | 283.8 / 32 B / 2 allocs | **7.8x** | 15242 / 15648 B / 81 allocs | **417.9x** |
| LT | 5.606 / 0 allocs | 111.8 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| MinItems | 11.84 / 0 allocs | 140.2 / 0 allocs | **11.8x** | N/A | N/A | N/A | N/A |
| Alpha | 14.73 / 0 allocs | 453.5 / 0 allocs | **30.8x** | 111.6 / 0 allocs | **7.6x** | 20274 / 16935 B / 101 allocs | **1376.4x** |
| Required | 9.346 / 0 allocs | 145.7 / 0 allocs | **15.6x** | 1.870 / 0 allocs | **0.2x** | 14738 / 15488 B / 73 allocs | **1576.9x** |
| IPV4 | 43.79 / 0 allocs | 135.0 / 0 allocs | **3.1x** | N/A | N/A | N/A | N/A |
| Length | 9.663 / 0 allocs | 112.2 / 0 allocs | **11.6x** | 249.6 / 32 B / 2 allocs | **25.8x** | 14966 / 15616 B / 79 allocs | **1548.8x** |
| IPV6 | 87.76 / 0 allocs | 188.8 / 0 allocs | **2.2x** | N/A | N/A | N/A | N/A |
| URL | 62.45 / 0 allocs | 483.1 / 144 B / 1 allocs | **7.7x** | 11993 / 146 B / 1 allocs | **192.0x** | 16199 / 15681 B / 77 allocs | **259.4x** |
| Numeric | 12.48 / 0 allocs | 88.23 / 0 allocs | **7.1x** | 127.6 / 0 allocs | **10.2x** | 15737 / 15573 B / 78 allocs | **1261.0x** |
| GT | 5.671 / 0 allocs | 113.8 / 0 allocs | **20.1x** | 88.74 / 0 allocs | **15.6x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.425 | 0 allocs |
| CELMultipleExpressions | 11.57 | 0 allocs |
| CELBasic | 11.54 | 0 allocs |
| CELCrossField | 8.112 | 0 allocs |
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
