package govalid

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/text/language"

	"github.com/sivchari/govalid/internal/validator"
)

func writeI18nFile(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "i18n.yaml")

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write temp i18n file: %v", err)
	}

	return path
}

func TestLoadI18nConfig(t *testing.T) {
	t.Parallel()

	path := writeI18nFile(t, "ja:\n  required: \"{Field}は必須です\"\n  gt: \"{Field}は{Param}より大きい値\"\nfr:\n  required: \"{Field} est obligatoire\"\n")

	cfg, err := loadI18nConfig(path)
	if err != nil {
		t.Fatalf("loadI18nConfig: %v", err)
	}

	if got := cfg.messages[language.Japanese]["required"]; got != "{Field}は必須です" {
		t.Errorf("ja/required = %q", got)
	}

	if got := cfg.messages[language.French]["required"]; got != "{Field} est obligatoire" {
		t.Errorf("fr/required = %q", got)
	}
}

func TestLoadI18nConfig_Errors(t *testing.T) {
	t.Parallel()

	t.Run("missing file", func(t *testing.T) {
		t.Parallel()

		if _, err := loadI18nConfig(filepath.Join(t.TempDir(), "nope.yaml")); err == nil {
			t.Error("expected error for missing file")
		}
	})

	t.Run("invalid yaml", func(t *testing.T) {
		t.Parallel()

		path := writeI18nFile(t, "ja: [unbalanced")
		if _, err := loadI18nConfig(path); err == nil {
			t.Error("expected error for invalid yaml")
		}
	})

	t.Run("invalid language", func(t *testing.T) {
		t.Parallel()

		path := writeI18nFile(t, "\"!!!\":\n  required: \"x\"\n")
		if _, err := loadI18nConfig(path); err == nil {
			t.Error("expected error for invalid language tag")
		}
	})
}

func TestI18nConfig_Build(t *testing.T) {
	t.Parallel()

	cfg := &i18nConfig{
		messages: map[language.Tag]map[string]string{
			language.Japanese: {"gt": "{Field}は{Param}より大きく{Value}は不正", "required": "{Field}は必須"},
			language.French:   {"gt": "{Field} > {Param}"},
		},
	}

	t.Run("substitutes field and param, keeps value token", func(t *testing.T) {
		t.Parallel()

		decl := validator.ErrDecl{Path: "User.Age", Type: "gt", Param: "18"}

		got := cfg.build(&decl)
		if got[language.Japanese] != "Ageは18より大きく{Value}は不正" {
			t.Errorf("ja = %q", got[language.Japanese])
		}

		if got[language.French] != "Age > 18" {
			t.Errorf("fr = %q", got[language.French])
		}
	})

	t.Run("returns nil when no language has the rule", func(t *testing.T) {
		t.Parallel()

		decl := validator.ErrDecl{Path: "User.Name", Type: "email"}
		if got := cfg.build(&decl); got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
}

func TestFieldFromPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path string
		want string
	}{
		{"User.Age", "Age"},
		{"User.Address.City", "City"},
		{"Name", "Name"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			t.Parallel()

			if got := fieldFromPath(tt.path); got != tt.want {
				t.Errorf("fieldFromPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}

func TestSortedMessages(t *testing.T) {
	t.Parallel()

	got := sortedMessages(map[language.Tag]string{
		language.Japanese: "ja",
		language.French:   "fr",
		language.English:  "en",
	})

	want := []langMessage{
		{Lang: "en", Msg: "en"},
		{Lang: "fr", Msg: "fr"},
		{Lang: "ja", Msg: "ja"},
	}

	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("sortedMessages[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestHasLocalizedMessages(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		decls []validator.ErrDecl
		want  bool
	}{
		{"none", []validator.ErrDecl{{Type: "required"}}, false},
		{"some", []validator.ErrDecl{{Type: "required"}, {Type: "gt", Messages: map[language.Tag]string{language.Japanese: "x"}}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := hasLocalizedMessages(tt.decls); got != tt.want {
				t.Errorf("hasLocalizedMessages = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestEnsureI18n exercises the -i18n flag loading path. It mutates the package-level
// i18nPath and is therefore not parallel.
func TestEnsureI18n(t *testing.T) {
	original := i18nPath

	t.Cleanup(func() { i18nPath = original })

	t.Run("no flag is a no-op", func(t *testing.T) {
		i18nPath = ""

		g := &generator{}
		if err := g.ensureI18n(); err != nil || g.i18n != nil {
			t.Errorf("ensureI18n() err=%v i18n=%v, want no-op", err, g.i18n)
		}
	})

	t.Run("loads config from flag", func(t *testing.T) {
		i18nPath = writeI18nFile(t, "ja:\n  required: \"{Field}は必須です\"\n")

		g := &generator{}
		if err := g.ensureI18n(); err != nil {
			t.Fatalf("ensureI18n() err=%v", err)
		}

		if g.i18n == nil || g.i18n.messages[language.Japanese]["required"] != "{Field}は必須です" {
			t.Errorf("config not loaded: %+v", g.i18n)
		}
	})

	t.Run("propagates load error", func(t *testing.T) {
		i18nPath = filepath.Join(t.TempDir(), "missing.yaml")

		g := &generator{}
		if err := g.ensureI18n(); err == nil {
			t.Error("expected error for missing config file")
		}
	})
}
