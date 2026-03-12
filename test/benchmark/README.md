# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-12  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	77951354	        14.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2647995	       471.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	 9543426	       117.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   57098	     19996 ns/op	   16936 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	272286534	         4.403 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        11.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	147937071	         8.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19479133	        61.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1083 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1333614	       900.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   75039	     15986 ns/op	   15798 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	92242292	        12.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	213701677	         5.611 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10817124	       114.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13592744	        88.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	214094124	         5.605 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10878030	       112.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	26890653	        43.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8848555	       135.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13890103	        88.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6432450	       184.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	124219825	         9.657 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10696908	       112.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4771186	       250.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80056	     14967 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	213804122	         5.612 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10912923	       111.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202353085	         5.935 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10346586	       118.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	98084955	        12.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8355026	       144.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	32904244	        36.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9610023	       126.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4229745	       280.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   80671	     15043 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	100000000	        11.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8646230	       138.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	47262736	        24.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9282996	       130.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4264038	       280.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	93373405	        12.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13704338	        88.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9363198	       128.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   77479	     15444 ns/op	   15573 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	127964106	         9.369 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8191747	       146.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641476428	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   84303	     14497 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19658492	        60.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2504377	       479.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   98707	     12175 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   77768	     15589 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	18274794	        65.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2604944	       463.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3403488	       351.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   76560	     15509 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.935 / 0 allocs | 118.3 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| Enum | 12.79 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 61.72 / 0 allocs | 1083 / 88 B / 5 allocs | **17.5x** | 900.1 / 0 allocs | **14.6x** | 15986 / 15798 B / 76 allocs | **259.0x** |
| GTE | 5.605 / 0 allocs | 112.9 / 0 allocs | **20.1x** | N/A | N/A | N/A | N/A |
| MinLength | 24.92 / 0 allocs | 130.1 / 0 allocs | **5.2x** | 280.7 / 32 B / 2 allocs | **11.3x** | N/A | N/A |
| UUID | 65.14 / 0 allocs | 463.1 / 0 allocs | **7.1x** | 351.9 / 0 allocs | **5.4x** | 15509 / 15541 B / 76 allocs | **238.1x** |
| MaxItems | 12.14 / 0 allocs | 144.9 / 0 allocs | **11.9x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.47 / 0 allocs | 126.2 / 0 allocs | **3.5x** | 280.5 / 32 B / 2 allocs | **7.7x** | 15043 / 15648 B / 81 allocs | **412.5x** |
| LT | 5.612 / 0 allocs | 111.8 / 0 allocs | **19.9x** | N/A | N/A | N/A | N/A |
| MinItems | 11.84 / 0 allocs | 138.4 / 0 allocs | **11.7x** | N/A | N/A | N/A | N/A |
| Alpha | 14.71 / 0 allocs | 471.1 / 0 allocs | **32.0x** | 117.9 / 0 allocs | **8.0x** | 19996 / 16936 B / 101 allocs | **1359.3x** |
| Required | 9.369 / 0 allocs | 146.1 / 0 allocs | **15.6x** | 1.869 / 0 allocs | **0.2x** | 14497 / 15488 B / 73 allocs | **1547.3x** |
| IPV4 | 43.93 / 0 allocs | 135.0 / 0 allocs | **3.1x** | N/A | N/A | N/A | N/A |
| Length | 9.657 / 0 allocs | 112.2 / 0 allocs | **11.6x** | 250.2 / 32 B / 2 allocs | **25.9x** | 14967 / 15616 B / 79 allocs | **1549.9x** |
| IPV6 | 88.49 / 0 allocs | 184.9 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 60.96 / 0 allocs | 479.8 / 144 B / 1 allocs | **7.9x** | 12175 / 146 B / 1 allocs | **199.7x** | 15589 / 15681 B / 77 allocs | **255.7x** |
| Numeric | 12.50 / 0 allocs | 88.18 / 0 allocs | **7.1x** | 128.4 / 0 allocs | **10.3x** | 15444 / 15573 B / 78 allocs | **1235.5x** |
| GT | 5.611 / 0 allocs | 114.0 / 0 allocs | **20.3x** | 88.66 / 0 allocs | **15.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.403 | 0 allocs |
| CELMultipleExpressions | 11.54 | 0 allocs |
| CELBasic | 11.58 | 0 allocs |
| CELCrossField | 8.113 | 0 allocs |
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
