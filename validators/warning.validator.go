package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func IsWarningText(fl validator.FieldLevel) bool {
	return !strings.Contains(fl.Field().String(), "warning")
}
