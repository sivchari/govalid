# govalid

## Overview

govalid is a Go package designed to generate type-safe validation code for structs based on markers. It provides a mechanism to apply validation rules directly in the code by marking struct fields with specific markers. The tool processes these markers and generates corresponding validation functions.

## Installation

To install govalid, use:

```bash
# Clone the repository
$ git clone https://github.com/sivchari/govalid.git

# Navigate to the project directory
$ cd govalid

# Run installation commands
$ go install ./...
```

## Usage

To use govalid, follow these steps:

1. Define your struct with field markers.
2. Run the govalid generation tool to create type-safe validation code.
3. Use the generated validation functions in your application.

### Example

```go
// Example struct with markers
// +govalid:required
type Person struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Generated code will include validation functions like:
func ValidatePerson(t *Person) error {
    if t == nil {
        return ErrNilPerson
    }

    if t.Name == "" {
        return ErrNameRequiredValidation
    }

    if t.Email == "" {
        return ErrEmailRequiredValidation
    }

    return nil
}
```

## Supported Markers

govalid supports the following markers:

### `govalid:required`
- **Description**: Ensures that the field is not empty or nil.
- **Example**:
  ```go
  // +govalid:required
  type User struct {
      Username string `json:"username"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateUser(t *User) error {
      if t == nil {
          return ErrNilUser
      }

      if t.Username == "" {
          return ErrUsernameRequiredValidation
      }

      return nil
  }
  ```

### `govalid:lt`
- **Description**: Ensures that a numeric field is less than a specified value.
- **Example**:
  ```go
  // +govalid:lt=18
  type Profile struct {
      Age int `json:"age"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateProfile(t *Profile) error {
      if t == nil {
          return ErrNilProfile
      }

      if t.Age >= 18 {
          return ErrAgeLtValidation
      }

      return nil
  }
  ```

### `govalid:gt`
- **Description**: Ensures that a numeric field is greater than a specified value.
- **Example**:
  ```go
  // +govalid:gt=100
  type Profile struct {
      Age int `json:"age"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateProfile(t *Profile) error {
      if t == nil {
          return ErrNilProfile
      }

      if t.Age <= 100 {
          return ErrAgeGtValidation
      }

      return nil
  }
  ```

### `govalid:maxlength`
- **Description**: Ensures that a string field's length does not exceed the specified maximum value (Unicode-aware).
- **Example**:
  ```go
  type User struct {
      // +govalid:maxlength=50
      Username string `json:"username"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateUser(t *User) error {
      if t == nil {
          return ErrNilUser
      }

      if utf8.RuneCountInString(t.Username) > 50 {
          return ErrUsernameMaxLengthValidation
      }

      return nil
  }
  ```

### `govalid:minlength`
- **Description**: Ensures that a string field's length is at least the specified minimum value (Unicode-aware).
- **Example**:
  ```go
  type User struct {
      // +govalid:minlength=3
      Username string `json:"username"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateUser(t *User) error {
      if t == nil {
          return ErrNilUser
      }

      if utf8.RuneCountInString(t.Username) < 3 {
          return ErrUsernameMinLengthValidation
      }

      return nil
  }
  ```

### `govalid:maxitems`
- **Description**: Ensures that a collection field's length does not exceed the specified maximum number of items. Supports slice, array, map, and channel types.
- **Example**:
  ```go
  type Collection struct {
      // +govalid:maxitems=10
      Items []string `json:"items"`
      
      // +govalid:maxitems=5
      Metadata map[string]string `json:"metadata"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateCollection(t *Collection) error {
      if t == nil {
          return ErrNilCollection
      }

      if len(t.Items) > 10 {
          return ErrItemsMaxItemsValidation
      }

      if len(t.Metadata) > 5 {
          return ErrMetadataMaxItemsValidation
      }

      return nil
  }
  ```

### `govalid:minitems`
- **Description**: Ensures that a collection field's length is at least the specified minimum number of items. Supports slice, array, map, and channel types.
- **Example**:
  ```go
  type Collection struct {
      // +govalid:minitems=1
      Items []string `json:"items"`
      
      // +govalid:minitems=2
      Tags []string `json:"tags"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateCollection(t *Collection) error {
      if t == nil {
          return ErrNilCollection
      }

      if len(t.Items) < 1 {
          return ErrItemsMinItemsValidation
      }

      if len(t.Tags) < 2 {
          return ErrTagsMinItemsValidation
      }

      return nil
  }
  ```


## License

govalid is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
