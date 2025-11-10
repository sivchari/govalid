# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-10  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124402392	         9.646 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2607459	       466.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11120148	       109.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   51214	     21710 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	642949932	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296399484	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296033667	         4.055 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350442292	         3.426 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21623012	        55.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1104 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1309041	       915.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   67730	     17714 ns/op	   15860 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275282664	         4.360 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481799775	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11197939	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13556689	        88.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482050182	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11281632	       106.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160575669	         7.470 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9607981	       124.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4764471	       249.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   72710	     16463 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481699926	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10992028	       105.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428003096	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10919542	       109.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	182731650	         6.559 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8375422	       143.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35756702	        33.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8659742	       139.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4191398	       286.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74014	     16425 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202871932	         5.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8336589	       144.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51845000	        23.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9983353	       119.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4212717	       287.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157548565	         7.603 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12174086	        99.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9492668	       128.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   69511	     16978 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	319039125	         3.743 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7972176	       149.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643422223	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   74026	     15987 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20661374	        57.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2446572	       489.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  101948	     11767 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70132	     17155 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24761434	        48.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2650411	       451.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3579603	       349.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70214	     17053 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.802 / 0 allocs | 109.9 / 0 allocs | **39.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.360 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.72 / 0 allocs | 1104 / 89 B / 5 allocs | **19.8x** | 915.9 / 0 allocs | **16.4x** | 17714 / 15860 B / 76 allocs | **317.9x** |
| GTE | 2.490 / 0 allocs | 106.7 / 0 allocs | **42.9x** | N/A | N/A | N/A | N/A |
| MinLength | 23.02 / 0 allocs | 119.4 / 0 allocs | **5.2x** | 287.1 / 32 B / 2 allocs | **12.5x** | N/A | N/A |
| UUID | 48.56 / 0 allocs | 451.3 / 0 allocs | **9.3x** | 349.8 / 0 allocs | **7.2x** | 17053 / 15542 B / 76 allocs | **351.2x** |
| MaxItems | 6.559 / 0 allocs | 143.9 / 0 allocs | **21.9x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.39 / 0 allocs | 139.2 / 0 allocs | **4.2x** | 286.3 / 32 B / 2 allocs | **8.6x** | 16425 / 15648 B / 81 allocs | **491.9x** |
| LT | 2.492 / 0 allocs | 105.9 / 0 allocs | **42.5x** | N/A | N/A | N/A | N/A |
| MinItems | 5.914 / 0 allocs | 144.2 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.646 / 0 allocs | 466.0 / 0 allocs | **48.3x** | 109.6 / 0 allocs | **11.4x** | 21710 / 16937 B / 101 allocs | **2250.7x** |
| Required | 3.743 / 0 allocs | 149.6 / 0 allocs | **40.0x** | 1.868 / 0 allocs | **0.5x** | 15987 / 15488 B / 73 allocs | **4271.2x** |
| Length | 7.470 / 0 allocs | 124.7 / 0 allocs | **16.7x** | 249.7 / 32 B / 2 allocs | **33.4x** | 16463 / 15616 B / 79 allocs | **2203.9x** |
| URL | 57.61 / 0 allocs | 489.1 / 144 B / 1 allocs | **8.5x** | 11767 / 147 B / 1 allocs | **204.3x** | 17155 / 15681 B / 77 allocs | **297.8x** |
| Numeric | 7.603 / 0 allocs | 99.09 / 0 allocs | **13.0x** | 128.0 / 0 allocs | **16.8x** | 16978 / 15574 B / 78 allocs | **2233.1x** |
| GT | 2.490 / 0 allocs | 106.6 / 0 allocs | **42.8x** | 88.09 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.869 | 0 allocs |
| CELMultipleExpressions | 4.048 | 0 allocs |
| CELBasic | 4.055 | 0 allocs |
| CELCrossField | 3.426 | 0 allocs |
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
