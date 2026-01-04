package validators

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var AliasRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
var Validate *validator.Validate

func ValidAlias(alias string) bool {
	return AliasRegex.MatchString(alias)
}

// ValidLongURL checks if a string is a valid long URL
func ValidLongURL(longUrl string) bool {
	u, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	if u.Host == "" {
		return false
	}
	parts := strings.Split(u.Hostname(), ".")
	tld := parts[len(parts)-1]
	return len(tld) >= 2
}

func RegisterValidators() {
	Validate = validator.New()

	// Set the tag name to "validate" (default, optional)
	Validate.SetTagName("validate")

	// Register custom validators
	Validate.RegisterValidation("alias", func(fl validator.FieldLevel) bool {
		return ValidAlias(fl.Field().String())
	})

	Validate.RegisterValidation("valid_long_url", func(fl validator.FieldLevel) bool {
		return ValidLongURL(fl.Field().String())
	})
}
