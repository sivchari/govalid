# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-24  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124472966	         9.641 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2559632	       472.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11333546	       108.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55210	     21467 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	640133498	         2.043 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	288252607	         4.088 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296212767	         4.050 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349944921	         3.430 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21572521	        56.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1093 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1319833	       909.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69512	     17344 ns/op	   15863 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	273687084	         4.374 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481044159	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11231926	       106.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13775710	        88.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482208615	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10941224	       109.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30789465	        39.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9179566	       131.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14096822	        86.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6900664	       174.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160643826	         7.471 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9904730	       121.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4872518	       248.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73525	     16424 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481925722	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10941955	       109.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427842357	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11212292	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183606577	         6.535 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8321991	       144.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35967678	        33.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8610346	       140.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4174270	       286.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72871	     16444 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202747909	         5.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8436386	       142.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51826674	        23.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10120142	       118.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4223370	       283.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157165618	         7.610 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12993990	        91.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9376443	       128.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71552	     16871 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321412858	         3.737 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7978142	       150.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641681289	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75430	     15873 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20689070	        57.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2503741	       477.7 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102849	     11710 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70942	     16912 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23526832	        51.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2670788	       458.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3568875	       336.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71072	     16976 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.801 / 0 allocs | 107.7 / 0 allocs | **38.5x** | N/A | N/A | N/A | N/A |
| Enum | 4.374 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 56.81 / 0 allocs | 1093 / 89 B / 5 allocs | **19.2x** | 909.1 / 0 allocs | **16.0x** | 17344 / 15863 B / 76 allocs | **305.3x** |
| GTE | 2.491 / 0 allocs | 109.5 / 0 allocs | **44.0x** | N/A | N/A | N/A | N/A |
| MinLength | 23.05 / 0 allocs | 118.7 / 0 allocs | **5.1x** | 283.1 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 51.83 / 0 allocs | 458.6 / 0 allocs | **8.8x** | 336.3 / 0 allocs | **6.5x** | 16976 / 15542 B / 76 allocs | **327.5x** |
| MaxItems | 6.535 / 0 allocs | 144.4 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.34 / 0 allocs | 140.5 / 0 allocs | **4.2x** | 286.9 / 32 B / 2 allocs | **8.6x** | 16444 / 15648 B / 81 allocs | **493.2x** |
| LT | 2.489 / 0 allocs | 109.4 / 0 allocs | **44.0x** | N/A | N/A | N/A | N/A |
| MinItems | 5.918 / 0 allocs | 142.6 / 0 allocs | **24.1x** | N/A | N/A | N/A | N/A |
| Alpha | 9.641 / 0 allocs | 472.8 / 0 allocs | **49.0x** | 108.2 / 0 allocs | **11.2x** | 21467 / 16937 B / 101 allocs | **2226.6x** |
| Required | 3.737 / 0 allocs | 150.7 / 0 allocs | **40.3x** | 1.869 / 0 allocs | **0.5x** | 15873 / 15488 B / 73 allocs | **4247.5x** |
| IPV4 | 39.07 / 0 allocs | 131.9 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 7.471 / 0 allocs | 121.2 / 0 allocs | **16.2x** | 248.3 / 32 B / 2 allocs | **33.2x** | 16424 / 15616 B / 79 allocs | **2198.4x** |
| IPV6 | 86.09 / 0 allocs | 174.1 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.66 / 0 allocs | 477.7 / 144 B / 1 allocs | **8.3x** | 11710 / 145 B / 1 allocs | **203.1x** | 16912 / 15681 B / 77 allocs | **293.3x** |
| Numeric | 7.610 / 0 allocs | 91.92 / 0 allocs | **12.1x** | 128.6 / 0 allocs | **16.9x** | 16871 / 15574 B / 78 allocs | **2217.0x** |
| GT | 2.489 / 0 allocs | 106.4 / 0 allocs | **42.7x** | 88.12 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.043 | 0 allocs |
| CELMultipleExpressions | 4.088 | 0 allocs |
| CELBasic | 4.050 | 0 allocs |
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
