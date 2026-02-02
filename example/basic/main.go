// Package main demonstrates basic usage of govalid validation.
package main

import (
	"fmt"
	"os"

	govaliderrors "github.com/sivchari/govalid/validation/errors"
)

func main() {
	// Example 1: Valid user
	fmt.Println("=== Example 1: Valid User ===")
	validUser := &User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}
	if err := ValidateUser(validUser); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("Validation passed!")
	}

	// Example 2: Invalid user with multiple errors
	fmt.Println("\n=== Example 2: Invalid User ===")
	invalidUser := &User{
		Name:  "",              // required, minlength=1 violation
		Email: "invalid-email", // email format violation
		Age:   200,             // lte=150 violation
	}
	if err := ValidateUser(invalidUser); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	}

	// Example 3: Loop through all validation errors
	fmt.Println("\n=== Example 3: Loop Through Errors ===")
	if err := ValidateUser(invalidUser); err != nil {
		// Type assert to ValidationErrors to iterate over individual errors
		if errs, ok := err.(govaliderrors.ValidationErrors); ok {
			fmt.Printf("Found %d validation errors:\n", len(errs))
			for i, e := range errs {
				fmt.Printf("  %d. Field: %s\n", i+1, e.Path)
				fmt.Printf("     Type: %s\n", e.Type)
				fmt.Printf("     Value: %v\n", e.Value)
				fmt.Printf("     Reason: %s\n", e.Reason)
			}
		}
	}

	// Example 4: Check for specific error type
	fmt.Println("\n=== Example 4: Check Specific Error ===")
	if err := ValidateUser(invalidUser); err != nil {
		if errs, ok := err.(govaliderrors.ValidationErrors); ok {
			for _, e := range errs {
				if e.Type == "required" {
					fmt.Printf("Required field missing: %s\n", e.Path)
				}
			}
		}
	}

	// Example 5: Using Validate() method (implements govalid.Validator interface)
	fmt.Println("\n=== Example 5: Using Validate() Method ===")
	if err := validUser.Validate(); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Validation passed using Validate() method!")
}
