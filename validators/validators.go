package validators

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var AliasRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
var Validate *validator.Validate

// ValidAlias checks if a string contains only letters, numbers, dash, or underscore
//
//	func ValidAlias(alias string) bool {
//		re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
//		return re.MatchString(alias)
//	}
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

// func RegisterValidators() {
// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		// regex: allow letters, numbers, dash, underscore
// 		v.RegisterValidation("alias", func(fl validator.FieldLevel) bool {
// 			alias := fl.Field().String()
// 			return ValidAlias(alias)

// 			// re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
// 			// return re.MatchString(alias)
// 		})
// 		v.RegisterValidation("valid_long_url", func(f1 validator.FieldLevel) bool {
// 			longUrl := f1.Field().String()
// 			return ValidLongURL(longUrl)

// 			// re := regexp.MustCompile(`^(http(s)?:\/\/)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)

// 			// return re.MatchString(longUrl)
// 		})

// 	}
// }

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
