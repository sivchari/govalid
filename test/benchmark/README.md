# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-22  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124423879	         9.644 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2564954	       471.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11205561	       110.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55783	     21230 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	642061548	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296313268	         4.061 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296024510	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349991900	         3.430 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21683330	        55.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1097 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1349887	       892.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68800	     17216 ns/op	   15842 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275366841	         4.360 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481757134	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11297882	       106.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13643383	        88.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482172960	         2.488 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11050692	       108.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30740601	        39.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9245656	       129.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14016684	        86.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6932856	       173.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160831354	         7.461 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9967629	       121.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4856934	       246.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74086	     16112 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	478304829	         2.494 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11097328	       108.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428274702	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11216644	       107.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183697633	         6.531 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8319944	       144.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36007594	        33.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8652615	       138.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4177942	       286.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73610	     16315 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202877604	         5.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8423554	       142.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52116444	        23.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10084284	       118.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4245200	       283.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	156529479	         7.654 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13153213	        91.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9445470	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72187	     16643 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320900931	         3.735 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8032248	       149.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642703068	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76389	     15617 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20734540	        57.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2488851	       480.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102206	     11783 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71042	     16715 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23148194	        52.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2686772	       447.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3576402	       338.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70720	     16783 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.802 / 0 allocs | 107.8 / 0 allocs | **38.5x** | N/A | N/A | N/A | N/A |
| Enum | 4.360 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.66 / 0 allocs | 1097 / 89 B / 5 allocs | **19.7x** | 892.2 / 0 allocs | **16.0x** | 17216 / 15842 B / 76 allocs | **309.3x** |
| GTE | 2.488 / 0 allocs | 108.6 / 0 allocs | **43.6x** | N/A | N/A | N/A | N/A |
| MinLength | 23.03 / 0 allocs | 118.8 / 0 allocs | **5.2x** | 283.0 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 52.65 / 0 allocs | 447.0 / 0 allocs | **8.5x** | 338.4 / 0 allocs | **6.4x** | 16783 / 15542 B / 76 allocs | **318.8x** |
| MaxItems | 6.531 / 0 allocs | 144.3 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.31 / 0 allocs | 138.9 / 0 allocs | **4.2x** | 286.5 / 32 B / 2 allocs | **8.6x** | 16315 / 15648 B / 81 allocs | **489.8x** |
| LT | 2.494 / 0 allocs | 108.5 / 0 allocs | **43.5x** | N/A | N/A | N/A | N/A |
| MinItems | 5.920 / 0 allocs | 142.3 / 0 allocs | **24.0x** | N/A | N/A | N/A | N/A |
| Alpha | 9.644 / 0 allocs | 471.0 / 0 allocs | **48.8x** | 110.1 / 0 allocs | **11.4x** | 21230 / 16937 B / 101 allocs | **2201.4x** |
| Required | 3.735 / 0 allocs | 149.0 / 0 allocs | **39.9x** | 1.867 / 0 allocs | **0.5x** | 15617 / 15488 B / 73 allocs | **4181.3x** |
| IPV4 | 39.08 / 0 allocs | 129.9 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 7.461 / 0 allocs | 121.2 / 0 allocs | **16.2x** | 246.6 / 32 B / 2 allocs | **33.1x** | 16112 / 15616 B / 79 allocs | **2159.5x** |
| IPV6 | 86.21 / 0 allocs | 173.8 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.62 / 0 allocs | 480.2 / 144 B / 1 allocs | **8.3x** | 11783 / 145 B / 1 allocs | **204.5x** | 16715 / 15681 B / 77 allocs | **290.1x** |
| Numeric | 7.654 / 0 allocs | 91.23 / 0 allocs | **11.9x** | 128.3 / 0 allocs | **16.8x** | 16643 / 15574 B / 78 allocs | **2174.4x** |
| GT | 2.489 / 0 allocs | 106.1 / 0 allocs | **42.6x** | 88.17 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.866 | 0 allocs |
| CELMultipleExpressions | 4.061 | 0 allocs |
| CELBasic | 4.054 | 0 allocs |
| CELCrossField | 3.430 | 0 allocs |
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
