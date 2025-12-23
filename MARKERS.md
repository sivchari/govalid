# Supported Markers

govalid supports the following markers:

## `govalid:required`
- **Description**: Ensures that the field is not empty or nil.
- **Example**:
  ```go
  //govalid:required
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

## `govalid:lt`
- **Description**: Ensures that a numeric field is less than a specified value.
- **Example**:
  ```go
  //govalid:lt=18
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

## `govalid:lte`
- **Description**: Ensures that a numeric field is less than or equal to a specified value.
- **Example**:
  ```go
  //govalid:lte=65
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

      if !(t.Age <= 65) {
          return ErrAgeLTEValidation
      }

      return nil
  }
  ```

## `govalid:gt`
- **Description**: Ensures that a numeric field is greater than a specified value.
- **Example**:
  ```go
  //govalid:gt=100
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

## `govalid:gte`
- **Description**: Ensures that a numeric field is greater than or equal to a specified value.
- **Example**:
  ```go
  //govalid:gte=18
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

      if !(t.Age >= 18) {
          return ErrAgeGTEValidation
      }

      return nil
  }
  ```

## `govalid:maxlength`
- **Description**: Ensures that a string field's length does not exceed the specified maximum value (Unicode-aware).
- **Example**:
  ```go
  type User struct {
      //govalid:maxlength=50
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

## `govalid:minlength`
- **Description**: Ensures that a string field's length is at least the specified minimum value (Unicode-aware).
- **Example**:
  ```go
  type User struct {
      //govalid:minlength=3
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

## `govalid:length`
- **Description**: Ensures that a string field has exactly the specified length (Unicode-aware).
- **Example**:
  ```go
  type User struct {
      //govalid:length=7
      Name string `json:"name"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateUser(t *User) error {
      if t == nil {
          return ErrNilUser
      }

      if utf8.RuneCountInString(t.Name) != 7 {
          return ErrUserNameLengthValidation
      }

      return nil
  }
  ```
- **Note**: Uses `utf8.RuneCountInString()` for proper Unicode character counting, ensuring accurate validation for international characters and emojis.

## `govalid:maxitems`
- **Description**: Ensures that a collection field's length does not exceed the specified maximum number of items. Supports slice, array, map, and channel types.
- **Example**:
  ```go
  type Collection struct {
      //govalid:maxitems=10
      Items []string `json:"items"`
      
      //govalid:maxitems=5
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

## `govalid:minitems`
- **Description**: Ensures that a collection field's length is at least the specified minimum number of items. Supports slice, array, map, and channel types.
- **Example**:
  ```go
  type Collection struct {
      //govalid:minitems=1
      Items []string `json:"items"`
      
      //govalid:minitems=2
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

## `govalid:enum`
- **Description**: Ensures that a field value is within a specified set of allowed values. Supports string, numeric, and custom types with comparable values. Values should be comma-separated.
- **Example**:
  ```go
  // Custom types
  type UserRole string
  type Priority int

  type User struct {
      // String enum
      //govalid:enum=admin,user,guest
      Role string `json:"role"`
      
      // Numeric enum
      //govalid:enum=1,2,3
      Level int `json:"level"`
      
      // Custom string type enum
      //govalid:enum=manager,developer,tester
      UserRole UserRole `json:"user_role"`
      
      // Custom numeric type enum
      //govalid:enum=10,20,30
      Priority Priority `json:"priority"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateUser(t *User) error {
      if t == nil {
          return ErrNilUser
      }

      if t.Role != "admin" && t.Role != "user" && t.Role != "guest" {
          return ErrRoleEnumValidation
      }

      if t.Level != 1 && t.Level != 2 && t.Level != 3 {
          return ErrLevelEnumValidation
      }

      if t.UserRole != "manager" && t.UserRole != "developer" && t.UserRole != "tester" {
          return ErrUserRoleEnumValidation
      }

      if t.Priority != 10 && t.Priority != 20 && t.Priority != 30 {
          return ErrPriorityEnumValidation
      }

      return nil
  }
  ```

## `govalid:email`
- **Description**: Ensures that a string field is a valid email address using HTML5-compliant validation.
- **Example**:
  ```go
  type User struct {
      //govalid:email
      Email string `validate:"email" json:"email"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateUser(t *User) error {
      if t == nil {
          return ErrNilUser
      }

      if !emailRegex.MatchString(t.Email) {
          return ErrEmailEmailValidation
      }

      return nil
  }
  ```

## `govalid:url`
- **Description**: Ensures that a string field is a valid URL using HTTP/HTTPS protocol validation.
- **Example**:
  ```go
  type Resource struct {
      //govalid:url
      URL string `validate:"url" json:"url"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateResource(t *Resource) error {
      if t == nil {
          return ErrNilResource
      }

      if !validationhelper.IsValidURL(t.URL) {
          return ErrURLURLValidation
      }

      return nil
  }
  ```

## `govalid:uuid`
- **Description**: Ensures that a string field is a valid UUID following RFC 4122 format.
- **Example**:
  ```go
  type Resource struct {
      //govalid:uuid
      ID string `json:"id"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateResource(t *Resource) error {
      if t == nil {
          return ErrNilResource
      }

      if !isValidUUID(t.ID) {
          return ErrIDUUIDValidation
      }

      return nil
  }
  ```

## `govalid:cel`
- **Description**: Validates fields using Google's Common Expression Language (CEL) for complex validation logic.
- **Available Variables**: 
  - `value`: The current field value being validated
- **Example**:
  ```go
  type Config struct {
      // Numeric validation
      //govalid:cel=value > 0.0 && value <= 100.0
      Score float64 `json:"score"`
      
      // String length validation
      //govalid:cel=size(value) >= 3 && size(value) <= 50
      Username string `json:"username"`
      
      // Pattern matching
      //govalid:cel=value.contains('@')
      Email string `json:"email"`
      
      // List validation
      //govalid:cel=size(value) > 0 && size(value) <= 10
      Tags []string `json:"tags"`
  }
  ```
- **Generated Code**:
  ```go
  func ValidateConfig(t *Config) error {
      if t == nil {
          return ErrNilConfig
      }

      if !validationhelper.IsValidCEL("value > 0.0 && value <= 100.0", t.Score, t) {
          return ErrScoreCELValidation
      }

      if !validationhelper.IsValidCEL("size(value) >= 3 && size(value) <= 50", t.Username, t) {
          return ErrUsernameCELValidation
      }

      if !validationhelper.IsValidCEL("value.contains('@')", t.Email, t) {
          return ErrEmailCELValidation
      }

      if !validationhelper.IsValidCEL("size(value) > 0 && size(value) <= 10", t.Tags, t) {
          return ErrTagsCELValidation
      }

      return nil
  }
  ```
- **Note**: CEL validation follows govalid's zero-reflection philosophy. Cross-field validation (accessing other struct fields) is not supported.

## `govalid:alpha`
- **Description**: Ensures that a string field is alphabetical, i.e. all its characters belong to the english alphabet.
- **Example**:
  ```go
    type User struct {
        //govalid:alpha
        FirstName string `json:"first_name"`
    }
    ```
- **Generated Code**:
  ```go
  func ValidateUser(t *User) error {
    if t == nil {
        return ErrNilUser
    }

    if !validationhelper.IsValidAlpha(t.FirstName) {
        return ErrUserFirstNameAlphaValidation
    }

    return nil
  }
  ```

## `govalid:numeric`

- **Description**: Ensures that a string field contains only digit characters (`0`–`9`). Leading zeros are allowed. Negative signs, decimal points, or scientific notation are **not** accepted.

- **Example**:

  ```go
  type Payload struct {
      //govalid:numeric
      Phone string `json:"phone"`
  }
  ```

- **Generated Code**:

  ```go
  func ValidatePayload(t *Payload) error {
      if t == nil {
          return ErrNilPayload
      }

      if !validationhelper.IsNumeric(t.Phone) {
          return ErrPayloadPhoneNumericValidation
      }

      return nil
  }
  ```

## `govalid:ipv4`

- **Description**: Ensure that a string field is a valid RFC 791-compliant IPv4 address.

- **Example**:

  ```go
  type Request struct {
      //govalid:ipv4
      IP string `json:"ip"`
  }
  ```

- **Generated Code**:

  ```go
  func ValidateRequest(t *Request) error {
      if t == nil {
          return ErrNilRequest
      }

      var errs govaliderrors.ValidationErrors

      if ip := net.ParseIP(t.IP); ip == nil || ip.To4() == nil {
          err := ErrRequestIPIpv4Validation
          err.Value = t.IP
          errs = append(errs, err)
      }

      if len(errs) > 0 {
          return errs
      }
      return nil
  }

  ```

## `govalid:ipv6`

- **Description**: Ensure that a string field is a valid RFC 4291-compliant IPv6 address.

- **Example**:

  ```go
  type Request struct {
      //govalid:ipv6
      IP string `json:"ip"`
  }
  ```

- **Generated Code**:

  ```go
  func ValidateRequest(t *Request) error {
      if t == nil {
          return ErrNilRequest
      }

      var errs govaliderrors.ValidationErrors

      if ip := net.ParseIP(t.IP); ip == nil || ip.To4() != nil {
          err := ErrRequestIPIpv6Validation
          err.Value = t.IP
          errs = append(errs, err)
      }

      if len(errs) > 0 {
          return errs
      }
      return nil
  }

  ```