package miniapp

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var shortNameRe = regexp.MustCompile(`^\w{3,30}$`)

func ValidateShortName(fl validator.FieldLevel) bool {
	return shortNameRe.MatchString(fl.Field().String())
}
