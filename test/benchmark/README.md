# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-17  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124243648	         9.661 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2565812	       471.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11175120	       112.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   48396	     22471 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	645373069	         1.864 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	294752461	         4.060 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295575196	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349922754	         3.437 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21390756	        55.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1102 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1347549	       889.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   67838	     17523 ns/op	   15864 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	274781262	         4.369 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481008354	         2.495 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11280273	       106.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13795730	        88.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	479455887	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11003590	       108.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30820288	        39.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9020047	       132.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14014843	        85.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6926838	       178.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160573009	         7.475 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9974319	       120.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4706950	       250.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73039	     16385 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481437564	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11022856	       109.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	424878859	         2.807 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11226273	       107.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183048954	         6.548 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8344935	       144.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36056906	        33.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8643289	       138.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4129873	       289.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72834	     16584 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202597435	         5.928 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8476945	       140.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52021852	        23.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10037979	       119.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4185608	       285.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157637295	         7.616 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12956428	        91.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9420602	       128.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   69243	     17261 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321413188	         3.734 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8004686	       150.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641503674	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   73022	     16191 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20807598	        57.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2454738	       486.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102715	     11724 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   68377	     17151 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23601666	        52.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2680249	       453.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3586891	       335.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   67782	     17452 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.807 / 0 allocs | 107.1 / 0 allocs | **38.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.369 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.40 / 0 allocs | 1102 / 89 B / 5 allocs | **19.9x** | 889.1 / 0 allocs | **16.0x** | 17523 / 15864 B / 76 allocs | **316.3x** |
| GTE | 2.493 / 0 allocs | 108.8 / 0 allocs | **43.6x** | N/A | N/A | N/A | N/A |
| MinLength | 23.07 / 0 allocs | 119.3 / 0 allocs | **5.2x** | 285.9 / 32 B / 2 allocs | **12.4x** | N/A | N/A |
| UUID | 52.44 / 0 allocs | 453.9 / 0 allocs | **8.7x** | 335.8 / 0 allocs | **6.4x** | 17452 / 15542 B / 76 allocs | **332.8x** |
| MaxItems | 6.548 / 0 allocs | 144.0 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.30 / 0 allocs | 138.8 / 0 allocs | **4.2x** | 289.4 / 32 B / 2 allocs | **8.7x** | 16584 / 15648 B / 81 allocs | **498.0x** |
| LT | 2.493 / 0 allocs | 109.2 / 0 allocs | **43.8x** | N/A | N/A | N/A | N/A |
| MinItems | 5.928 / 0 allocs | 140.6 / 0 allocs | **23.7x** | N/A | N/A | N/A | N/A |
| Alpha | 9.661 / 0 allocs | 471.8 / 0 allocs | **48.8x** | 112.1 / 0 allocs | **11.6x** | 22471 / 16937 B / 101 allocs | **2325.9x** |
| Required | 3.734 / 0 allocs | 150.7 / 0 allocs | **40.4x** | 1.869 / 0 allocs | **0.5x** | 16191 / 15488 B / 73 allocs | **4336.1x** |
| IPV4 | 39.10 / 0 allocs | 132.6 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 7.475 / 0 allocs | 120.5 / 0 allocs | **16.1x** | 250.4 / 32 B / 2 allocs | **33.5x** | 16385 / 15616 B / 79 allocs | **2192.0x** |
| IPV6 | 85.92 / 0 allocs | 178.6 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 57.66 / 0 allocs | 486.8 / 144 B / 1 allocs | **8.4x** | 11724 / 145 B / 1 allocs | **203.3x** | 17151 / 15681 B / 77 allocs | **297.5x** |
| Numeric | 7.616 / 0 allocs | 91.43 / 0 allocs | **12.0x** | 128.4 / 0 allocs | **16.9x** | 17261 / 15574 B / 78 allocs | **2266.4x** |
| GT | 2.495 / 0 allocs | 106.7 / 0 allocs | **42.8x** | 88.07 / 0 allocs | **35.3x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.864 | 0 allocs |
| CELMultipleExpressions | 4.060 | 0 allocs |
| CELBasic | 4.054 | 0 allocs |
| CELCrossField | 3.437 | 0 allocs |
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
