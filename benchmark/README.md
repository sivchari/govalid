# Benchmark Results: govalid vs go-playground/validator

This document provides a structured comparison of the benchmark results for `govalid` and `go-playground/validator`. The benchmarks focus on the `required`, `lt`, and `gt` validation markers, with metrics derived from tests conducted using `-benchmem` and `-count 10`.

## Specifications

- **Processor:** Apple M3 Max
- **Architecture:** arm64

---

## Benchmark Summary

### `govalid`

#### Required Validation
- **Average Time/op:** ~1.75 ns
- **Memory Allocations/op:** 0
- **Bytes Allocated/op:** 0

#### LT Validation
- **Average Time/op:** ~1.75 ns
- **Memory Allocations/op:** 0
- **Bytes Allocated/op:** 0

#### GT Validation
- **Average Time/op:** ~1.75 ns
- **Memory Allocations/op:** 0
- **Bytes Allocated/op:** 0

### `go-playground/validator`

#### Required Validation
- **Average Time/op:** ~245 ns
- **Memory Allocations/op:** 8
- **Bytes Allocated/op:** 440

#### LT Validation
- **Average Time/op:** ~135 ns
- **Memory Allocations/op:** 4
- **Bytes Allocated/op:** 192

#### GT Validation
- **Average Time/op:** ~135 ns
- **Memory Allocations/op:** 4
- **Bytes Allocated/op:** 192

---

## Observations

### Performance Comparison

#### Required Validation
- `govalid` is **140x faster** than `go-playground/validator`.
- `govalid` performs no memory allocations, while `go-playground/validator` allocates 440 bytes across 8 blocks.

#### LT Validation
- `govalid` is **77x faster** than `go-playground/validator`.
- `govalid` performs no memory allocations, while `go-playground/validator` allocates 192 bytes across 4 blocks.

#### GT Validation
- `govalid` is **77x faster** than `go-playground/validator`.
- `govalid` performs no memory allocations, while `go-playground/validator` allocates 192 bytes across 4 blocks.

---

## Structured Benchmarks

### `Required` Validation
| Library                 | Avg Time/op (ns) | Allocations/op | Bytes/op |
|-------------------------|------------------|----------------|----------|
| govalid                | ~1.75            | 0              | 0        |
| go-playground/validator | ~245             | 8              | 440      |

### `LT` Validation
| Library                 | Avg Time/op (ns) | Allocations/op | Bytes/op |
|-------------------------|------------------|----------------|----------|
| govalid                | ~1.75            | 0              | 0        |
| go-playground/validator | ~135             | 4              | 192      |

### `GT` Validation
| Library                 | Avg Time/op (ns) | Allocations/op | Bytes/op |
|-------------------------|------------------|----------------|----------|
| govalid                | ~1.75            | 0              | 0        |
| go-playground/validator | ~135             | 4              | 192      |

---

## Conclusion

The `govalid` package demonstrates superior performance for validation tasks compared to `go-playground/validator`. Its zero memory allocation and extremely low execution time make it ideal for performance-critical applications. The structured presentation ensures easy updates for future marker additions.

