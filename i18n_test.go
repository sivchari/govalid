package govalid_test

import (
	"context"
	"testing"

	"golang.org/x/text/language"

	"github.com/sivchari/govalid"
	govaliderrors "github.com/sivchari/govalid/validation/errors"
)

func TestLanguageFromContext(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ctx  context.Context
		want language.Tag
	}{
		{"default is english", context.Background(), language.English},
		{"set japanese", govalid.WithLanguage(context.Background(), language.Japanese), language.Japanese},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := govalid.LanguageFromContext(tt.ctx); got != tt.want {
				t.Errorf("LanguageFromContext() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestLocalize(t *testing.T) {
	t.Parallel()

	newErrs := func() govaliderrors.ValidationErrors {
		return govaliderrors.ValidationErrors{
			{
				Path:     "User.Name",
				Type:     "required",
				Reason:   "field Name is required",
				Messages: map[language.Tag]string{language.Japanese: "Nameは必須です"},
			},
		}
	}

	t.Run("japanese context localizes", func(t *testing.T) {
		t.Parallel()

		ctx := govalid.WithLanguage(context.Background(), language.Japanese)

		err := govalid.Localize(ctx, newErrs())
		if err.Error() != "field User.Name with value <nil> has failed validation required because Nameは必須です" {
			t.Errorf("unexpected localized error: %q", err.Error())
		}
	})

	t.Run("default context keeps english", func(t *testing.T) {
		t.Parallel()

		err := govalid.Localize(context.Background(), newErrs())
		if err.Error() != "field User.Name with value <nil> has failed validation required because field Name is required" {
			t.Errorf("unexpected default error: %q", err.Error())
		}
	})
}
