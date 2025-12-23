# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-23  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124384980	         9.648 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2612270	       464.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11197273	       109.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   48334	     21860 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	650508447	         1.846 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	294371932	         4.070 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296036956	         4.055 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320833008	         3.738 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21591075	        59.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1092 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1347546	       889.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69918	     17187 ns/op	   15855 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256388407	         4.695 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427955989	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11309092	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13637761	        88.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	427822053	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10977272	       109.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30705860	        39.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9237396	       130.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14135746	        85.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6927046	       173.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	148414354	         8.086 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10036760	       119.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4889744	       246.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74679	     16113 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	425610280	         2.815 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11360358	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385750306	         3.110 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11230504	       107.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202708632	         5.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8159312	       147.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37720953	        31.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8628373	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4231396	       284.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73700	     16226 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	175270522	         6.845 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8493165	       141.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48745988	        24.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10111749	       118.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4245134	       282.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124120054	         9.670 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13667767	        87.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9412920	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71564	     16668 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296420181	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8023750	       149.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642874198	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76576	     15716 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20796403	        57.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2513739	       475.7 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  100846	     11837 ns/op	     149 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71361	     16862 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	20463828	        55.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2688624	       447.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3466704	       351.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71126	     16808 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.110 / 0 allocs | 107.0 / 0 allocs | **34.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.695 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.61 / 0 allocs | 1092 / 89 B / 5 allocs | **18.3x** | 889.1 / 0 allocs | **14.9x** | 17187 / 15855 B / 76 allocs | **288.3x** |
| GTE | 2.801 / 0 allocs | 109.4 / 0 allocs | **39.1x** | N/A | N/A | N/A | N/A |
| MinLength | 24.63 / 0 allocs | 118.5 / 0 allocs | **4.8x** | 282.3 / 32 B / 2 allocs | **11.5x** | N/A | N/A |
| UUID | 55.24 / 0 allocs | 447.5 / 0 allocs | **8.1x** | 351.9 / 0 allocs | **6.4x** | 16808 / 15542 B / 76 allocs | **304.3x** |
| MaxItems | 5.917 / 0 allocs | 147.3 / 0 allocs | **24.9x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.76 / 0 allocs | 139.3 / 0 allocs | **4.4x** | 284.7 / 32 B / 2 allocs | **9.0x** | 16226 / 15648 B / 81 allocs | **510.9x** |
| LT | 2.815 / 0 allocs | 105.6 / 0 allocs | **37.5x** | N/A | N/A | N/A | N/A |
| MinItems | 6.845 / 0 allocs | 141.4 / 0 allocs | **20.7x** | N/A | N/A | N/A | N/A |
| Alpha | 9.648 / 0 allocs | 464.3 / 0 allocs | **48.1x** | 109.3 / 0 allocs | **11.3x** | 21860 / 16937 B / 101 allocs | **2265.8x** |
| Required | 4.052 / 0 allocs | 149.0 / 0 allocs | **36.8x** | 1.869 / 0 allocs | **0.5x** | 15716 / 15488 B / 73 allocs | **3878.6x** |
| IPV4 | 39.07 / 0 allocs | 130.0 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.086 / 0 allocs | 119.9 / 0 allocs | **14.8x** | 246.1 / 32 B / 2 allocs | **30.4x** | 16113 / 15616 B / 79 allocs | **1992.7x** |
| IPV6 | 85.34 / 0 allocs | 173.1 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.68 / 0 allocs | 475.7 / 144 B / 1 allocs | **8.2x** | 11837 / 149 B / 1 allocs | **205.2x** | 16862 / 15681 B / 77 allocs | **292.3x** |
| Numeric | 9.670 / 0 allocs | 87.24 / 0 allocs | **9.0x** | 128.3 / 0 allocs | **13.3x** | 16668 / 15574 B / 78 allocs | **1723.7x** |
| GT | 2.801 / 0 allocs | 106.3 / 0 allocs | **38.0x** | 88.16 / 0 allocs | **31.5x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.846 | 0 allocs |
| CELMultipleExpressions | 4.070 | 0 allocs |
| CELBasic | 4.055 | 0 allocs |
| CELCrossField | 3.738 | 0 allocs |
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
