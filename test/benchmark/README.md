# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-23  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124362986	         9.647 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2582198	       476.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11226124	       108.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56610	     21219 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	653028502	         1.838 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295534905	         4.319 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295629612	         4.051 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320700574	         3.738 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21031668	        59.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1092 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1351482	       886.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68962	     17133 ns/op	   15868 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256973607	         4.672 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427560982	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11235043	       106.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13640294	        88.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	428944531	         2.799 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11017986	       109.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30606342	        38.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9206137	       131.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14042714	        85.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6933048	       174.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147842655	         8.108 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9972555	       120.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4846492	       246.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   72598	     16327 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	428467978	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11350929	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385092829	         3.115 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11211068	       107.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	203003946	         5.916 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8275600	       145.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37094714	        31.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8623315	       139.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4219987	       284.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73143	     16596 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	174853021	         6.860 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8538435	       140.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48753962	        24.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10054608	       119.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4211100	       282.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124329319	         9.651 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13603773	        87.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9325149	       128.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   70356	     16840 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296335333	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8016472	       149.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642432932	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75093	     15786 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20657876	        58.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2487832	       482.3 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  101457	     11766 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70854	     16890 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	20952420	        54.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2689272	       447.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3459144	       344.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71276	     17139 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.115 / 0 allocs | 107.0 / 0 allocs | **34.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.672 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.30 / 0 allocs | 1092 / 89 B / 5 allocs | **18.4x** | 886.9 / 0 allocs | **15.0x** | 17133 / 15868 B / 76 allocs | **288.9x** |
| GTE | 2.799 / 0 allocs | 109.6 / 0 allocs | **39.2x** | N/A | N/A | N/A | N/A |
| MinLength | 24.60 / 0 allocs | 119.3 / 0 allocs | **4.8x** | 282.6 / 32 B / 2 allocs | **11.5x** | N/A | N/A |
| UUID | 54.97 / 0 allocs | 447.0 / 0 allocs | **8.1x** | 344.0 / 0 allocs | **6.3x** | 17139 / 15542 B / 76 allocs | **311.8x** |
| MaxItems | 5.916 / 0 allocs | 145.3 / 0 allocs | **24.6x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.80 / 0 allocs | 139.1 / 0 allocs | **4.4x** | 284.4 / 32 B / 2 allocs | **8.9x** | 16596 / 15648 B / 81 allocs | **521.9x** |
| LT | 2.801 / 0 allocs | 105.6 / 0 allocs | **37.7x** | N/A | N/A | N/A | N/A |
| MinItems | 6.860 / 0 allocs | 140.9 / 0 allocs | **20.5x** | N/A | N/A | N/A | N/A |
| Alpha | 9.647 / 0 allocs | 476.1 / 0 allocs | **49.4x** | 108.2 / 0 allocs | **11.2x** | 21219 / 16937 B / 101 allocs | **2199.5x** |
| Required | 4.048 / 0 allocs | 149.6 / 0 allocs | **37.0x** | 1.869 / 0 allocs | **0.5x** | 15786 / 15488 B / 73 allocs | **3899.7x** |
| IPV4 | 38.98 / 0 allocs | 131.6 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 8.108 / 0 allocs | 120.1 / 0 allocs | **14.8x** | 246.4 / 32 B / 2 allocs | **30.4x** | 16327 / 15616 B / 79 allocs | **2013.7x** |
| IPV6 | 85.19 / 0 allocs | 174.6 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 58.14 / 0 allocs | 482.3 / 144 B / 1 allocs | **8.3x** | 11766 / 145 B / 1 allocs | **202.4x** | 16890 / 15681 B / 77 allocs | **290.5x** |
| Numeric | 9.651 / 0 allocs | 87.82 / 0 allocs | **9.1x** | 128.6 / 0 allocs | **13.3x** | 16840 / 15574 B / 78 allocs | **1744.9x** |
| GT | 2.802 / 0 allocs | 106.9 / 0 allocs | **38.2x** | 88.95 / 0 allocs | **31.7x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.838 | 0 allocs |
| CELMultipleExpressions | 4.319 | 0 allocs |
| CELBasic | 4.051 | 0 allocs |
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
