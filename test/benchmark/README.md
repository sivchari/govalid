# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-06-12  
**Platform:** Linux 6.17.0-1018-azure x86_64  
**Go version:** go1.25.6

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	79652227	        14.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2701729	       444.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11331501	       108.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   58498	     20756 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	272693979	         4.401 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        12.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	95216602	        11.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	146839237	         8.185 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19678999	        57.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	  928532	      1125 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1306128	       915.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70854	     16573 ns/op	   15723 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	90957759	        12.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	210714984	         5.701 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10748542	       113.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	11598748	        89.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	209123761	         5.726 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10270959	       117.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27181270	        44.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9331240	       130.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13436925	        88.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6517123	       182.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	123504213	         9.706 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11646156	       105.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4608481	       257.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   71282	     15834 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	198823924	         6.020 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10535649	       109.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	209406470	         5.719 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10439433	       113.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	89915113	        12.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8492098	       139.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	29751583	        36.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8165428	       143.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4315942	       274.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   77078	     15400 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	95178714	        12.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8495888	       139.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	53881240	        19.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 8929130	       131.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4205144	       285.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	108381620	        11.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13255618	        89.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9508006	       127.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   75032	     15826 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	127014415	         9.438 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8287932	       145.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	544694235	         2.199 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   79591	     15261 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	18268688	        64.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2335768	       511.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   94825	     12427 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   75258	     15968 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	18019824	        64.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2637602	       458.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3323529	       360.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   73168	     16131 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.719 / 0 allocs | 113.4 / 0 allocs | **19.8x** | N/A | N/A | N/A | N/A |
| Enum | 12.96 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.59 / 0 allocs | 1125 / 88 B / 5 allocs | **19.5x** | 915.4 / 0 allocs | **15.9x** | 16573 / 15723 B / 74 allocs | **287.8x** |
| GTE | 5.726 / 0 allocs | 117.4 / 0 allocs | **20.5x** | N/A | N/A | N/A | N/A |
| MinLength | 19.89 / 0 allocs | 131.4 / 0 allocs | **6.6x** | 285.7 / 32 B / 2 allocs | **14.4x** | N/A | N/A |
| UUID | 64.47 / 0 allocs | 458.0 / 0 allocs | **7.1x** | 360.9 / 0 allocs | **5.6x** | 16131 / 15501 B / 74 allocs | **250.2x** |
| MaxItems | 12.19 / 0 allocs | 139.8 / 0 allocs | **11.5x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.64 / 0 allocs | 143.2 / 0 allocs | **3.9x** | 274.9 / 32 B / 2 allocs | **7.5x** | 15400 / 15632 B / 80 allocs | **420.3x** |
| LT | 6.020 / 0 allocs | 109.2 / 0 allocs | **18.1x** | N/A | N/A | N/A | N/A |
| MinItems | 12.52 / 0 allocs | 139.9 / 0 allocs | **11.2x** | N/A | N/A | N/A | N/A |
| Alpha | 14.67 / 0 allocs | 444.6 / 0 allocs | **30.3x** | 108.2 / 0 allocs | **7.4x** | 20756 / 16815 B / 97 allocs | **1414.9x** |
| Required | 9.438 / 0 allocs | 145.6 / 0 allocs | **15.4x** | 2.199 / 0 allocs | **0.2x** | 15261 / 15472 B / 72 allocs | **1617.0x** |
| IPV4 | 44.47 / 0 allocs | 130.3 / 0 allocs | **2.9x** | N/A | N/A | N/A | N/A |
| Length | 9.706 / 0 allocs | 105.2 / 0 allocs | **10.8x** | 257.7 / 32 B / 2 allocs | **26.6x** | 15834 / 15600 B / 78 allocs | **1631.4x** |
| IPV6 | 88.89 / 0 allocs | 182.7 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 64.94 / 0 allocs | 511.9 / 144 B / 1 allocs | **7.9x** | 12427 / 146 B / 1 allocs | **191.4x** | 15968 / 15641 B / 75 allocs | **245.9x** |
| Numeric | 11.05 / 0 allocs | 89.38 / 0 allocs | **8.1x** | 127.8 / 0 allocs | **11.6x** | 15826 / 15533 B / 76 allocs | **1432.2x** |
| GT | 5.701 / 0 allocs | 113.1 / 0 allocs | **19.8x** | 89.93 / 0 allocs | **15.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.401 | 0 allocs |
| CELMultipleExpressions | 12.09 | 0 allocs |
| CELBasic | 11.88 | 0 allocs |
| CELCrossField | 8.185 | 0 allocs |
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
