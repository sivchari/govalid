# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-02-17  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124166246	         9.664 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2595321	       464.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11264083	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55311	     21586 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	652329520	         1.840 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	294677804	         4.065 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295278594	         4.058 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320006888	         3.765 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21664406	        59.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1090 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1351796	       889.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69050	     17416 ns/op	   15877 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256667638	         4.674 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427540752	         2.816 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11208024	       107.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13709418	        88.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	426825878	         2.806 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11006048	       112.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30569211	        39.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9240052	       130.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14160849	        87.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6931982	       174.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147560697	         8.126 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10021808	       120.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4888586	       246.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73968	     16364 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	424009507	         2.815 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11361561	       105.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	383695995	         3.122 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11187834	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202465897	         5.921 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8200248	       146.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37457002	        31.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8625096	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4202091	       285.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   71424	     16621 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	175073970	         6.862 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8472150	       142.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48596800	        24.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9936924	       121.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4225189	       283.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124280866	         9.659 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13739083	        87.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9347227	       128.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   68774	     16998 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296448847	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8019379	       149.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642022964	         1.874 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75781	     15853 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20774944	        57.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2511846	       478.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102722	     11897 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70581	     17117 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	25847869	        53.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2694936	       446.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3563232	       336.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70190	     17070 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.122 / 0 allocs | 107.4 / 0 allocs | **34.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.674 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.03 / 0 allocs | 1090 / 89 B / 5 allocs | **18.5x** | 889.3 / 0 allocs | **15.1x** | 17416 / 15877 B / 76 allocs | **295.0x** |
| GTE | 2.806 / 0 allocs | 112.3 / 0 allocs | **40.0x** | N/A | N/A | N/A | N/A |
| MinLength | 24.66 / 0 allocs | 121.2 / 0 allocs | **4.9x** | 283.2 / 32 B / 2 allocs | **11.5x** | N/A | N/A |
| UUID | 53.79 / 0 allocs | 446.7 / 0 allocs | **8.3x** | 336.8 / 0 allocs | **6.3x** | 17070 / 15542 B / 76 allocs | **317.3x** |
| MaxItems | 5.921 / 0 allocs | 146.3 / 0 allocs | **24.7x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.79 / 0 allocs | 139.3 / 0 allocs | **4.4x** | 285.7 / 32 B / 2 allocs | **9.0x** | 16621 / 15648 B / 81 allocs | **522.8x** |
| LT | 2.815 / 0 allocs | 105.9 / 0 allocs | **37.6x** | N/A | N/A | N/A | N/A |
| MinItems | 6.862 / 0 allocs | 142.3 / 0 allocs | **20.7x** | N/A | N/A | N/A | N/A |
| Alpha | 9.664 / 0 allocs | 464.1 / 0 allocs | **48.0x** | 107.7 / 0 allocs | **11.1x** | 21586 / 16937 B / 101 allocs | **2233.7x** |
| Required | 4.048 / 0 allocs | 149.7 / 0 allocs | **37.0x** | 1.874 / 0 allocs | **0.5x** | 15853 / 15488 B / 73 allocs | **3916.3x** |
| IPV4 | 39.07 / 0 allocs | 130.4 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.126 / 0 allocs | 120.1 / 0 allocs | **14.8x** | 246.2 / 32 B / 2 allocs | **30.3x** | 16364 / 15616 B / 79 allocs | **2013.8x** |
| IPV6 | 87.99 / 0 allocs | 174.9 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.71 / 0 allocs | 478.6 / 144 B / 1 allocs | **8.3x** | 11897 / 146 B / 1 allocs | **206.2x** | 17117 / 15681 B / 77 allocs | **296.6x** |
| Numeric | 9.659 / 0 allocs | 87.74 / 0 allocs | **9.1x** | 128.4 / 0 allocs | **13.3x** | 16998 / 15574 B / 78 allocs | **1759.8x** |
| GT | 2.816 / 0 allocs | 107.5 / 0 allocs | **38.2x** | 88.32 / 0 allocs | **31.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.840 | 0 allocs |
| CELMultipleExpressions | 4.065 | 0 allocs |
| CELBasic | 4.058 | 0 allocs |
| CELCrossField | 3.765 | 0 allocs |
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
