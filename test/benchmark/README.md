# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-25  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	77470124	        14.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2595364	       467.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10907457	       112.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   50590	     20535 ns/op	   16936 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	243013137	         5.001 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	147922168	         8.110 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19390138	        61.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1074 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1335056	       913.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   73970	     16020 ns/op	   15793 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	93179247	        12.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	213942122	         5.640 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10832148	       113.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13525780	        88.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	213756792	         5.613 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10628217	       114.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27565731	        44.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8901099	       134.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13709012	        87.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6550912	       183.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	124362391	         9.650 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10723294	       111.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4785064	       248.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   79360	     15040 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	213230461	         5.619 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11027442	       111.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202580958	         5.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10117772	       118.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	97850500	        12.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8356180	       144.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	32801961	        36.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9535279	       126.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4276148	       280.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   79434	     15133 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	99118652	        11.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8494478	       139.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48157990	        24.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9361723	       130.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4236074	       280.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	96033188	        12.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13566697	        88.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9388500	       127.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   77160	     15541 ns/op	   15573 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	128288610	         9.354 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8269508	       145.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642747696	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   81145	     14562 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19707387	        60.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2526668	       474.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   93154	     12979 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   77209	     15731 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	18481441	        64.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2613668	       463.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3411218	       361.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   76196	     15575 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.925 / 0 allocs | 118.6 / 0 allocs | **20.0x** | N/A | N/A | N/A | N/A |
| Enum | 12.77 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 61.79 / 0 allocs | 1074 / 89 B / 5 allocs | **17.4x** | 913.4 / 0 allocs | **14.8x** | 16020 / 15793 B / 76 allocs | **259.3x** |
| GTE | 5.613 / 0 allocs | 114.4 / 0 allocs | **20.4x** | N/A | N/A | N/A | N/A |
| MinLength | 24.94 / 0 allocs | 130.4 / 0 allocs | **5.2x** | 280.7 / 32 B / 2 allocs | **11.3x** | N/A | N/A |
| UUID | 64.48 / 0 allocs | 463.0 / 0 allocs | **7.2x** | 361.0 / 0 allocs | **5.6x** | 15575 / 15541 B / 76 allocs | **241.5x** |
| MaxItems | 12.15 / 0 allocs | 144.2 / 0 allocs | **11.9x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.51 / 0 allocs | 126.9 / 0 allocs | **3.5x** | 280.0 / 32 B / 2 allocs | **7.7x** | 15133 / 15648 B / 81 allocs | **414.5x** |
| LT | 5.619 / 0 allocs | 111.7 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| MinItems | 11.85 / 0 allocs | 139.5 / 0 allocs | **11.8x** | N/A | N/A | N/A | N/A |
| Alpha | 14.66 / 0 allocs | 467.7 / 0 allocs | **31.9x** | 112.8 / 0 allocs | **7.7x** | 20535 / 16936 B / 101 allocs | **1400.8x** |
| Required | 9.354 / 0 allocs | 145.2 / 0 allocs | **15.5x** | 1.867 / 0 allocs | **0.2x** | 14562 / 15488 B / 73 allocs | **1556.8x** |
| IPV4 | 44.73 / 0 allocs | 134.2 / 0 allocs | **3.0x** | N/A | N/A | N/A | N/A |
| Length | 9.650 / 0 allocs | 111.4 / 0 allocs | **11.5x** | 248.8 / 32 B / 2 allocs | **25.8x** | 15040 / 15616 B / 79 allocs | **1558.5x** |
| IPV6 | 87.44 / 0 allocs | 183.8 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 60.95 / 0 allocs | 474.6 / 144 B / 1 allocs | **7.8x** | 12979 / 145 B / 1 allocs | **212.9x** | 15731 / 15681 B / 77 allocs | **258.1x** |
| Numeric | 12.46 / 0 allocs | 88.39 / 0 allocs | **7.1x** | 127.9 / 0 allocs | **10.3x** | 15541 / 15573 B / 78 allocs | **1247.3x** |
| GT | 5.640 / 0 allocs | 113.2 / 0 allocs | **20.1x** | 88.50 / 0 allocs | **15.7x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 5.001 | 0 allocs |
| CELMultipleExpressions | 11.72 | 0 allocs |
| CELBasic | 11.59 | 0 allocs |
| CELCrossField | 8.110 | 0 allocs |
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
