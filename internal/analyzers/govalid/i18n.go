package govalid

import (
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/language"

	"github.com/sivchari/govalid/internal/validator"
)

// ensureI18n loads the i18n configuration from the -i18n flag on first use.
// It is a no-op when -i18n is not set or the configuration is already loaded.
func (g *generator) ensureI18n() error {
	if i18nPath == "" || g.i18n != nil {
		return nil
	}

	cfg, err := loadI18nConfig(i18nPath)
	if err != nil {
		return err
	}

	g.i18n = cfg

	return nil
}

// i18nConfig holds localized message templates keyed by language and rule type.
// It is populated from the YAML file passed via the -i18n flag.
type i18nConfig struct {
	// messages maps language -> rule type (e.g. "gt") -> message template.
	messages map[language.Tag]map[string]string
}

// loadI18nConfig reads and parses the YAML i18n configuration at path. The file
// maps a BCP-47 language tag to a map of rule type to message template:
//
//	ja:
//	  required: "{Field}は必須です"
//	  gt:       "{Field}は{Param}より大きい値にしてください"
func loadI18nConfig(path string) (*i18nConfig, error) {
	data, err := os.ReadFile(path) //nolint:gosec // path comes from a trusted CLI flag.
	if err != nil {
		return nil, fmt.Errorf("read i18n config: %w", err)
	}

	raw := map[string]map[string]string{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parse i18n config %s: %w", path, err)
	}

	cfg := &i18nConfig{messages: make(map[language.Tag]map[string]string, len(raw))}

	for langStr, msgs := range raw {
		tag, err := language.Parse(langStr)
		if err != nil {
			return nil, fmt.Errorf("invalid language %q in i18n config %s: %w", langStr, path, err)
		}

		cfg.messages[tag] = msgs
	}

	return cfg, nil
}

// build returns the localized messages for a single error declaration, with the
// {Field}, {Param} and {Path} placeholders substituted. The {Value} placeholder is
// left untouched for substitution at validation time. It returns nil when no
// language defines a template for the declaration's rule type.
func (c *i18nConfig) build(decl *validator.ErrDecl) map[language.Tag]string {
	field := fieldFromPath(decl.Path)
	replacer := strings.NewReplacer("{Field}", field, "{Param}", decl.Param, "{Path}", decl.Path)

	var out map[language.Tag]string

	for tag, msgs := range c.messages {
		tmpl, ok := msgs[decl.Type]
		if !ok {
			continue
		}

		if out == nil {
			out = make(map[language.Tag]string)
		}

		out[tag] = replacer.Replace(tmpl)
	}

	return out
}

// fieldFromPath returns the leaf field name from a dot-separated path,
// e.g. "User.Address.City" -> "City".
func fieldFromPath(path string) string {
	if i := strings.LastIndex(path, "."); i >= 0 {
		return path[i+1:]
	}

	return path
}
