# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-02-06  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124244052	         9.658 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2574770	       465.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11246246	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55886	     21302 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	653046190	         1.838 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295690069	         4.055 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295432585	         4.057 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320831044	         3.741 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21720728	        57.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1093 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1357826	       883.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69712	     17201 ns/op	   15861 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256421875	         4.676 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	429008860	         2.799 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11307480	       106.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13651636	        88.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	428110674	         2.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10977746	       109.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30421262	        39.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9219936	       130.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14149879	        86.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6923358	       173.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147482942	         8.146 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9916951	       120.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4874522	       246.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74517	     16196 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	427453000	         2.803 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11320039	       105.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	384932121	         3.120 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11232658	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202849080	         5.928 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8261325	       143.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36779170	        31.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8600325	       139.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4250743	       283.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73729	     16274 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	174974205	         6.855 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8450204	       141.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48091350	        24.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10084492	       118.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4123514	       285.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124391444	         9.646 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13787421	        87.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9439648	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71775	     16682 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296215800	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7990176	       150.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641541012	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76216	     15584 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20751493	        57.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2507056	       477.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103096	     11708 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70588	     16831 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	22805840	        51.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2645035	       460.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3420470	       347.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71704	     16731 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.120 / 0 allocs | 107.4 / 0 allocs | **34.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.676 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.33 / 0 allocs | 1093 / 88 B / 5 allocs | **19.1x** | 883.7 / 0 allocs | **15.4x** | 17201 / 15861 B / 76 allocs | **300.0x** |
| GTE | 2.800 / 0 allocs | 109.2 / 0 allocs | **39.0x** | N/A | N/A | N/A | N/A |
| MinLength | 24.63 / 0 allocs | 118.6 / 0 allocs | **4.8x** | 285.9 / 32 B / 2 allocs | **11.6x** | N/A | N/A |
| UUID | 51.26 / 0 allocs | 460.7 / 0 allocs | **9.0x** | 347.6 / 0 allocs | **6.8x** | 16731 / 15542 B / 76 allocs | **326.4x** |
| MaxItems | 5.928 / 0 allocs | 143.3 / 0 allocs | **24.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.76 / 0 allocs | 139.1 / 0 allocs | **4.4x** | 283.6 / 32 B / 2 allocs | **8.9x** | 16274 / 15648 B / 81 allocs | **512.4x** |
| LT | 2.803 / 0 allocs | 105.4 / 0 allocs | **37.6x** | N/A | N/A | N/A | N/A |
| MinItems | 6.855 / 0 allocs | 141.1 / 0 allocs | **20.6x** | N/A | N/A | N/A | N/A |
| Alpha | 9.658 / 0 allocs | 465.7 / 0 allocs | **48.2x** | 108.3 / 0 allocs | **11.2x** | 21302 / 16937 B / 101 allocs | **2205.6x** |
| Required | 4.053 / 0 allocs | 150.2 / 0 allocs | **37.1x** | 1.866 / 0 allocs | **0.5x** | 15584 / 15488 B / 73 allocs | **3845.1x** |
| IPV4 | 39.12 / 0 allocs | 130.3 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.146 / 0 allocs | 120.3 / 0 allocs | **14.8x** | 246.1 / 32 B / 2 allocs | **30.2x** | 16196 / 15616 B / 79 allocs | **1988.2x** |
| IPV6 | 86.38 / 0 allocs | 173.5 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.70 / 0 allocs | 477.2 / 144 B / 1 allocs | **8.3x** | 11708 / 145 B / 1 allocs | **202.9x** | 16831 / 15681 B / 77 allocs | **291.7x** |
| Numeric | 9.646 / 0 allocs | 87.67 / 0 allocs | **9.1x** | 128.3 / 0 allocs | **13.3x** | 16682 / 15574 B / 78 allocs | **1729.4x** |
| GT | 2.799 / 0 allocs | 106.0 / 0 allocs | **37.9x** | 88.27 / 0 allocs | **31.5x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.838 | 0 allocs |
| CELMultipleExpressions | 4.055 | 0 allocs |
| CELBasic | 4.057 | 0 allocs |
| CELCrossField | 3.741 | 0 allocs |
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
