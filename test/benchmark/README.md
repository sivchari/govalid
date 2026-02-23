# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-02-23  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	75546541	        15.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 3057470	       393.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11428159	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   59348	     19773 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	220183363	         4.714 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	97273275	        12.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	97063662	        12.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	149037564	         8.067 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.063 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19130654	        62.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1011 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1348194	       889.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   74112	     16085 ns/op	   15867 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	96120434	        12.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	212249635	         5.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	12346194	        97.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13680603	        88.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	212989658	         5.658 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11910116	       100.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30019478	        39.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	10992489	       109.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	12793209	       106.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 7288011	       175.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	100000000	        11.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	12732373	        93.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 5358958	       225.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   79371	     15125 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	212161058	         5.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	12268012	        98.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	200273530	         5.948 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11712378	       102.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	99457154	        12.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 9891879	       121.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	42996344	        26.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9399159	       133.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4566849	       262.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   78428	     15322 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	96804092	        12.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 9845283	       121.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	41434854	        28.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	11176332	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4565250	       261.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	91143044	        13.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	15662756	        77.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9312591	       131.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   75008	     15693 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	100000000	        10.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 9178411	       130.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	567217902	         2.119 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   80991	     14614 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19102528	        63.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2844889	       420.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  109333	     10987 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   75972	     15772 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	20905570	        62.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2913214	       417.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3668060	       329.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   74850	     15946 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.948 / 0 allocs | 102.6 / 0 allocs | **17.2x** | N/A | N/A | N/A | N/A |
| Enum | 12.33 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 62.63 / 0 allocs | 1011 / 89 B / 5 allocs | **16.1x** | 889.3 / 0 allocs | **14.2x** | 16085 / 15867 B / 76 allocs | **256.8x** |
| GTE | 5.658 / 0 allocs | 100.8 / 0 allocs | **17.8x** | N/A | N/A | N/A | N/A |
| MinLength | 28.92 / 0 allocs | 108.0 / 0 allocs | **3.7x** | 261.8 / 32 B / 2 allocs | **9.1x** | N/A | N/A |
| UUID | 62.62 / 0 allocs | 417.9 / 0 allocs | **6.7x** | 329.2 / 0 allocs | **5.3x** | 15946 / 15542 B / 76 allocs | **254.6x** |
| MaxItems | 12.67 / 0 allocs | 121.7 / 0 allocs | **9.6x** | N/A | N/A | N/A | N/A |
| MaxLength | 26.06 / 0 allocs | 133.0 / 0 allocs | **5.1x** | 262.3 / 32 B / 2 allocs | **10.1x** | 15322 / 15648 B / 81 allocs | **588.0x** |
| LT | 5.645 / 0 allocs | 98.25 / 0 allocs | **17.4x** | N/A | N/A | N/A | N/A |
| MinItems | 12.28 / 0 allocs | 121.9 / 0 allocs | **9.9x** | N/A | N/A | N/A | N/A |
| Alpha | 15.85 / 0 allocs | 393.4 / 0 allocs | **24.8x** | 107.4 / 0 allocs | **6.8x** | 19773 / 16937 B / 101 allocs | **1247.5x** |
| Required | 10.58 / 0 allocs | 130.7 / 0 allocs | **12.4x** | 2.119 / 0 allocs | **0.2x** | 14614 / 15488 B / 73 allocs | **1381.3x** |
| IPV4 | 39.88 / 0 allocs | 109.6 / 0 allocs | **2.7x** | N/A | N/A | N/A | N/A |
| Length | 11.29 / 0 allocs | 93.65 / 0 allocs | **8.3x** | 225.4 / 32 B / 2 allocs | **20.0x** | 15125 / 15616 B / 79 allocs | **1339.7x** |
| IPV6 | 106.5 / 0 allocs | 175.6 / 0 allocs | **1.6x** | N/A | N/A | N/A | N/A |
| URL | 63.16 / 0 allocs | 420.9 / 144 B / 1 allocs | **6.7x** | 10987 / 145 B / 1 allocs | **174.0x** | 15772 / 15681 B / 77 allocs | **249.7x** |
| Numeric | 13.03 / 0 allocs | 77.63 / 0 allocs | **6.0x** | 131.6 / 0 allocs | **10.1x** | 15693 / 15574 B / 78 allocs | **1204.4x** |
| GT | 5.645 / 0 allocs | 97.83 / 0 allocs | **17.3x** | 88.87 / 0 allocs | **15.7x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.714 | 0 allocs |
| CELMultipleExpressions | 12.36 | 0 allocs |
| CELBasic | 12.33 | 0 allocs |
| CELCrossField | 8.067 | 0 allocs |
| CELStringLength | 1.000 | 0 allocs |
| CELNumericComparison | 1.063 | 0 allocs |

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
