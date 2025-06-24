# Benchmark Results: govalid vs go-playground/validator

This document provides a detailed comparison of the benchmark results for `govalid` and `go-playground/validator` focusing on the `required` validation marker. The tests were run with `-benchmem` and `-count 10` options to ensure comprehensive performance data, including memory allocations.

## Hardware Specification

- **Processor:** Apple M3 Max
- **Architecture:** arm64

## Benchmark Results

### govalid

**Benchmark Function:** `BenchmarkGovalidValidatorRequired`

| Trial | Operations | Time/op | Allocations/op | Bytes/op |
|-------|------------|---------|----------------|----------|
| 1     | 666,628,705 | 1.751 ns | 0              | 0        |
| 2     | 690,410,060 | 1.758 ns | 0              | 0        |
| 3     | 677,892,536 | 1.753 ns | 0              | 0        |
| 4     | 681,072,290 | 1.796 ns | 0              | 0        |
| 5     | 675,411,998 | 1.762 ns | 0              | 0        |
| 6     | 683,944,821 | 1.783 ns | 0              | 0        |
| 7     | 676,838,354 | 1.754 ns | 0              | 0        |
| 8     | 692,204,857 | 1.754 ns | 0              | 0        |
| 9     | 691,844,354 | 1.746 ns | 0              | 0        |
| 10    | 691,850,835 | 1.742 ns | 0              | 0        |

### go-playground/validator

**Benchmark Function:** `BenchmarkPlaygroundValidatorRequired`

| Trial | Operations | Time/op   | Allocations/op | Bytes/op |
|-------|------------|-----------|----------------|----------|
| 1     | 5,107,802   | 233.7 ns  | 8              | 432      |
| 2     | 5,050,735   | 235.0 ns  | 8              | 432      |
| 3     | 5,110,531   | 240.2 ns  | 8              | 432      |
| 4     | 5,090,628   | 235.0 ns  | 8              | 432      |
| 5     | 4,965,501   | 240.2 ns  | 8              | 432      |
| 6     | 5,113,801   | 236.3 ns  | 8              | 432      |
| 7     | 5,105,175   | 236.4 ns  | 8              | 432      |
| 8     | 5,100,423   | 237.6 ns  | 8              | 432      |
| 9     | 5,054,350   | 236.0 ns  | 8              | 432      |
| 10    | 5,136,982   | 235.1 ns  | 8              | 432      |

## Observations

- **govalid**:
  - Extremely fast, with an average execution time of ~1.75 ns per operation.
  - Consumes zero bytes and performs zero allocations per operation.

- **go-playground/validator**:
  - Slower, with an average execution time of ~235 ns per operation.
  - Allocates 8 blocks per operation, consuming 432 bytes.

## Conclusion

The `govalid` package demonstrates superior performance compared to `go-playground/validator`, particularly for the `required` validation marker. It achieves faster execution times and zero memory allocation, making it ideal for high-performance applications.

