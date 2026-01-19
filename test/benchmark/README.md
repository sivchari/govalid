# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-01-19  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	100000000	        10.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 3113727	       384.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	12780076	        93.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   61496	     19396 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	489107445	         2.454 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	362390031	         3.291 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	374373256	         3.201 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	416070813	         2.883 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.107 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	652261768	         1.839 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	22016468	        54.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1026 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1542940	       776.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   76711	     15510 ns/op	   15870 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	310705126	         3.876 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	565387572	         2.125 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11777908	       101.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	15736405	        75.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	567293554	         2.116 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11975092	        99.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	38399370	        31.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	10103240	       118.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14827494	        79.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 7099734	       169.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	159544164	         7.682 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	12024817	        99.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 5325066	       223.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   81499	     14708 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	570684260	         2.104 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11906871	       100.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	520757836	         2.310 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11795803	       101.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	277712347	         4.324 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 9370250	       127.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	48968600	        24.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9400606	       127.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4671360	       255.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   78808	     14942 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	231214701	         5.190 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 9501932	       125.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	54066498	        22.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10741826	       111.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4766977	       251.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	134204174	         8.941 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	15314828	        77.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	10719194	       111.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   78573	     15147 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	378052772	         3.176 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8942403	       132.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	1000000000	         1.155 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   83362	     14318 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20595538	        58.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2608173	       460.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  117243	     10655 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   78735	     15315 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	28693975	        41.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2845606	       421.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3744453	       320.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   75867	     15594 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.310 / 0 allocs | 101.6 / 0 allocs | **44.0x** | N/A | N/A | N/A | N/A |
| Enum | 3.876 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 54.78 / 0 allocs | 1026 / 89 B / 5 allocs | **18.7x** | 776.8 / 0 allocs | **14.2x** | 15510 / 15870 B / 76 allocs | **283.1x** |
| GTE | 2.116 / 0 allocs | 99.70 / 0 allocs | **47.1x** | N/A | N/A | N/A | N/A |
| MinLength | 22.21 / 0 allocs | 111.2 / 0 allocs | **5.0x** | 251.6 / 32 B / 2 allocs | **11.3x** | N/A | N/A |
| UUID | 41.51 / 0 allocs | 421.5 / 0 allocs | **10.2x** | 320.0 / 0 allocs | **7.7x** | 15594 / 15542 B / 76 allocs | **375.7x** |
| MaxItems | 4.324 / 0 allocs | 127.7 / 0 allocs | **29.5x** | N/A | N/A | N/A | N/A |
| MaxLength | 24.62 / 0 allocs | 127.3 / 0 allocs | **5.2x** | 255.9 / 32 B / 2 allocs | **10.4x** | 14942 / 15648 B / 81 allocs | **606.9x** |
| LT | 2.104 / 0 allocs | 100.8 / 0 allocs | **47.9x** | N/A | N/A | N/A | N/A |
| MinItems | 5.190 / 0 allocs | 125.8 / 0 allocs | **24.2x** | N/A | N/A | N/A | N/A |
| Alpha | 10.75 / 0 allocs | 384.8 / 0 allocs | **35.8x** | 93.27 / 0 allocs | **8.7x** | 19396 / 16937 B / 101 allocs | **1804.3x** |
| Required | 3.176 / 0 allocs | 132.3 / 0 allocs | **41.7x** | 1.155 / 0 allocs | **0.4x** | 14318 / 15488 B / 73 allocs | **4508.2x** |
| IPV4 | 31.49 / 0 allocs | 118.3 / 0 allocs | **3.8x** | N/A | N/A | N/A | N/A |
| Length | 7.682 / 0 allocs | 99.64 / 0 allocs | **13.0x** | 223.1 / 32 B / 2 allocs | **29.0x** | 14708 / 15616 B / 79 allocs | **1914.6x** |
| IPV6 | 79.93 / 0 allocs | 169.1 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 58.21 / 0 allocs | 460.2 / 144 B / 1 allocs | **7.9x** | 10655 / 145 B / 1 allocs | **183.0x** | 15315 / 15681 B / 77 allocs | **263.1x** |
| Numeric | 8.941 / 0 allocs | 77.84 / 0 allocs | **8.7x** | 111.7 / 0 allocs | **12.5x** | 15147 / 15574 B / 78 allocs | **1694.1x** |
| GT | 2.125 / 0 allocs | 101.4 / 0 allocs | **47.7x** | 75.29 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.454 | 0 allocs |
| CELMultipleExpressions | 3.291 | 0 allocs |
| CELBasic | 3.201 | 0 allocs |
| CELCrossField | 2.883 | 0 allocs |
| CELStringLength | 1.107 | 0 allocs |
| CELNumericComparison | 1.839 | 0 allocs |

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
