# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-12  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	81518332	        14.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2640489	       465.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10884592	       113.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60403	     20657 ns/op	   16936 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	271443921	         4.410 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	147977426	         8.110 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19441687	        61.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1077 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1335008	       898.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   74629	     15954 ns/op	   15820 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	92931922	        12.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	213901120	         5.609 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10878016	       112.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13554586	        88.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	212922164	         5.653 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10883875	       112.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27531571	        44.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8828104	       135.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13874488	        88.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6485300	       185.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	124368450	         9.654 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10797582	       112.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4756790	       249.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80464	     14897 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	214225734	         5.602 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10975116	       111.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202708531	         5.923 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10380974	       118.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	99006526	        12.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8320531	       145.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	32948656	        36.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9597271	       126.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4278187	       280.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74745	     15430 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	100000000	        11.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8713447	       138.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48042265	        24.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9234140	       130.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4243989	       280.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	95281803	        12.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13765134	        87.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9376909	       128.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   76503	     15461 ns/op	   15573 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	128332468	         9.348 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8248807	       145.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641859711	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   81619	     14519 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19701556	        60.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2493891	       480.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102165	     12067 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   75661	     15796 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	18513874	        64.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2603150	       464.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3427191	       351.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   76466	     15663 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.923 / 0 allocs | 118.1 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| Enum | 12.79 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 61.73 / 0 allocs | 1077 / 89 B / 5 allocs | **17.4x** | 898.7 / 0 allocs | **14.6x** | 15954 / 15820 B / 76 allocs | **258.4x** |
| GTE | 5.653 / 0 allocs | 112.8 / 0 allocs | **20.0x** | N/A | N/A | N/A | N/A |
| MinLength | 24.91 / 0 allocs | 130.7 / 0 allocs | **5.2x** | 280.3 / 32 B / 2 allocs | **11.3x** | N/A | N/A |
| UUID | 64.58 / 0 allocs | 464.5 / 0 allocs | **7.2x** | 351.2 / 0 allocs | **5.4x** | 15663 / 15541 B / 76 allocs | **242.5x** |
| MaxItems | 12.15 / 0 allocs | 145.1 / 0 allocs | **11.9x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.49 / 0 allocs | 126.7 / 0 allocs | **3.5x** | 280.0 / 32 B / 2 allocs | **7.7x** | 15430 / 15648 B / 81 allocs | **422.9x** |
| LT | 5.602 / 0 allocs | 111.9 / 0 allocs | **20.0x** | N/A | N/A | N/A | N/A |
| MinItems | 11.84 / 0 allocs | 138.9 / 0 allocs | **11.7x** | N/A | N/A | N/A | N/A |
| Alpha | 14.64 / 0 allocs | 465.8 / 0 allocs | **31.8x** | 113.1 / 0 allocs | **7.7x** | 20657 / 16936 B / 101 allocs | **1411.0x** |
| Required | 9.348 / 0 allocs | 145.6 / 0 allocs | **15.6x** | 1.868 / 0 allocs | **0.2x** | 14519 / 15488 B / 73 allocs | **1553.2x** |
| IPV4 | 44.14 / 0 allocs | 135.9 / 0 allocs | **3.1x** | N/A | N/A | N/A | N/A |
| Length | 9.654 / 0 allocs | 112.1 / 0 allocs | **11.6x** | 249.4 / 32 B / 2 allocs | **25.8x** | 14897 / 15616 B / 79 allocs | **1543.1x** |
| IPV6 | 88.05 / 0 allocs | 185.5 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 60.96 / 0 allocs | 480.2 / 144 B / 1 allocs | **7.9x** | 12067 / 146 B / 1 allocs | **197.9x** | 15796 / 15681 B / 77 allocs | **259.1x** |
| Numeric | 12.56 / 0 allocs | 87.63 / 0 allocs | **7.0x** | 128.0 / 0 allocs | **10.2x** | 15461 / 15573 B / 78 allocs | **1231.0x** |
| GT | 5.609 / 0 allocs | 112.2 / 0 allocs | **20.0x** | 88.61 / 0 allocs | **15.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.410 | 0 allocs |
| CELMultipleExpressions | 11.54 | 0 allocs |
| CELBasic | 11.54 | 0 allocs |
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
