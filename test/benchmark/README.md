# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-10  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	125193379	         9.655 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 3121952	       384.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	12824558	        94.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60924	     19671 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	488678674	         2.456 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	378355960	         3.174 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	377611681	         3.188 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	461745411	         2.598 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	655460634	         1.833 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21502234	        55.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1037 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1552857	       772.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   74358	     16444 ns/op	   15874 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	336469676	         3.550 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	565115967	         2.145 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11393124	       103.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	16030677	        74.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	564473594	         2.125 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11984750	       100.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	36362810	        32.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	10264686	       116.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	15277354	        78.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 7084354	       169.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	207233251	         5.780 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11842899	       101.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 5296814	       223.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   79150	     15057 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	565580797	         2.145 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11989784	        99.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	565378693	         2.128 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11338321	       101.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	244319166	         4.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 9332382	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	40017944	        30.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9522944	       126.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4562193	       257.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   79050	     15072 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	259565760	         4.623 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 9352672	       128.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	66561760	        17.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10688041	       112.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4717188	       256.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	180410359	         6.651 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	15436080	        77.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	10437544	       110.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   75963	     15463 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	415191410	         2.891 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8876875	       134.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	1000000000	         1.168 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   81238	     15254 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19341122	        56.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2567788	       462.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  104428	     11542 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   77949	     15680 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	32794706	        36.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2839377	       422.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3522171	       338.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   75326	     16443 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.128 / 0 allocs | 101.9 / 0 allocs | **47.9x** | N/A | N/A | N/A | N/A |
| Enum | 3.550 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.83 / 0 allocs | 1037 / 89 B / 5 allocs | **18.6x** | 772.9 / 0 allocs | **13.8x** | 16444 / 15874 B / 76 allocs | **294.5x** |
| GTE | 2.125 / 0 allocs | 100.0 / 0 allocs | **47.1x** | N/A | N/A | N/A | N/A |
| MinLength | 17.90 / 0 allocs | 112.2 / 0 allocs | **6.3x** | 256.8 / 32 B / 2 allocs | **14.3x** | N/A | N/A |
| UUID | 36.45 / 0 allocs | 422.1 / 0 allocs | **11.6x** | 338.0 / 0 allocs | **9.3x** | 16443 / 15542 B / 76 allocs | **451.1x** |
| MaxItems | 4.918 / 0 allocs | 128.3 / 0 allocs | **26.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 30.01 / 0 allocs | 126.1 / 0 allocs | **4.2x** | 257.0 / 32 B / 2 allocs | **8.6x** | 15072 / 15648 B / 81 allocs | **502.2x** |
| LT | 2.145 / 0 allocs | 99.91 / 0 allocs | **46.6x** | N/A | N/A | N/A | N/A |
| MinItems | 4.623 / 0 allocs | 128.7 / 0 allocs | **27.8x** | N/A | N/A | N/A | N/A |
| Alpha | 9.655 / 0 allocs | 384.5 / 0 allocs | **39.8x** | 94.27 / 0 allocs | **9.8x** | 19671 / 16937 B / 101 allocs | **2037.4x** |
| Required | 2.891 / 0 allocs | 134.2 / 0 allocs | **46.4x** | 1.168 / 0 allocs | **0.4x** | 15254 / 15488 B / 73 allocs | **5276.4x** |
| IPV4 | 32.14 / 0 allocs | 116.9 / 0 allocs | **3.6x** | N/A | N/A | N/A | N/A |
| Length | 5.780 / 0 allocs | 101.2 / 0 allocs | **17.5x** | 223.5 / 32 B / 2 allocs | **38.7x** | 15057 / 15616 B / 79 allocs | **2605.0x** |
| IPV6 | 78.69 / 0 allocs | 169.2 / 0 allocs | **2.2x** | N/A | N/A | N/A | N/A |
| URL | 56.60 / 0 allocs | 462.9 / 144 B / 1 allocs | **8.2x** | 11542 / 145 B / 1 allocs | **203.9x** | 15680 / 15681 B / 77 allocs | **277.0x** |
| Numeric | 6.651 / 0 allocs | 77.79 / 0 allocs | **11.7x** | 110.5 / 0 allocs | **16.6x** | 15463 / 15574 B / 78 allocs | **2324.9x** |
| GT | 2.145 / 0 allocs | 103.7 / 0 allocs | **48.3x** | 74.56 / 0 allocs | **34.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.456 | 0 allocs |
| CELMultipleExpressions | 3.174 | 0 allocs |
| CELBasic | 3.188 | 0 allocs |
| CELCrossField | 2.598 | 0 allocs |
| CELStringLength | 1.113 | 0 allocs |
| CELNumericComparison | 1.833 | 0 allocs |

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
