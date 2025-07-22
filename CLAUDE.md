# Development Workflow for govalid

This document outlines the efficient development workflow established with our new automated registry system.

## 🚀 Implementation Pattern for New Markers

### 1. Feature Branch Setup
```bash
git checkout -b feature/{marker-name}-marker
```

### 2. Automated Validator Creation

#### A. One-Command Scaffold and Generation
```bash
# Create a new validator AND generate all registry files in one command:
make generate-validator MARKER=phoneNumber

# This:
# ✓ Generates scaffold: internal/validator/rules/phonenumber.go
# ✓ Updates internal/markers/markers_generated.go
# ✓ Creates internal/validator/registry/initializers/phonenumber.go
# ✓ Updates internal/validator/registry/initializers/all.go
# ✓ Updates internal/analyzers/govalid/registry_init.go
```

#### B. Implement Validator Logic
Edit the generated file with your validation logic:
```go
func (v *phonenumberValidator) Validate() string {
    // Return Go expression that evaluates to true when validation FAILS
    return fmt.Sprintf("!isValidPhoneNumber(t.%s)", v.FieldName())
}

func (v *phonenumberValidator) Imports() []string {
    return []string{"regexp"} // Add required imports
}
```

### 3. Testing Structure

#### A. Golden Tests (`internal/analyzers/govalid/testdata/src/{markername}/`)
- `{markername}.go` - Test input with marker comments
- `govalid.golden` - Expected generated output

#### B. Unit Tests (`test/unit/{markername}_test.go`)
```go
func Test{MarkerName}Validation(t *testing.T) {
    tests := []struct {
        name        string
        data        test.{MarkerName}
        expectError bool
    }{
        {"valid", test.{MarkerName}{Field: "valid_value"}, false},
        {"limit_minus_one", test.{MarkerName}{Field: "boundary-1"}, false},
        {"exactly_at_limit", test.{MarkerName}{Field: "boundary"}, false},
        {"limit_plus_one", test.{MarkerName}{Field: "boundary+1"}, true},
    }
    // Test both govalid and go-playground/validator
}
```

#### C. Benchmark Tests (`test/benchmark/benchmark_{markername}_test.go`)
```go
func BenchmarkGoValid{MarkerName}(b *testing.B) {
    instance := test.{MarkerName}{Field: "test_value"}
    b.ResetTimer()
    for b.Loop() {
        err := test.Validate{MarkerName}(&instance)
        if err != nil {
            b.Fatal("unexpected error:", err)
        }
    }
    b.StopTimer()
}

func BenchmarkGoPlayground{MarkerName}(b *testing.B) {
    validate := validator.New()
    instance := test.{MarkerName}{Field: "test_value"}
    b.ResetTimer()
    for b.Loop() {
        err := validate.Struct(&instance)
        if err != nil {
            b.Fatal("unexpected error:", err)
        }
    }
    b.StopTimer()
}
```

### 4. Test Execution Order
```bash
# 1. Build and install updated binary
go install ./cmd/govalid/

# 2. Generate test files
cd test && go generate

# 3. Run golden tests
cd .. && go test ./internal/analyzers/govalid/ -v

# 4. Run unit tests
cd test && go test ./unit/ -v

# 5. Run benchmarks
go test ./benchmark/ -bench=Benchmark.*{MarkerName} -benchmem

# 6. Update benchmark README
# Edit test/benchmark/README.md with results

# 7. Run lint checks and fix any issues
cd .. && make golangci-lint

# 8. Re-run benchmarks after any optimization changes
# If code changes were made to fix lint issues or optimize performance:
cd test && go test ./benchmark/ -bench=Benchmark.*{MarkerName} -benchmem

# 9. Update benchmark README again if performance changed
# Edit test/benchmark/README.md with updated results
```

### 5. Documentation Updates
- Update main README.md with marker explanation
- Update benchmark/README.md with performance results
- Document any behavior differences from go-playground/validator

## 🆕 New System Architecture

### Registry-Based Validator Discovery

The new system eliminates manual registration through an automated discovery and registry pattern:

```go
// internal/validator/registry/registry.go
type Registry interface {
    Markers() []string
    Validator(marker string) (ValidatorFactory, error)
    Init() error
}

// internal/validator/registry/initializers/
// Each validator has its own initializer automatically generated
type PhoneNumberInitializer struct{}

func (p PhoneNumberInitializer) Marker() string {
    return markers.GoValidMarkerPhoneNumber
}

func (p PhoneNumberInitializer) Init() registry.ValidatorFactory {
    return rules.ValidatePhoneNumber
}
```

### Automatic Code Generation Flow

```
make generate-validator MARKER=phonenumber
                    ↓
    ┌───────────────────────────────────┐
    │ Creates scaffold in rules/        │
    │ • internal/validator/rules/       │
    │   phonenumber.go                  │
    └───────────────────────────────────┘
                    ↓
    ┌───────────────────────────────────┐
    │ Automatically discovers all       │
    │ validators in rules/              │
    └───────────────────────────────────┘
                    ↓
    ┌───────────────────────────────────┐
    │ Generates:                        │
    │ • markers_generated.go            │
    │ • initializers/*.go               │
    │ • all.go                          │
    │ • registry_init.go                │
    └───────────────────────────────────┘
                    ↓
    ┌───────────────────────────────────┐
    │ makeValidator uses registry to    │
    │ dynamically resolve validators    │
    └───────────────────────────────────┘
```

### Template Management

All templates are now externalized and embedded:

```
cmd/generate-validators/templates/
├── initializer.go.tmpl      # Validator initializer template
├── all.go.tmpl              # All initializers aggregation
├── registry_init.go.tmpl    # Registry initialization
├── markers.go.tmpl          # Marker definitions
└── validator.go.tmpl        # Validator scaffold template

internal/analyzers/govalid/templates/
└── validation.go.tmpl       # Generated validation function template
```

## 🔧 Key Technical Patterns

### Validation Helper Functions

For complex validation logic that generates large amounts of code, use external helper functions:

```go
// Location: validation/validationhelper/{validator_name}.go
package validationhelper

func IsValid{ValidationName}(input string) bool {
    // Complex validation logic
    return true
}
```

**When to use helper functions:**
- Complex validation logic (> 50 lines)
- Multiple string operations or loops
- Functions that would not be inlined by the compiler
- Logic that benefits from centralized maintenance

**When to use inline generation:**
- Simple comparisons (GT, LT, Required)
- Single-line validations
- Performance-critical simple operations

**Helper function integration:**
```go
// In validator implementation
func (v *{validator}Validator) Validate() string {
    return fmt.Sprintf("!validationhelper.IsValid{ValidationName}(t.%s)", v.FieldName())
}

func (v *{validator}Validator) Imports() []string {
    return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// No need for GeneratorMemory management - external function handles it
func (v *{validator}Validator) Err() string {
    // Only generate error variables, no inline functions
}
```

**Performance optimization cycle for helper functions:**
1. **Initial implementation**: Focus on correctness and functionality
2. **Lint compliance**: Run `make golangci-lint` and fix all issues
3. **Function decomposition**: Break complex functions into smaller, optimized components
4. **Benchmark verification**: Ensure optimizations improve performance
5. **Documentation update**: Update benchmark README with improved results

**Helper function best practices:**
- **Decompose complex logic**: Break functions into smaller, focused components
- **Minimize allocations**: Avoid `strings.Split`, `strings.Contains` in hot paths
- **Use manual parsing**: Character-by-character parsing for zero allocations
- **Optimize for compiler**: Small functions are more likely to be inlined
- **Lint compliance**: Ensure all helper functions pass golangci-lint checks

## 🔧 Advanced Technical Patterns

### Error Variable Naming Pattern

**Since Issue #72 Implementation**: Error variables now include struct name prefixes to prevent naming conflicts:

```go
// Before (Issue #72)
ErrNameRequiredValidation = errors.New("field Name is required")

// After (Issue #72)
ErrUserNameRequiredValidation = errors.New("field Name is required")
ErrProductNameRequiredValidation = errors.New("field Name is required")
```

**Benefits:**
- **Prevents naming conflicts**: Multiple structs can have fields with same names
- **Improves code clarity**: Clear which struct the error belongs to
- **Maintains backward compatibility**: Generated code compiles without changes

**Implementation Pattern:**
```go
// Validator struct includes struct name
type requiredValidator struct {
    pass       *codegen.Pass
    field      *ast.Field
    structName string  // Added for Issue #72
}

// Error variable generation uses struct name prefix
func (r *requiredValidator) ErrVariable() string {
    return strings.ReplaceAll("Err@RequiredValidation", "@", r.structName+r.FieldName())
}
```

### Interface-Based Import System
```go
// Validator interface
type Validator interface {
    Validate() string
    FieldName() string
    Err() string
    ErrVariable() string
    Imports() []string  // Dynamic import declaration
}

// Collector function
func collectImportPackages(metadata []*AnalyzedMetadata) map[string]struct{} {
    packages := make(map[string]struct{})
    for _, meta := range metadata {
        for _, validator := range meta.Validators {
            for _, pkg := range validator.Imports() {
                packages[pkg] = struct{}{}
            }
        }
    }
    return packages
}
```

### Template Integration
```go
// Template data structure
type TemplateData struct {
    PackageName     string
    TypeName        string
    Metadata        []*AnalyzedMetadata
    ImportPackages  map[string]struct{}  // Dynamic imports
}

// Template usage
{{- range $pkg, $_ := .ImportPackages }}
"{{ $pkg }}"
{{- end }}
```

## 📊 Benchmark Best Practices (Go 1.24+)

### Correct Benchmark Structure
```go
func BenchmarkFunction(b *testing.B) {
    // Setup (runs once)
    instance := setupData()
    
    b.ResetTimer()  // Exclude setup time
    for b.Loop() {  // Go 1.24+ preferred method
        result := functionUnderTest(instance)
        if result != expected {
            b.Fatal("unexpected result")
        }
    }
    b.StopTimer()  // Optional, for cleanup exclusion
}
```

### Key Points:
- **Use `b.Loop()`** instead of `for i := 0; i < b.N; i++` (Go 1.24+)
- **Call `b.ResetTimer()`** after setup to exclude initialization time
- **Verify results** to ensure compiler doesn't optimize away the work
- **Use `b.StopTimer()`** if cleanup time should be excluded

## 🎯 Performance Expectations

Based on MaxLength implementation:
- **Simple validators (GT/LT/Required)**: ~1-2ns, 0 allocs
- **String validators (MaxLength)**: ~14ns, 0 allocs  
- **Improvement over go-playground/validator**: 5x to 50x faster
- **Memory efficiency**: Always 0 allocations vs competitor's 0-5 allocs

## 🔍 Validation Behavior Patterns

### Required Field Handling
- **nil slice/map/chan**: Invalid (required violation)
- **empty slice []**: Valid (initialized)
- **zero values**: Invalid for primitives, follow Go zero-value semantics
- **Arrays**: Check `len(array) == 0` (arrays can't be nil)

### String Validation
- **Use `utf8.RuneCountInString()`** for character counting (not `len()`)
- **Matches go-playground/validator Unicode behavior**

## ⚠️ Common Pitfalls to Avoid

1. **Benchmark Measurement**: Always verify actual processing occurs
2. **Import Management**: Use interface-based system, not hardcoded strings.Contains()
3. **Test Structure**: Keep boundary value tests simple and focused
4. **Error Generation**: Use GeneratorMemory to avoid duplicate error definitions
5. **Template Formatting**: Use consistent indentation and error message format
6. **Lint Compliance**: Always run `make golangci-lint` after implementation changes
7. **Performance Regression**: Re-run benchmarks after any code changes (lint fixes, refactoring)
8. **Documentation Updates**: Always update benchmark README if performance numbers change

## 🔄 Post-Implementation Workflow

### Lint and Optimization Cycle
```bash
# After initial implementation
make golangci-lint

# If lint issues found:
# 1. Fix lint issues (may involve code refactoring)
# 2. Re-run benchmarks to check for performance changes
go test ./benchmark/ -bench=Benchmark.*{MarkerName} -benchmem

# 3. Update benchmark README if numbers changed
# 4. Verify tests still pass
go test ./unit/ -v

# 5. Re-run lint to ensure fixes are correct
make golangci-lint
```

### Performance Monitoring
- **Before optimization**: Record baseline performance
- **After lint fixes**: Check for performance impact (usually positive due to better compiler optimization)
- **Document improvements**: Update README with new performance numbers
- **Verify against competitors**: Ensure go-playground/validator comparison is still accurate

## 📝 Commit Message Pattern
```
{Action} {MarkerName} marker implementation

- Add {markername} validator with {specific_features}
- Implement {validation_logic} with {performance_notes}
- Add comprehensive unit tests with boundary value testing
- Add benchmarks showing {performance_improvement}
- Update documentation and README

🤖 Generated with [Claude Code](https://claude.ai/code)

Co-Authored-By: Claude <noreply@anthropic.com>
```

## 🔄 Next Markers to Implement

From GitHub issues (in priority order):
1. MaxItems marker (Issue #7) - Array/slice length validation ✓
2. MinLength marker (Issue #11) - String minimum length ✓
3. MinItems marker (Issue #10) - Array/slice minimum length ✓
4. GTE marker (Issue #15) - Greater than or equal numeric validation ✓
5. LTE marker (Issue #16) - Less than or equal numeric validation ✓
6. Enum marker (Issue #17) - Enumeration validation ✓

All basic markers have been implemented! Future work:
- Custom validators support
- Performance optimizations
- Additional complex validators based on user feedback

Follow this pattern for each implementation to maintain consistency and quality.