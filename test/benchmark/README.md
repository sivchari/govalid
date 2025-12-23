# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-23  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124064253	         9.669 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2630595	       463.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11259284	       108.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56982	     20915 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	653075600	         1.838 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295802794	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295179690	         4.056 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320528595	         3.741 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21675168	        55.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1089 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1325092	       905.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69955	     17070 ns/op	   15852 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256358214	         4.674 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	423650546	         2.811 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11297217	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13643685	        88.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	428555276	         2.799 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10927927	       110.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30767257	        39.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9227337	       132.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14140812	        85.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6923648	       173.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	146721222	         8.155 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9946578	       120.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4853818	       246.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74491	     16124 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	428068453	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11400985	       105.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	384772112	         3.116 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11128634	       108.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202822825	         5.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8217931	       146.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37487952	        31.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8601662	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4239240	       282.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73531	     16196 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	175213294	         6.845 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8469361	       141.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48879384	        24.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10139155	       119.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4222854	       283.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124395798	         9.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13788402	        87.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9403443	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72022	     16609 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296589588	         4.046 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8028414	       149.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642034430	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76213	     15536 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20726053	        57.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2526723	       475.4 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103344	     11660 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71892	     16762 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	19561993	        56.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2683383	       448.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3581947	       336.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   72187	     16730 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.116 / 0 allocs | 108.6 / 0 allocs | **34.9x** | N/A | N/A | N/A | N/A |
| Enum | 4.674 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.52 / 0 allocs | 1089 / 89 B / 5 allocs | **19.6x** | 905.5 / 0 allocs | **16.3x** | 17070 / 15852 B / 76 allocs | **307.5x** |
| GTE | 2.799 / 0 allocs | 110.0 / 0 allocs | **39.3x** | N/A | N/A | N/A | N/A |
| MinLength | 24.58 / 0 allocs | 119.3 / 0 allocs | **4.9x** | 283.5 / 32 B / 2 allocs | **11.5x** | N/A | N/A |
| UUID | 56.65 / 0 allocs | 448.1 / 0 allocs | **7.9x** | 336.2 / 0 allocs | **5.9x** | 16730 / 15542 B / 76 allocs | **295.3x** |
| MaxItems | 5.918 / 0 allocs | 146.2 / 0 allocs | **24.7x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.75 / 0 allocs | 139.3 / 0 allocs | **4.4x** | 282.5 / 32 B / 2 allocs | **8.9x** | 16196 / 15648 B / 81 allocs | **510.1x** |
| LT | 2.802 / 0 allocs | 105.3 / 0 allocs | **37.6x** | N/A | N/A | N/A | N/A |
| MinItems | 6.845 / 0 allocs | 141.6 / 0 allocs | **20.7x** | N/A | N/A | N/A | N/A |
| Alpha | 9.669 / 0 allocs | 463.3 / 0 allocs | **47.9x** | 108.1 / 0 allocs | **11.2x** | 20915 / 16937 B / 101 allocs | **2163.1x** |
| Required | 4.046 / 0 allocs | 149.2 / 0 allocs | **36.9x** | 1.867 / 0 allocs | **0.5x** | 15536 / 15488 B / 73 allocs | **3839.8x** |
| IPV4 | 39.16 / 0 allocs | 132.1 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 8.155 / 0 allocs | 120.5 / 0 allocs | **14.8x** | 246.2 / 32 B / 2 allocs | **30.2x** | 16124 / 15616 B / 79 allocs | **1977.2x** |
| IPV6 | 85.40 / 0 allocs | 173.3 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.66 / 0 allocs | 475.4 / 144 B / 1 allocs | **8.2x** | 11660 / 145 B / 1 allocs | **202.2x** | 16762 / 15681 B / 77 allocs | **290.7x** |
| Numeric | 9.645 / 0 allocs | 87.69 / 0 allocs | **9.1x** | 128.3 / 0 allocs | **13.3x** | 16609 / 15574 B / 78 allocs | **1722.0x** |
| GT | 2.811 / 0 allocs | 106.3 / 0 allocs | **37.8x** | 88.15 / 0 allocs | **31.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.838 | 0 allocs |
| CELMultipleExpressions | 4.054 | 0 allocs |
| CELBasic | 4.056 | 0 allocs |
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
