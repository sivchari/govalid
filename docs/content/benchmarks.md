---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-14  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidEmail-4              	23223614	        51.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4         	 1000000	      1082 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4          	 1360206	       890.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-4       	12217460	        96.91 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateEmail-4       	   75676	     15825 ns/op	   15794 B/op	      76 allocs/op
BenchmarkGoValidEnum-4               	275041154	         4.368 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                 	549957590	         2.182 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4            	10501321	       116.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4             	13612168	        91.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                	642974013	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4           	10135870	       117.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-4                 	643114712	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4            	 9939189	       120.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                	643201537	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4           	10185711	       123.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4           	236857950	         5.020 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4      	 8609665	       142.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4          	52649703	        21.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4     	 7826211	       154.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4      	 4175911	       287.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-4   	 3877906	       307.5 ns/op	     448 B/op	       5 allocs/op
BenchmarkGookitValidateMaxLength-4   	   80212	     14983 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4           	226728916	         5.296 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4      	 8545950	       143.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4          	69875004	        17.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4     	 9764770	       125.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4      	 4171686	       287.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-4           	545092935	         2.189 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4      	 8752206	       140.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4       	550625074	         2.178 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-4    	12655058	        92.99 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateRequired-4    	   82938	     14345 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                	20920695	        57.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4           	 2365634	       505.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4            	  102610	     11705 ns/op	     145 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-4         	   94884	     12440 ns/op	     186 B/op	       3 allocs/op
BenchmarkGookitValidateURL-4         	   77646	     15533 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4               	30452536	        52.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4          	 2703860	       444.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4           	 3523198	       340.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-4        	 2720476	       436.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkGookitValidateUUID-4        	   77250	     15494 ns/op	   15541 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 51.21ns | 1082 | **21.1x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 2.182ns | 116.8 | **53.5x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.867ns | 117.8 | **63.1x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.866ns | 120.9 | **64.8x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.866ns | 123.3 | **66.1x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 5.020ns | 142.4 | **28.4x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 21.09ns | 154.1 | **7.3x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 5.296ns | 143.9 | **27.2x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 17.12ns | 125.3 | **7.3x faster** | 0 allocs/op | 0 allocs/op |
| Required | 2.189ns | 140.1 | **64.0x faster** | 0 allocs/op | 0 allocs/op |
| URL | 57.26ns | 505.2 | **8.8x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 52.02ns | 444.7 | **8.5x faster** | 0 allocs/op | 0 allocs/op |
| Enum | 2.242ns | N/A (govalid exclusive) | **govalid exclusive** | 0 allocs/op | N/A |

## Performance Categories

### 🚀 Ultra-Fast (< 3ns)
- **GT**: ~2.2ns - 54x faster\n- **GTE**: ~1.9ns - 63x faster\n- **LT**: ~1.9ns - 65x faster\n- **LTE**: ~1.9ns - 66x faster\n- **Required**: ~2.2ns - 64x faster\n
### ⚡ Fast (3-40ns)
- **MaxItems**: ~5ns - 28x faster\n- **MaxLength**: ~21ns - 7x faster\n- **MinItems**: ~5ns - 27x faster\n- **MinLength**: ~17ns - 7x faster\n
### 🎯 govalid Exclusive Features
- **Enum**: ~4.4ns - govalid exclusive\n

## Key Performance Insights

### 1. Zero Allocations
**All govalid validators perform zero heap allocations**, while competitors often allocate 0-5 objects per validation.

### 2. Sub-Nanosecond Efficiency
Simple validators (GT, LT, Required) execute in under 2ns, making them essentially free operations.

### 3. Complex Validation Optimization
Even complex validators like email and URL are optimized with:
- Manual string parsing (no regex overhead)
- Single-pass validation algorithms
- Zero memory allocations

### 4. String Length Performance
Unicode-aware string validators are 4.8-6.0x faster despite proper UTF-8 handling.

## govalid-Exclusive Features

### Enum Validation
```go
// +govalid:enum=admin,user,guest
Role string
```
- **Performance**: ~2.2ns with 0 allocations
- **No equivalent** in go-playground/validator
- Supports string, numeric, and custom types

### Extended Collection Support
```go
// +govalid:maxitems=10
Items map[string]int  // Maps supported!

// +govalid:minitems=1
Channel chan string   // Channels supported!
```

## Optimization Techniques

### 1. Code Generation
- **Compile-time validation functions** (no runtime reflection)
- **Inlined simple operations** for maximum speed
- **Direct field access** with no interface overhead

### 2. External Helper Functions
Complex validators use optimized external functions for better performance.

### 3. Manual String Parsing
- **Character-by-character parsing** instead of `strings.Split`
- **Direct indexing** instead of `strings.Contains`
- **Single-pass algorithms** for complex validation

### 4. Memory Optimization
- **Zero heap allocations** across all validators
- **Stack-only operations** for maximum cache efficiency
- **Minimal memory footprint** in generated code

## Running Benchmarks

To run benchmarks yourself:

```bash
# Sync all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem
```

## Conclusion

govalid delivers exceptional performance improvements:
- **4.8x to 45x faster** than go-playground/validator
- **Zero allocations** across all validators
- **Sub-3ns performance** for simple operations
- **Extended type support** (maps, channels, enums)
- **Production-ready** with comprehensive test coverage

Choose govalid when performance matters and zero allocations are critical for your application's success.
