# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-23  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124289545	         9.654 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2561686	       464.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11177696	       108.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56287	     21213 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	652905866	         1.838 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296254633	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295386962	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	321502228	         3.734 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21711842	        56.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1089 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1351089	       887.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68732	     17182 ns/op	   15857 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256961666	         4.670 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	428714733	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11254742	       106.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13780908	        88.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	427949570	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10707882	       112.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30597500	        39.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9165651	       131.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14082566	        85.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6933754	       173.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	148320897	         8.094 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9987439	       120.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4349982	       252.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73342	     16265 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	427972210	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11374296	       105.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385584241	         3.112 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11268943	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202927125	         5.910 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8182923	       145.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37573360	        31.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8631192	       138.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4215328	       283.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74484	     16754 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	174965108	         6.859 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8398362	       142.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48798207	        24.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10095775	       119.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4260246	       281.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124203886	         9.659 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13771815	        87.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9282594	       128.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71482	     16764 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296158898	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7989716	       149.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642863784	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75390	     15769 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20828630	        57.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2521330	       475.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103071	     11694 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71044	     16821 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24126283	        55.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2692015	       447.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3557564	       336.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71629	     16778 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.112 / 0 allocs | 107.7 / 0 allocs | **34.6x** | N/A | N/A | N/A | N/A |
| Enum | 4.670 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 56.37 / 0 allocs | 1089 / 89 B / 5 allocs | **19.3x** | 887.1 / 0 allocs | **15.7x** | 17182 / 15857 B / 76 allocs | **304.8x** |
| GTE | 2.801 / 0 allocs | 112.0 / 0 allocs | **40.0x** | N/A | N/A | N/A | N/A |
| MinLength | 24.60 / 0 allocs | 119.0 / 0 allocs | **4.8x** | 281.9 / 32 B / 2 allocs | **11.5x** | N/A | N/A |
| UUID | 55.73 / 0 allocs | 447.4 / 0 allocs | **8.0x** | 336.8 / 0 allocs | **6.0x** | 16778 / 15542 B / 76 allocs | **301.1x** |
| MaxItems | 5.910 / 0 allocs | 145.8 / 0 allocs | **24.7x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.73 / 0 allocs | 138.7 / 0 allocs | **4.4x** | 283.6 / 32 B / 2 allocs | **8.9x** | 16754 / 15648 B / 81 allocs | **528.0x** |
| LT | 2.802 / 0 allocs | 105.5 / 0 allocs | **37.7x** | N/A | N/A | N/A | N/A |
| MinItems | 6.859 / 0 allocs | 142.7 / 0 allocs | **20.8x** | N/A | N/A | N/A | N/A |
| Alpha | 9.654 / 0 allocs | 464.6 / 0 allocs | **48.1x** | 108.1 / 0 allocs | **11.2x** | 21213 / 16937 B / 101 allocs | **2197.3x** |
| Required | 4.049 / 0 allocs | 149.8 / 0 allocs | **37.0x** | 1.867 / 0 allocs | **0.5x** | 15769 / 15488 B / 73 allocs | **3894.5x** |
| IPV4 | 39.93 / 0 allocs | 131.2 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.094 / 0 allocs | 120.4 / 0 allocs | **14.9x** | 252.7 / 32 B / 2 allocs | **31.2x** | 16265 / 15616 B / 79 allocs | **2009.5x** |
| IPV6 | 85.25 / 0 allocs | 173.7 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.65 / 0 allocs | 475.8 / 144 B / 1 allocs | **8.3x** | 11694 / 145 B / 1 allocs | **202.8x** | 16821 / 15681 B / 77 allocs | **291.8x** |
| Numeric | 9.659 / 0 allocs | 87.50 / 0 allocs | **9.1x** | 128.4 / 0 allocs | **13.3x** | 16764 / 15574 B / 78 allocs | **1735.6x** |
| GT | 2.801 / 0 allocs | 106.7 / 0 allocs | **38.1x** | 88.20 / 0 allocs | **31.5x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.838 | 0 allocs |
| CELMultipleExpressions | 4.049 | 0 allocs |
| CELBasic | 4.054 | 0 allocs |
| CELCrossField | 3.734 | 0 allocs |
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
