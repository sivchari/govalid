package errors

import (
	"testing"

	"golang.org/x/text/language"
)

func TestValidationError_Localize(t *testing.T) {
	t.Parallel()

	base := ValidationError{
		Path:   "User.Age",
		Type:   "gt",
		Value:  10,
		Reason: "field Age must be greater than 18",
		Messages: map[language.Tag]string{
			language.Japanese: "Ageは18より大きい値にしてください",
			language.French:   "Age doit être supérieur à 18",
		},
	}

	tests := []struct {
		name string
		lang language.Tag
		want string
	}{
		{"exact japanese", language.Japanese, "Ageは18より大きい値にしてください"},
		{"exact french", language.French, "Age doit être supérieur à 18"},
		{"region falls back to base language", language.MustParse("ja-JP"), "Ageは18より大きい値にしてください"},
		{"unregistered language keeps default reason", language.German, "field Age must be greater than 18"},
		{"english keeps default reason", language.English, "field Age must be greater than 18"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := base.Localize(tt.lang).Reason
			if got != tt.want {
				t.Errorf("Localize(%s).Reason = %q, want %q", tt.lang, got, tt.want)
			}
		})
	}
}

func TestValidationError_Localize_Value(t *testing.T) {
	t.Parallel()

	err := ValidationError{
		Type:   "gt",
		Value:  10,
		Reason: "field Age must be greater than 18",
		Messages: map[language.Tag]string{
			language.Japanese: "Ageの値{Value}は18より大きくしてください",
		},
	}

	got := err.Localize(language.Japanese).Reason
	want := "Ageの値10は18より大きくしてください"

	if got != want {
		t.Errorf("Localize value substitution = %q, want %q", got, want)
	}
}

func TestValidationError_Localize_NoMessages(t *testing.T) {
	t.Parallel()

	err := ValidationError{Type: "required", Reason: "field Name is required"}

	if got := err.Localize(language.Japanese).Reason; got != "field Name is required" {
		t.Errorf("Localize without messages = %q, want default reason", got)
	}
}

func TestValidationErrors_Localize(t *testing.T) {
	t.Parallel()

	errs := ValidationErrors{
		{
			Type:     "required",
			Reason:   "field Name is required",
			Messages: map[language.Tag]string{language.Japanese: "Nameは必須です"},
		},
		{
			Type:   "email",
			Reason: "field Email is not valid email",
			// no Japanese message -> keeps default reason
		},
	}

	got := errs.Localize(language.Japanese)

	if got[0].Reason != "Nameは必須です" {
		t.Errorf("errs[0].Reason = %q, want localized", got[0].Reason)
	}

	if got[1].Reason != "field Email is not valid email" {
		t.Errorf("errs[1].Reason = %q, want default", got[1].Reason)
	}
}
