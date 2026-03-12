# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-12  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	80133351	        14.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 3090622	       386.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	12716778	        98.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   59624	     19800 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	217805168	         5.505 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	128480263	         9.355 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	132047014	         9.072 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	187486515	         6.275 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	915147562	         1.326 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	653911508	         1.831 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	20475369	        58.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1029 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1569976	       765.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   75716	     15672 ns/op	   15885 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	124226901	         9.658 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	292743204	         4.097 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11844534	       101.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	16252045	        74.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	293090554	         4.095 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11964016	        99.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	35599724	        33.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	10262659	       116.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14845032	        80.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6948259	       172.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	125080584	         9.596 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	12043932	        99.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 5348545	       223.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   81242	     14932 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	291738721	         4.127 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11855110	       101.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	292237712	         4.106 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11774860	       101.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	132136269	         9.093 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 9410542	       127.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	46477453	        25.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9454484	       126.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4721584	       254.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   80554	     14942 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	100000000	        11.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 9400356	       127.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	49864584	        26.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10818753	       110.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4750090	       252.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	100000000	        10.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	15301425	        78.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	10956886	       109.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   77072	     15332 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	152302075	         7.891 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 9055737	       132.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	831953641	         1.462 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   83174	     14377 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20776083	        57.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2580849	       470.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  116966	     10266 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   78163	     15407 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	27632360	        43.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2830240	       422.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3724658	       322.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   76468	     15453 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 4.106 / 0 allocs | 101.9 / 0 allocs | **24.8x** | N/A | N/A | N/A | N/A |
| Enum | 9.658 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.31 / 0 allocs | 1029 / 89 B / 5 allocs | **17.6x** | 765.6 / 0 allocs | **13.1x** | 15672 / 15885 B / 76 allocs | **268.8x** |
| GTE | 4.095 / 0 allocs | 99.75 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| MinLength | 26.61 / 0 allocs | 110.4 / 0 allocs | **4.1x** | 252.9 / 32 B / 2 allocs | **9.5x** | N/A | N/A |
| UUID | 43.08 / 0 allocs | 422.7 / 0 allocs | **9.8x** | 322.9 / 0 allocs | **7.5x** | 15453 / 15542 B / 76 allocs | **358.7x** |
| MaxItems | 9.093 / 0 allocs | 127.0 / 0 allocs | **14.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 25.79 / 0 allocs | 126.8 / 0 allocs | **4.9x** | 254.1 / 32 B / 2 allocs | **9.9x** | 14942 / 15648 B / 81 allocs | **579.4x** |
| LT | 4.127 / 0 allocs | 101.1 / 0 allocs | **24.5x** | N/A | N/A | N/A | N/A |
| MinItems | 11.02 / 0 allocs | 127.5 / 0 allocs | **11.6x** | N/A | N/A | N/A | N/A |
| Alpha | 14.89 / 0 allocs | 386.6 / 0 allocs | **26.0x** | 98.25 / 0 allocs | **6.6x** | 19800 / 16937 B / 101 allocs | **1329.8x** |
| Required | 7.891 / 0 allocs | 132.4 / 0 allocs | **16.8x** | 1.462 / 0 allocs | **0.2x** | 14377 / 15488 B / 73 allocs | **1821.9x** |
| IPV4 | 33.30 / 0 allocs | 116.6 / 0 allocs | **3.5x** | N/A | N/A | N/A | N/A |
| Length | 9.596 / 0 allocs | 99.45 / 0 allocs | **10.4x** | 223.1 / 32 B / 2 allocs | **23.2x** | 14932 / 15616 B / 79 allocs | **1556.1x** |
| IPV6 | 80.60 / 0 allocs | 172.0 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 57.71 / 0 allocs | 470.9 / 144 B / 1 allocs | **8.2x** | 10266 / 145 B / 1 allocs | **177.9x** | 15407 / 15681 B / 77 allocs | **267.0x** |
| Numeric | 10.81 / 0 allocs | 78.76 / 0 allocs | **7.3x** | 109.4 / 0 allocs | **10.1x** | 15332 / 15574 B / 78 allocs | **1418.3x** |
| GT | 4.097 / 0 allocs | 101.4 / 0 allocs | **24.7x** | 74.01 / 0 allocs | **18.1x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 5.505 | 0 allocs |
| CELMultipleExpressions | 9.355 | 0 allocs |
| CELBasic | 9.072 | 0 allocs |
| CELCrossField | 6.275 | 0 allocs |
| CELStringLength | 1.326 | 0 allocs |
| CELNumericComparison | 1.831 | 0 allocs |

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
