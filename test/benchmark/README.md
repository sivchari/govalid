# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-22  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124276400	         9.657 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2482436	       476.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11203899	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55630	     21382 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	646708596	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296137621	         4.050 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296103232	         4.061 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	348851626	         3.439 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	20467652	        57.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1100 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1317594	       910.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   67588	     17457 ns/op	   15860 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	274861186	         4.365 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481331074	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11271708	       106.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13475611	        89.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481971484	         2.495 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11018852	       109.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30746563	        39.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9243681	       130.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14018169	        85.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6938077	       173.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160733118	         7.467 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10030114	       120.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4789377	       251.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73063	     16379 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	480606504	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11077809	       108.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427536979	         2.803 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11112009	       107.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183370147	         6.560 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8228246	       145.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35853850	        33.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8653354	       138.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4145924	       289.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   70778	     16714 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202654381	         5.919 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8493715	       140.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52109704	        23.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10085784	       118.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4171924	       285.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	154121054	         7.751 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13118400	        91.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9409929	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   69622	     17013 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320574645	         3.738 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7933803	       150.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641864652	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75530	     16160 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20769007	        57.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2458605	       495.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   96816	     12391 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   69082	     17986 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	22527466	        53.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2492655	       476.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3300966	       364.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   64062	     18906 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.803 / 0 allocs | 107.3 / 0 allocs | **38.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.365 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.42 / 0 allocs | 1100 / 89 B / 5 allocs | **19.2x** | 910.2 / 0 allocs | **15.9x** | 17457 / 15860 B / 76 allocs | **304.0x** |
| GTE | 2.495 / 0 allocs | 109.2 / 0 allocs | **43.8x** | N/A | N/A | N/A | N/A |
| MinLength | 23.06 / 0 allocs | 118.5 / 0 allocs | **5.1x** | 285.6 / 32 B / 2 allocs | **12.4x** | N/A | N/A |
| UUID | 53.82 / 0 allocs | 476.1 / 0 allocs | **8.8x** | 364.2 / 0 allocs | **6.8x** | 18906 / 15542 B / 76 allocs | **351.3x** |
| MaxItems | 6.560 / 0 allocs | 145.7 / 0 allocs | **22.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.36 / 0 allocs | 138.9 / 0 allocs | **4.2x** | 289.2 / 32 B / 2 allocs | **8.7x** | 16714 / 15648 B / 81 allocs | **501.0x** |
| LT | 2.493 / 0 allocs | 108.7 / 0 allocs | **43.6x** | N/A | N/A | N/A | N/A |
| MinItems | 5.919 / 0 allocs | 140.9 / 0 allocs | **23.8x** | N/A | N/A | N/A | N/A |
| Alpha | 9.657 / 0 allocs | 476.8 / 0 allocs | **49.4x** | 108.3 / 0 allocs | **11.2x** | 21382 / 16937 B / 101 allocs | **2214.1x** |
| Required | 3.738 / 0 allocs | 150.6 / 0 allocs | **40.3x** | 1.869 / 0 allocs | **0.5x** | 16160 / 15488 B / 73 allocs | **4323.2x** |
| IPV4 | 39.06 / 0 allocs | 130.9 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 7.467 / 0 allocs | 120.4 / 0 allocs | **16.1x** | 251.2 / 32 B / 2 allocs | **33.6x** | 16379 / 15616 B / 79 allocs | **2193.5x** |
| IPV6 | 85.73 / 0 allocs | 173.0 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.67 / 0 allocs | 495.1 / 144 B / 1 allocs | **8.6x** | 12391 / 145 B / 1 allocs | **214.9x** | 17986 / 15681 B / 77 allocs | **311.9x** |
| Numeric | 7.751 / 0 allocs | 91.34 / 0 allocs | **11.8x** | 128.2 / 0 allocs | **16.5x** | 17013 / 15574 B / 78 allocs | **2194.9x** |
| GT | 2.493 / 0 allocs | 106.7 / 0 allocs | **42.8x** | 89.26 / 0 allocs | **35.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.868 | 0 allocs |
| CELMultipleExpressions | 4.050 | 0 allocs |
| CELBasic | 4.061 | 0 allocs |
| CELCrossField | 3.439 | 0 allocs |
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
