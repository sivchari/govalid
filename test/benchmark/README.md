# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-12-24  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	123790930	         9.697 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2572950	       469.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11229477	       108.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55755	     21266 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644407250	         1.851 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	276177000	         4.160 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295386060	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320327034	         3.742 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21659466	        57.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1099 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1358004	       883.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69178	     17253 ns/op	   15880 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	255716114	         4.683 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	422836791	         2.814 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11307860	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13661169	        88.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	427971492	         2.806 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10964242	       109.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30214101	        38.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9248691	       131.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14032346	        85.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6947787	       173.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	147248554	         8.151 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9928974	       120.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4820006	       249.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73659	     16261 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	427688871	         2.813 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11258468	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385121257	         3.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11112056	       107.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202627372	         5.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8253514	       145.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37249813	        31.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8575851	       139.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4188495	       286.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73533	     16371 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	175077242	         6.862 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8519740	       140.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	45149082	        24.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10119726	       118.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4172503	       284.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	123016434	         9.749 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13686003	        87.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9402465	       130.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   70941	     16700 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296392832	         4.050 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8036241	       149.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	640938771	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   74745	     15798 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20423626	        57.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2448961	       487.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102344	     11768 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70395	     16979 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	19197388	        52.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2687061	       447.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3582789	       346.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70045	     16912 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.113 / 0 allocs | 107.0 / 0 allocs | **34.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.683 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.22 / 0 allocs | 1099 / 89 B / 5 allocs | **19.2x** | 883.9 / 0 allocs | **15.4x** | 17253 / 15880 B / 76 allocs | **301.5x** |
| GTE | 2.806 / 0 allocs | 109.5 / 0 allocs | **39.0x** | N/A | N/A | N/A | N/A |
| MinLength | 24.84 / 0 allocs | 118.9 / 0 allocs | **4.8x** | 284.9 / 32 B / 2 allocs | **11.5x** | N/A | N/A |
| UUID | 52.34 / 0 allocs | 447.9 / 0 allocs | **8.6x** | 346.1 / 0 allocs | **6.6x** | 16912 / 15542 B / 76 allocs | **323.1x** |
| MaxItems | 5.917 / 0 allocs | 145.7 / 0 allocs | **24.6x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.76 / 0 allocs | 139.1 / 0 allocs | **4.4x** | 286.0 / 32 B / 2 allocs | **9.0x** | 16371 / 15648 B / 81 allocs | **515.5x** |
| LT | 2.813 / 0 allocs | 105.6 / 0 allocs | **37.5x** | N/A | N/A | N/A | N/A |
| MinItems | 6.862 / 0 allocs | 140.5 / 0 allocs | **20.5x** | N/A | N/A | N/A | N/A |
| Alpha | 9.697 / 0 allocs | 469.5 / 0 allocs | **48.4x** | 108.2 / 0 allocs | **11.2x** | 21266 / 16937 B / 101 allocs | **2193.0x** |
| Required | 4.050 / 0 allocs | 149.9 / 0 allocs | **37.0x** | 1.870 / 0 allocs | **0.5x** | 15798 / 15488 B / 73 allocs | **3900.7x** |
| IPV4 | 38.98 / 0 allocs | 131.0 / 0 allocs | **3.4x** | N/A | N/A | N/A | N/A |
| Length | 8.151 / 0 allocs | 120.8 / 0 allocs | **14.8x** | 249.5 / 32 B / 2 allocs | **30.6x** | 16261 / 15616 B / 79 allocs | **1995.0x** |
| IPV6 | 85.15 / 0 allocs | 173.3 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.98 / 0 allocs | 487.8 / 144 B / 1 allocs | **8.4x** | 11768 / 145 B / 1 allocs | **203.0x** | 16979 / 15681 B / 77 allocs | **292.8x** |
| Numeric | 9.749 / 0 allocs | 87.62 / 0 allocs | **9.0x** | 130.1 / 0 allocs | **13.3x** | 16700 / 15574 B / 78 allocs | **1713.0x** |
| GT | 2.814 / 0 allocs | 106.6 / 0 allocs | **37.9x** | 88.27 / 0 allocs | **31.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.851 | 0 allocs |
| CELMultipleExpressions | 4.160 | 0 allocs |
| CELBasic | 4.052 | 0 allocs |
| CELCrossField | 3.742 | 0 allocs |
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
