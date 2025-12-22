package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError returns a string suitable for AppError.Details
func FormatValidationError(err error) string {
	if err == nil {
		return ""
	}

	var sb strings.Builder

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := fe.Field()
			var msg string

			switch fe.Tag() {
			case "min":
				msg = fmt.Sprintf("must be at least %s characters", fe.Param())
			case "max":
				msg = fmt.Sprintf("must be at most %s characters", fe.Param())
			case "required":
				msg = "is required"
			case "valid_long_url":
				msg = "must be a valid HTTP or HTTPS URL"
			case "alias":
				msg = "can only contain letters, numbers, dashes, or underscores"
			default:
				msg = fmt.Sprintf("invalid value for '%s'", field)
			}

			sb.WriteString(fmt.Sprintf("%s: %s; ", field, msg))
		}

		// remove last semicolon and space
		details := strings.TrimSuffix(sb.String(), "; ")
		return details
	}

	// fallback
	return err.Error()
}
