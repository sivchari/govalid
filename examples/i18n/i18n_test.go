package i18n

import (
	"context"
	"testing"

	"golang.org/x/text/language"

	"github.com/sivchari/govalid"
)

func TestUserValidation_Localized(t *testing.T) {
	t.Parallel()

	invalid := &User{Name: "", Age: 10, Nickname: "ok"}

	tests := []struct {
		name string
		lang language.Tag
		want string
	}{
		{
			name: "japanese",
			lang: language.Japanese,
			want: "field User.Name with value  has failed validation required because Nameは必須です\n" +
				"field User.Age with value 10 has failed validation gt because Ageは18より大きい値にしてください",
		},
		{
			name: "french",
			lang: language.French,
			want: "field User.Name with value  has failed validation required because Name est obligatoire\n" +
				"field User.Age with value 10 has failed validation gt because Age doit être supérieur à 18",
		},
		{
			name: "default english",
			lang: language.English,
			want: "field User.Name with value  has failed validation required because field Name is required\n" +
				"field User.Age with value 10 has failed validation gt because field Age must be greater than 18",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := govalid.WithLanguage(context.Background(), tt.lang)

			err := ValidateUserContext(ctx, invalid)
			if err == nil {
				t.Fatal("expected validation error, got nil")
			}

			if err.Error() != tt.want {
				t.Errorf("localized error mismatch:\n got: %q\nwant: %q", err.Error(), tt.want)
			}
		})
	}
}

func TestUserValidation_Valid(t *testing.T) {
	t.Parallel()

	valid := &User{Name: "Alice", Age: 20, Nickname: "ali"}

	ctx := govalid.WithLanguage(context.Background(), language.Japanese)
	if err := ValidateUserContext(ctx, valid); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
