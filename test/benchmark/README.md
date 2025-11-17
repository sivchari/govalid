# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-11-17  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124465195	         9.641 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2561749	       473.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11256972	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56446	     21338 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644549823	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296390485	         4.061 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296421126	         4.051 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	348322929	         3.431 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21520556	        55.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1094 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1320170	       912.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68668	     17361 ns/op	   15858 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275109340	         4.372 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481368273	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11313842	       106.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13709576	        88.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481728891	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11016687	       109.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30860154	        39.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9232340	       129.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14131914	        85.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6914020	       173.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160512554	         7.474 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9988878	       120.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4846678	       248.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73257	     16309 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	482192634	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11035357	       108.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428237139	         2.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11245808	       106.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	182846058	         6.553 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8288018	       144.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35850174	        33.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8638849	       139.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4150077	       289.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73766	     16380 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202859095	         5.919 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8491552	       141.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52022101	        23.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10111590	       119.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4206610	       284.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157352798	         7.615 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13145541	        91.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9429699	       129.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71228	     16787 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320638898	         3.742 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8019147	       149.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642303189	         1.868 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75075	     15933 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20676901	        57.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2507934	       479.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103687	     11672 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71126	     17039 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23397001	        51.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2675090	       450.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3456387	       352.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71002	     16957 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.804 / 0 allocs | 106.9 / 0 allocs | **38.1x** | N/A | N/A | N/A | N/A |
| Enum | 4.372 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.45 / 0 allocs | 1094 / 89 B / 5 allocs | **19.7x** | 912.8 / 0 allocs | **16.5x** | 17361 / 15858 B / 76 allocs | **313.1x** |
| GTE | 2.490 / 0 allocs | 109.1 / 0 allocs | **43.8x** | N/A | N/A | N/A | N/A |
| MinLength | 23.04 / 0 allocs | 119.0 / 0 allocs | **5.2x** | 284.9 / 32 B / 2 allocs | **12.4x** | N/A | N/A |
| UUID | 51.61 / 0 allocs | 450.3 / 0 allocs | **8.7x** | 352.8 / 0 allocs | **6.8x** | 16957 / 15542 B / 76 allocs | **328.6x** |
| MaxItems | 6.553 / 0 allocs | 144.6 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.36 / 0 allocs | 139.2 / 0 allocs | **4.2x** | 289.6 / 32 B / 2 allocs | **8.7x** | 16380 / 15648 B / 81 allocs | **491.0x** |
| LT | 2.492 / 0 allocs | 108.9 / 0 allocs | **43.7x** | N/A | N/A | N/A | N/A |
| MinItems | 5.919 / 0 allocs | 141.2 / 0 allocs | **23.9x** | N/A | N/A | N/A | N/A |
| Alpha | 9.641 / 0 allocs | 473.7 / 0 allocs | **49.1x** | 108.0 / 0 allocs | **11.2x** | 21338 / 16937 B / 101 allocs | **2213.3x** |
| Required | 3.742 / 0 allocs | 149.5 / 0 allocs | **40.0x** | 1.868 / 0 allocs | **0.5x** | 15933 / 15488 B / 73 allocs | **4257.9x** |
| IPV4 | 39.35 / 0 allocs | 129.9 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 7.474 / 0 allocs | 120.4 / 0 allocs | **16.1x** | 248.0 / 32 B / 2 allocs | **33.2x** | 16309 / 15616 B / 79 allocs | **2182.1x** |
| IPV6 | 85.55 / 0 allocs | 173.1 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.73 / 0 allocs | 479.2 / 144 B / 1 allocs | **8.3x** | 11672 / 145 B / 1 allocs | **202.2x** | 17039 / 15681 B / 77 allocs | **295.1x** |
| Numeric | 7.615 / 0 allocs | 91.50 / 0 allocs | **12.0x** | 129.2 / 0 allocs | **17.0x** | 16787 / 15574 B / 78 allocs | **2204.5x** |
| GT | 2.492 / 0 allocs | 106.1 / 0 allocs | **42.6x** | 88.33 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.870 | 0 allocs |
| CELMultipleExpressions | 4.061 | 0 allocs |
| CELBasic | 4.051 | 0 allocs |
| CELCrossField | 3.431 | 0 allocs |
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
