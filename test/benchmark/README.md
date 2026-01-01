# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-01-01  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124301976	         9.653 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2292151	       511.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11063991	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56379	     21084 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	652882634	         1.897 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	282504823	         4.120 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296321090	         4.047 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320662036	         3.741 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21683862	        59.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1089 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1312645	       916.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69230	     17336 ns/op	   15840 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	251816469	         4.799 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427887999	         2.805 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11322126	       106.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13682394	        88.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	427717826	         2.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11009286	       109.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30816211	        39.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9252567	       130.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13923348	        85.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6582582	       174.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147720204	         8.122 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9948267	       121.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4758618	       250.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   72456	     16297 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	427428177	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11375050	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385283264	         3.139 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11218114	       107.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	203056057	         5.973 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8304711	       144.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37806648	        31.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8654614	       139.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4208985	       282.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73153	     16387 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	174987714	         6.857 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8526855	       143.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48267865	        24.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10135010	       118.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4221223	       282.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124404854	         9.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13738434	        87.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9355123	       128.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71510	     17425 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296547414	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8013408	       149.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643051137	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76476	     15712 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20643855	        57.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2502816	       478.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103232	     11723 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71349	     16764 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	26302767	        51.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2646027	       449.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3591724	       337.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70995	     16814 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.139 / 0 allocs | 107.0 / 0 allocs | **34.1x** | N/A | N/A | N/A | N/A |
| Enum | 4.799 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.15 / 0 allocs | 1089 / 88 B / 5 allocs | **18.4x** | 916.3 / 0 allocs | **15.5x** | 17336 / 15840 B / 76 allocs | **293.1x** |
| GTE | 2.804 / 0 allocs | 109.7 / 0 allocs | **39.1x** | N/A | N/A | N/A | N/A |
| MinLength | 24.61 / 0 allocs | 118.8 / 0 allocs | **4.8x** | 282.2 / 32 B / 2 allocs | **11.5x** | N/A | N/A |
| UUID | 51.24 / 0 allocs | 449.9 / 0 allocs | **8.8x** | 337.4 / 0 allocs | **6.6x** | 16814 / 15542 B / 76 allocs | **328.1x** |
| MaxItems | 5.973 / 0 allocs | 144.7 / 0 allocs | **24.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.73 / 0 allocs | 139.0 / 0 allocs | **4.4x** | 282.9 / 32 B / 2 allocs | **8.9x** | 16387 / 15648 B / 81 allocs | **516.5x** |
| LT | 2.802 / 0 allocs | 105.6 / 0 allocs | **37.7x** | N/A | N/A | N/A | N/A |
| MinItems | 6.857 / 0 allocs | 143.1 / 0 allocs | **20.9x** | N/A | N/A | N/A | N/A |
| Alpha | 9.653 / 0 allocs | 511.5 / 0 allocs | **53.0x** | 108.3 / 0 allocs | **11.2x** | 21084 / 16937 B / 101 allocs | **2184.2x** |
| Required | 4.049 / 0 allocs | 149.6 / 0 allocs | **36.9x** | 1.866 / 0 allocs | **0.5x** | 15712 / 15488 B / 73 allocs | **3880.5x** |
| IPV4 | 39.11 / 0 allocs | 130.0 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.122 / 0 allocs | 121.1 / 0 allocs | **14.9x** | 250.0 / 32 B / 2 allocs | **30.8x** | 16297 / 15616 B / 79 allocs | **2006.5x** |
| IPV6 | 85.42 / 0 allocs | 174.9 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.77 / 0 allocs | 478.1 / 144 B / 1 allocs | **8.3x** | 11723 / 145 B / 1 allocs | **202.9x** | 16764 / 15681 B / 77 allocs | **290.2x** |
| Numeric | 9.645 / 0 allocs | 87.13 / 0 allocs | **9.0x** | 128.4 / 0 allocs | **13.3x** | 17425 / 15574 B / 78 allocs | **1806.6x** |
| GT | 2.805 / 0 allocs | 106.5 / 0 allocs | **38.0x** | 88.05 / 0 allocs | **31.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.897 | 0 allocs |
| CELMultipleExpressions | 4.120 | 0 allocs |
| CELBasic | 4.047 | 0 allocs |
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
