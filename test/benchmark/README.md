# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-03-12  
**Platform:** Linux 6.14.0-1017-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	77465858	        14.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2664171	       464.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10601084	       112.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   60106	     20189 ns/op	   16936 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	271022649	         4.406 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	100000000	        12.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        11.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	147685318	         8.134 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19418270	        62.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1083 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1334667	       899.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   73670	     16037 ns/op	   15788 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	93313104	        12.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	214000555	         5.608 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10637168	       113.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13529566	        88.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	214110303	         5.606 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10643092	       113.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27126091	        44.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 8927676	       134.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13864135	        87.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6535041	       184.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	124163220	         9.692 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10707106	       112.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4799490	       249.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80263	     15084 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	213954668	         5.606 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11050347	       110.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	202880061	         5.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10309946	       118.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	98953214	        12.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8367937	       144.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	32208777	        36.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 9471019	       126.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4263028	       280.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   77445	     15196 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	99858790	        11.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8746920	       138.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48203900	        24.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9351094	       130.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4279137	       280.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	95197408	        12.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13710913	        87.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9344005	       128.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   76914	     15764 ns/op	   15573 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	128411992	         9.360 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8295058	       145.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642671250	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   82419	     14678 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19736496	        60.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2517272	       476.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  101382	     11930 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   76633	     15825 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	18421938	        64.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2589160	       470.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3277988	       364.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   76725	     15728 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.914 / 0 allocs | 118.9 / 0 allocs | **20.1x** | N/A | N/A | N/A | N/A |
| Enum | 12.79 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 62.78 / 0 allocs | 1083 / 89 B / 5 allocs | **17.3x** | 899.2 / 0 allocs | **14.3x** | 16037 / 15788 B / 76 allocs | **255.4x** |
| GTE | 5.606 / 0 allocs | 113.3 / 0 allocs | **20.2x** | N/A | N/A | N/A | N/A |
| MinLength | 24.91 / 0 allocs | 130.1 / 0 allocs | **5.2x** | 280.4 / 32 B / 2 allocs | **11.3x** | N/A | N/A |
| UUID | 64.67 / 0 allocs | 470.2 / 0 allocs | **7.3x** | 364.9 / 0 allocs | **5.6x** | 15728 / 15541 B / 76 allocs | **243.2x** |
| MaxItems | 12.14 / 0 allocs | 144.5 / 0 allocs | **11.9x** | N/A | N/A | N/A | N/A |
| MaxLength | 36.51 / 0 allocs | 126.6 / 0 allocs | **3.5x** | 280.1 / 32 B / 2 allocs | **7.7x** | 15196 / 15648 B / 81 allocs | **416.2x** |
| LT | 5.606 / 0 allocs | 110.2 / 0 allocs | **19.7x** | N/A | N/A | N/A | N/A |
| MinItems | 11.90 / 0 allocs | 138.6 / 0 allocs | **11.6x** | N/A | N/A | N/A | N/A |
| Alpha | 14.69 / 0 allocs | 464.8 / 0 allocs | **31.6x** | 112.8 / 0 allocs | **7.7x** | 20189 / 16936 B / 101 allocs | **1374.3x** |
| Required | 9.360 / 0 allocs | 145.2 / 0 allocs | **15.5x** | 1.870 / 0 allocs | **0.2x** | 14678 / 15488 B / 73 allocs | **1568.2x** |
| IPV4 | 44.36 / 0 allocs | 134.1 / 0 allocs | **3.0x** | N/A | N/A | N/A | N/A |
| Length | 9.692 / 0 allocs | 112.0 / 0 allocs | **11.6x** | 249.6 / 32 B / 2 allocs | **25.8x** | 15084 / 15616 B / 79 allocs | **1556.3x** |
| IPV6 | 87.82 / 0 allocs | 184.1 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 60.87 / 0 allocs | 476.0 / 144 B / 1 allocs | **7.8x** | 11930 / 147 B / 1 allocs | **196.0x** | 15825 / 15681 B / 77 allocs | **260.0x** |
| Numeric | 12.46 / 0 allocs | 87.79 / 0 allocs | **7.0x** | 128.5 / 0 allocs | **10.3x** | 15764 / 15573 B / 78 allocs | **1265.2x** |
| GT | 5.608 / 0 allocs | 113.3 / 0 allocs | **20.2x** | 88.79 / 0 allocs | **15.8x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.406 | 0 allocs |
| CELMultipleExpressions | 12.33 | 0 allocs |
| CELBasic | 11.53 | 0 allocs |
| CELCrossField | 8.134 | 0 allocs |
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
