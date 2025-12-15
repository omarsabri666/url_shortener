package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// ValidAlias checks if a string contains only letters, numbers, dash, or underscore
func ValidAlias(alias string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return re.MatchString(alias)
}

// ValidLongURL checks if a string is a valid long URL
func ValidLongURL(longUrl string) bool {
	re := regexp.MustCompile(`^(http(s)?:\/\/)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)
	return re.MatchString(longUrl)
}
func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// regex: allow letters, numbers, dash, underscore
		v.RegisterValidation("alias", func(fl validator.FieldLevel) bool {
			alias := fl.Field().String()
			return ValidAlias(alias)

			// re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
			// return re.MatchString(alias)
		})
		v.RegisterValidation("valid_long_url", func(f1 validator.FieldLevel) bool {
			longUrl := f1.Field().String()
			return ValidLongURL(longUrl)

			// re := regexp.MustCompile(`^(http(s)?:\/\/)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)

			// return re.MatchString(longUrl)
		})

	}
}
