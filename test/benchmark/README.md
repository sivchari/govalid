# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-22  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124456981	         9.643 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2559432	       474.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11212633	       107.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56474	     21081 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644334424	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295824458	         4.056 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295870549	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350318841	         3.428 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21679204	        59.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1094 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1334197	       897.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70305	     17116 ns/op	   15867 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	274868973	         4.377 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481825898	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10814230	       111.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13491120	        89.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482742150	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10560042	       113.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30932727	        39.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8834103	       135.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14102875	        85.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6716197	       178.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160627100	         7.468 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9665055	       124.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4848718	       246.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74793	     16155 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	480784252	         2.497 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10621965	       112.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428392106	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10773942	       113.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	176958726	         6.678 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8066539	       149.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35874386	        33.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8390269	       143.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4150484	       288.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74306	     16277 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202547192	         5.923 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8209869	       146.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52081242	        23.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9742654	       123.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4217425	       283.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	156435510	         7.671 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13222842	        91.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9410758	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72128	     16669 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320535546	         3.741 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7814082	       153.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642477285	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76719	     15639 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20769158	        57.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2474490	       485.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102718	     11687 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71155	     16755 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23212707	        51.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2647113	       453.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3398742	       349.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   72063	     16731 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.801 / 0 allocs | 113.0 / 0 allocs | **40.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.377 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.44 / 0 allocs | 1094 / 88 B / 5 allocs | **18.4x** | 897.7 / 0 allocs | **15.1x** | 17116 / 15867 B / 76 allocs | **288.0x** |
| GTE | 2.491 / 0 allocs | 113.5 / 0 allocs | **45.6x** | N/A | N/A | N/A | N/A |
| MinLength | 23.27 / 0 allocs | 123.3 / 0 allocs | **5.3x** | 283.1 / 32 B / 2 allocs | **12.2x** | N/A | N/A |
| UUID | 51.53 / 0 allocs | 453.5 / 0 allocs | **8.8x** | 349.1 / 0 allocs | **6.8x** | 16731 / 15542 B / 76 allocs | **324.7x** |
| MaxItems | 6.678 / 0 allocs | 149.2 / 0 allocs | **22.3x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.31 / 0 allocs | 143.2 / 0 allocs | **4.3x** | 288.6 / 32 B / 2 allocs | **8.7x** | 16277 / 15648 B / 81 allocs | **488.7x** |
| LT | 2.497 / 0 allocs | 112.8 / 0 allocs | **45.2x** | N/A | N/A | N/A | N/A |
| MinItems | 5.923 / 0 allocs | 146.3 / 0 allocs | **24.7x** | N/A | N/A | N/A | N/A |
| Alpha | 9.643 / 0 allocs | 474.7 / 0 allocs | **49.2x** | 107.6 / 0 allocs | **11.2x** | 21081 / 16937 B / 101 allocs | **2186.1x** |
| Required | 3.741 / 0 allocs | 153.8 / 0 allocs | **41.1x** | 1.868 / 0 allocs | **0.5x** | 15639 / 15488 B / 73 allocs | **4180.4x** |
| IPV4 | 39.03 / 0 allocs | 135.8 / 0 allocs | **3.5x** | N/A | N/A | N/A | N/A |
| Length | 7.468 / 0 allocs | 124.0 / 0 allocs | **16.6x** | 246.4 / 32 B / 2 allocs | **33.0x** | 16155 / 15616 B / 79 allocs | **2163.2x** |
| IPV6 | 85.62 / 0 allocs | 178.8 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 57.78 / 0 allocs | 485.0 / 144 B / 1 allocs | **8.4x** | 11687 / 145 B / 1 allocs | **202.3x** | 16755 / 15681 B / 77 allocs | **290.0x** |
| Numeric | 7.671 / 0 allocs | 91.00 / 0 allocs | **11.9x** | 128.3 / 0 allocs | **16.7x** | 16669 / 15574 B / 78 allocs | **2173.0x** |
| GT | 2.491 / 0 allocs | 111.7 / 0 allocs | **44.8x** | 89.11 / 0 allocs | **35.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.866 | 0 allocs |
| CELMultipleExpressions | 4.056 | 0 allocs |
| CELBasic | 4.054 | 0 allocs |
| CELCrossField | 3.428 | 0 allocs |
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
