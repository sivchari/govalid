# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-13  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	123827024	         9.681 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2425776	       498.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11228383	       108.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56155	     21062 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644008099	         1.865 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296048942	         4.057 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296324568	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	348214801	         3.437 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21826947	        58.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1084 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1355352	       884.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69728	     17131 ns/op	   15856 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275651758	         4.354 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	482210384	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11159126	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13681626	        88.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482398137	         2.488 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11089777	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30759777	        39.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9273693	       132.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14086770	        85.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6869503	       175.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160257594	         7.481 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10081105	       119.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4851303	       247.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74130	     16055 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481401577	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11367673	       105.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427113376	         2.805 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11166806	       107.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183755925	         6.535 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8297233	       144.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35973205	        33.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8655844	       138.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4230553	       284.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73623	     16377 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202718607	         5.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8501988	       141.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52057110	        23.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10086650	       118.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4198611	       283.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157789860	         7.779 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13751349	        87.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9455714	       128.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71839	     16618 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	318345338	         3.745 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7910457	       151.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643092592	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76503	     15567 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20587238	        58.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2526262	       479.4 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103038	     11723 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71733	     16660 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24613377	        48.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2672464	       449.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3581451	       336.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   72340	     16730 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.805 / 0 allocs | 107.8 / 0 allocs | **38.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.354 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.71 / 0 allocs | 1084 / 88 B / 5 allocs | **18.5x** | 884.8 / 0 allocs | **15.1x** | 17131 / 15856 B / 76 allocs | **291.8x** |
| GTE | 2.488 / 0 allocs | 108.3 / 0 allocs | **43.5x** | N/A | N/A | N/A | N/A |
| MinLength | 23.06 / 0 allocs | 118.9 / 0 allocs | **5.2x** | 283.4 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.57 / 0 allocs | 449.2 / 0 allocs | **9.2x** | 336.1 / 0 allocs | **6.9x** | 16730 / 15542 B / 76 allocs | **344.5x** |
| MaxItems | 6.535 / 0 allocs | 144.4 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.31 / 0 allocs | 138.7 / 0 allocs | **4.2x** | 284.6 / 32 B / 2 allocs | **8.5x** | 16377 / 15648 B / 81 allocs | **491.7x** |
| LT | 2.492 / 0 allocs | 105.2 / 0 allocs | **42.2x** | N/A | N/A | N/A | N/A |
| MinItems | 5.920 / 0 allocs | 141.6 / 0 allocs | **23.9x** | N/A | N/A | N/A | N/A |
| Alpha | 9.681 / 0 allocs | 498.4 / 0 allocs | **51.5x** | 108.1 / 0 allocs | **11.2x** | 21062 / 16937 B / 101 allocs | **2175.6x** |
| Required | 3.745 / 0 allocs | 151.4 / 0 allocs | **40.4x** | 1.868 / 0 allocs | **0.5x** | 15567 / 15488 B / 73 allocs | **4156.7x** |
| IPV4 | 39.21 / 0 allocs | 132.4 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 7.481 / 0 allocs | 119.2 / 0 allocs | **15.9x** | 247.1 / 32 B / 2 allocs | **33.0x** | 16055 / 15616 B / 79 allocs | **2146.1x** |
| IPV6 | 85.34 / 0 allocs | 175.6 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 58.18 / 0 allocs | 479.4 / 144 B / 1 allocs | **8.2x** | 11723 / 145 B / 1 allocs | **201.5x** | 16660 / 15681 B / 77 allocs | **286.4x** |
| Numeric | 7.779 / 0 allocs | 87.49 / 0 allocs | **11.2x** | 128.7 / 0 allocs | **16.5x** | 16618 / 15574 B / 78 allocs | **2136.3x** |
| GT | 2.490 / 0 allocs | 106.3 / 0 allocs | **42.7x** | 88.73 / 0 allocs | **35.6x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.865 | 0 allocs |
| CELMultipleExpressions | 4.057 | 0 allocs |
| CELBasic | 4.049 | 0 allocs |
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
