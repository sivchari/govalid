package validator_test

import (
	"testing"

	"github.com/sivchari/govalid/internal/validator"
)

func TestNewFieldPath(t *testing.T) {
	tests := []struct {
		name       string
		components []string
		want       validator.FieldPath
	}{
		{
			name:       "simple",
			components: []string{"User", "Name"},
			want:       validator.FieldPath("User.Name"),
		},
		{
			name:       "with empty string in middle",
			components: []string{"User", "", "Name"},
			want:       validator.FieldPath("User.Name"),
		},
		{
			name:       "with empty string at start",
			components: []string{"", "User", "Name"},
			want:       validator.FieldPath("User.Name"),
		},
		{
			name:       "with empty string at end",
			components: []string{"User", "Name", ""},
			want:       validator.FieldPath("User.Name"),
		},
		{
			name:       "multiple empty strings",
			components: []string{"", "User", "", "Address", "", "City", ""},
			want:       validator.FieldPath("User.Address.City"),
		},
		{
			name:       "all empty strings",
			components: []string{"", "", ""},
			want:       validator.FieldPath(""),
		},
		{
			name:       "no components",
			components: []string{},
			want:       validator.FieldPath(""),
		},
		{
			name:       "with whitespace strings",
			components: []string{"User", "  ", "Name"},
			want:       validator.FieldPath("User.Name"),
		},
		{
			name:       "with tab and newline",
			components: []string{"User", "\t", "\n", "Name"},
			want:       validator.FieldPath("User.Name"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validator.NewFieldPath(tt.components...)
			if got != tt.want {
				t.Errorf("NewFieldPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldPath_CleanedPath(t *testing.T) {
	tests := []struct {
		name      string
		fieldPath validator.FieldPath
		want      string
	}{
		{
			name:      "simple",
			fieldPath: validator.FieldPath("User.Name"),
			want:      "UserName",
		},
		{
			name:      "single component",
			fieldPath: validator.FieldPath("User"),
			want:      "User",
		},
		{
			name:      "empty path",
			fieldPath: validator.FieldPath(""),
			want:      "",
		},
		{
			name:      "deeply nested",
			fieldPath: validator.FieldPath("Company.Department.Team.Member.Profile.Name"),
			want:      "CompanyDepartmentTeamMemberProfileName",
		},
		{
			name:      "with underscores",
			fieldPath: validator.FieldPath("User_Info.Home_Address.City_Name"),
			want:      "User_InfoHome_AddressCity_Name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fieldPath.CleanedPath()
			if got != tt.want {
				t.Errorf("FieldPath.CleanedPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldPath_String(t *testing.T) {
	tests := []struct {
		name      string
		fieldPath validator.FieldPath
		want      string
	}{
		{
			name:      "simple",
			fieldPath: validator.FieldPath("User.Name"),
			want:      "User.Name",
		},
		{
			name:      "empty path",
			fieldPath: validator.FieldPath(""),
			want:      "",
		},
		{
			name:      "single component",
			fieldPath: validator.FieldPath("User"),
			want:      "User",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fieldPath.String()
			if got != tt.want {
				t.Errorf("FieldPath.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
