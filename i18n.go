package govalid

import (
	"context"

	"golang.org/x/text/language"

	govaliderrors "github.com/sivchari/govalid/validation/errors"
)

// languageKey is the context key under which the validation language is stored.
type languageKey struct{}

// WithLanguage returns a copy of ctx that carries lang, which generated validators
// use to localize error messages. Use it together with code generated via the
// -i18n option; without localized messages the language has no effect.
func WithLanguage(ctx context.Context, lang language.Tag) context.Context {
	return context.WithValue(ctx, languageKey{}, lang)
}

// LanguageFromContext returns the language stored in ctx by WithLanguage.
// It returns language.English when no language is set.
func LanguageFromContext(ctx context.Context) language.Tag {
	if lang, ok := ctx.Value(languageKey{}).(language.Tag); ok {
		return lang
	}

	return language.English
}

// Localize localizes the validation errors using the language stored in ctx and
// returns them. It is called by generated validators when code is generated with
// the -i18n option. Errors without a localized message for the language keep their
// default (English) Reason.
func Localize(ctx context.Context, errs govaliderrors.ValidationErrors) error {
	return errs.Localize(LanguageFromContext(ctx))
}
