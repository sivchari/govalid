# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-25  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	78977935	        14.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2650788	       448.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10484689	       116.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60399	     19664 ns/op	   16936 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	272081869	         4.416 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	99370322	        11.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	148080636	         8.103 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19361654	        61.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1083 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1333552	       900.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   74236	     16040 ns/op	   15806 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	93036496	        12.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	214099408	         5.604 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10818502	       112.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13525929	        90.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	214019821	         5.610 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10603944	       113.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27165466	        44.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8814380	       137.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13845765	        86.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 5242719	       193.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	124217716	         9.661 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10736835	       112.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4781638	       249.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80592	     14855 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	214142697	         5.617 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10805581	       112.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202839217	         5.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10330128	       118.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	98776740	        12.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8260544	       145.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	32892267	        36.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9503980	       126.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4262342	       279.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   80893	     15044 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	95935936	        11.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8665980	       139.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48110310	        24.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9252135	       130.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4237183	       280.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	94746548	        12.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13752228	        88.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9470562	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   77325	     15462 ns/op	   15573 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	128405968	         9.346 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8249384	       145.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642263964	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   82317	     14890 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19408346	        61.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2518646	       474.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102012	     12120 ns/op	     148 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   78024	     15568 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	18470944	        64.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2556279	       474.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3424929	       353.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   75411	     15609 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.915 / 0 allocs | 118.0 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| Enum | 12.78 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 61.75 / 0 allocs | 1083 / 88 B / 5 allocs | **17.5x** | 900.0 / 0 allocs | **14.6x** | 16040 / 15806 B / 76 allocs | **259.8x** |
| GTE | 5.610 / 0 allocs | 113.6 / 0 allocs | **20.2x** | N/A | N/A | N/A | N/A |
| MinLength | 24.90 / 0 allocs | 130.4 / 0 allocs | **5.2x** | 280.5 / 32 B / 2 allocs | **11.3x** | N/A | N/A |
| UUID | 64.81 / 0 allocs | 474.3 / 0 allocs | **7.3x** | 353.6 / 0 allocs | **5.5x** | 15609 / 15541 B / 76 allocs | **240.8x** |
| MaxItems | 12.16 / 0 allocs | 145.4 / 0 allocs | **12.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.43 / 0 allocs | 126.5 / 0 allocs | **3.5x** | 279.7 / 32 B / 2 allocs | **7.7x** | 15044 / 15648 B / 81 allocs | **413.0x** |
| LT | 5.617 / 0 allocs | 112.0 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| MinItems | 11.87 / 0 allocs | 139.7 / 0 allocs | **11.8x** | N/A | N/A | N/A | N/A |
| Alpha | 14.75 / 0 allocs | 448.6 / 0 allocs | **30.4x** | 116.5 / 0 allocs | **7.9x** | 19664 / 16936 B / 101 allocs | **1333.2x** |
| Required | 9.346 / 0 allocs | 145.2 / 0 allocs | **15.5x** | 1.868 / 0 allocs | **0.2x** | 14890 / 15488 B / 73 allocs | **1593.2x** |
| IPV4 | 44.24 / 0 allocs | 137.0 / 0 allocs | **3.1x** | N/A | N/A | N/A | N/A |
| Length | 9.661 / 0 allocs | 112.0 / 0 allocs | **11.6x** | 249.8 / 32 B / 2 allocs | **25.9x** | 14855 / 15616 B / 79 allocs | **1537.6x** |
| IPV6 | 86.99 / 0 allocs | 193.2 / 0 allocs | **2.2x** | N/A | N/A | N/A | N/A |
| URL | 61.73 / 0 allocs | 474.9 / 144 B / 1 allocs | **7.7x** | 12120 / 148 B / 1 allocs | **196.3x** | 15568 / 15681 B / 77 allocs | **252.2x** |
| Numeric | 12.51 / 0 allocs | 88.14 / 0 allocs | **7.0x** | 128.3 / 0 allocs | **10.3x** | 15462 / 15573 B / 78 allocs | **1236.0x** |
| GT | 5.604 / 0 allocs | 112.7 / 0 allocs | **20.1x** | 90.17 / 0 allocs | **16.1x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.416 | 0 allocs |
| CELMultipleExpressions | 11.54 | 0 allocs |
| CELBasic | 11.55 | 0 allocs |
| CELCrossField | 8.103 | 0 allocs |
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
