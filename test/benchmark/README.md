# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-25  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.25.0

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	72251019	        16.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2573250	       467.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10954734	       110.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60877	     20250 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	252015238	         4.757 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	154013840	         7.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	20499940	        58.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1077 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1286460	       928.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   75850	     15768 ns/op	   15699 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	100000000	        11.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	202679461	         5.922 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10484720	       114.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13382047	        89.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	202904088	         5.922 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11101618	       108.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	29466584	        40.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8954044	       134.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	12885423	        93.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6722763	       180.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        10.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11469370	       105.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4630777	       257.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80754	     14946 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	214017919	         5.608 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11631589	       103.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202506951	         5.922 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11195806	       107.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	100000000	        11.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 7946836	       140.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	46786770	        23.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9054597	       132.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4291184	       278.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   80151	     15018 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	96136251	        12.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8264692	       144.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	44368798	        27.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9819988	       122.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4200892	       284.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	95510337	        12.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12754965	        93.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9153693	       130.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   78276	     15301 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7955584	       150.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	640698158	         1.871 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   83420	     14231 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	19601956	        61.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2375870	       505.7 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  101206	     11848 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   79201	     15376 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	19283028	        59.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2582965	       466.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3396217	       353.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   77718	     15360 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.922 / 0 allocs | 107.2 / 0 allocs | **18.1x** | N/A | N/A | N/A | N/A |
| Enum | 11.86 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.50 / 0 allocs | 1077 / 88 B / 5 allocs | **18.4x** | 928.0 / 0 allocs | **15.9x** | 15768 / 15699 B / 74 allocs | **269.5x** |
| GTE | 5.922 / 0 allocs | 108.4 / 0 allocs | **18.3x** | N/A | N/A | N/A | N/A |
| MinLength | 27.10 / 0 allocs | 122.4 / 0 allocs | **4.5x** | 284.5 / 32 B / 2 allocs | **10.5x** | N/A | N/A |
| UUID | 59.41 / 0 allocs | 466.0 / 0 allocs | **7.8x** | 353.4 / 0 allocs | **5.9x** | 15360 / 15501 B / 74 allocs | **258.5x** |
| MaxItems | 11.57 / 0 allocs | 140.4 / 0 allocs | **12.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 23.56 / 0 allocs | 132.5 / 0 allocs | **5.6x** | 278.9 / 32 B / 2 allocs | **11.8x** | 15018 / 15632 B / 80 allocs | **637.4x** |
| LT | 5.608 / 0 allocs | 103.1 / 0 allocs | **18.4x** | N/A | N/A | N/A | N/A |
| MinItems | 12.55 / 0 allocs | 144.7 / 0 allocs | **11.5x** | N/A | N/A | N/A | N/A |
| Alpha | 16.07 / 0 allocs | 467.0 / 0 allocs | **29.1x** | 110.3 / 0 allocs | **6.9x** | 20250 / 16815 B / 97 allocs | **1260.1x** |
| Required | 10.59 / 0 allocs | 150.6 / 0 allocs | **14.2x** | 1.871 / 0 allocs | **0.2x** | 14231 / 15472 B / 72 allocs | **1343.8x** |
| IPV4 | 40.50 / 0 allocs | 134.0 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 10.62 / 0 allocs | 105.0 / 0 allocs | **9.9x** | 257.5 / 32 B / 2 allocs | **24.2x** | 14946 / 15600 B / 78 allocs | **1407.3x** |
| IPV6 | 93.07 / 0 allocs | 180.3 / 0 allocs | **1.9x** | N/A | N/A | N/A | N/A |
| URL | 61.17 / 0 allocs | 505.7 / 144 B / 1 allocs | **8.3x** | 11848 / 146 B / 1 allocs | **193.7x** | 15376 / 15641 B / 75 allocs | **251.4x** |
| Numeric | 12.48 / 0 allocs | 93.83 / 0 allocs | **7.5x** | 130.9 / 0 allocs | **10.5x** | 15301 / 15533 B / 76 allocs | **1226.0x** |
| GT | 5.922 / 0 allocs | 114.5 / 0 allocs | **19.3x** | 89.95 / 0 allocs | **15.2x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.757 | 0 allocs |
| CELMultipleExpressions | 11.96 | 0 allocs |
| CELBasic | 11.86 | 0 allocs |
| CELCrossField | 7.800 | 0 allocs |
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
