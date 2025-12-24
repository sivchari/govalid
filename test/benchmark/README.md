# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-24  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124447027	         9.644 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2555088	       473.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11323509	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56457	     21057 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	645194024	         1.842 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296359464	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296412852	         4.051 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	321305780	         3.733 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21670358	        56.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1088 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1319605	       909.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70082	     17090 ns/op	   15852 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256642600	         4.670 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427690879	         2.824 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11207223	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13775716	        88.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	428188420	         2.806 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10757442	       111.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30712414	        39.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9156530	       132.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13397280	        89.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6799776	       176.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147376358	         8.123 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9949262	       120.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4869396	       246.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74780	     16062 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	428679946	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11330743	       105.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	384559315	         3.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11265580	       106.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	203014582	         5.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8187332	       146.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37470865	        31.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8562859	       139.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4227192	       285.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73768	     16304 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	174554048	         6.874 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8263612	       142.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48631489	        24.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9650022	       123.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4252650	       282.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124072713	         9.667 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13646388	        87.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9491690	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   70824	     16652 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296709296	         4.046 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8025663	       149.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	640711972	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76578	     15546 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20738275	        57.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2524962	       478.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102561	     12688 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71499	     16731 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	20153078	        56.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2677831	       449.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3583146	       336.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70389	     16687 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.113 / 0 allocs | 106.7 / 0 allocs | **34.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.670 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 56.07 / 0 allocs | 1088 / 89 B / 5 allocs | **19.4x** | 909.2 / 0 allocs | **16.2x** | 17090 / 15852 B / 76 allocs | **304.8x** |
| GTE | 2.806 / 0 allocs | 111.5 / 0 allocs | **39.7x** | N/A | N/A | N/A | N/A |
| MinLength | 24.64 / 0 allocs | 123.6 / 0 allocs | **5.0x** | 282.0 / 32 B / 2 allocs | **11.4x** | N/A | N/A |
| UUID | 56.09 / 0 allocs | 449.2 / 0 allocs | **8.0x** | 336.1 / 0 allocs | **6.0x** | 16687 / 15542 B / 76 allocs | **297.5x** |
| MaxItems | 5.914 / 0 allocs | 146.4 / 0 allocs | **24.8x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.96 / 0 allocs | 139.4 / 0 allocs | **4.4x** | 285.7 / 32 B / 2 allocs | **8.9x** | 16304 / 15648 B / 81 allocs | **510.1x** |
| LT | 2.802 / 0 allocs | 105.4 / 0 allocs | **37.6x** | N/A | N/A | N/A | N/A |
| MinItems | 6.874 / 0 allocs | 142.2 / 0 allocs | **20.7x** | N/A | N/A | N/A | N/A |
| Alpha | 9.644 / 0 allocs | 473.4 / 0 allocs | **49.1x** | 107.7 / 0 allocs | **11.2x** | 21057 / 16937 B / 101 allocs | **2183.4x** |
| Required | 4.046 / 0 allocs | 149.6 / 0 allocs | **37.0x** | 1.868 / 0 allocs | **0.5x** | 15546 / 15488 B / 73 allocs | **3842.3x** |
| IPV4 | 39.12 / 0 allocs | 132.1 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 8.123 / 0 allocs | 120.3 / 0 allocs | **14.8x** | 246.1 / 32 B / 2 allocs | **30.3x** | 16062 / 15616 B / 79 allocs | **1977.3x** |
| IPV6 | 89.60 / 0 allocs | 176.6 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.73 / 0 allocs | 478.2 / 144 B / 1 allocs | **8.3x** | 12688 / 146 B / 1 allocs | **219.8x** | 16731 / 15681 B / 77 allocs | **289.8x** |
| Numeric | 9.667 / 0 allocs | 87.03 / 0 allocs | **9.0x** | 128.3 / 0 allocs | **13.3x** | 16652 / 15574 B / 78 allocs | **1722.6x** |
| GT | 2.824 / 0 allocs | 106.6 / 0 allocs | **37.7x** | 88.03 / 0 allocs | **31.2x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.842 | 0 allocs |
| CELMultipleExpressions | 4.049 | 0 allocs |
| CELBasic | 4.051 | 0 allocs |
| CELCrossField | 3.733 | 0 allocs |
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
