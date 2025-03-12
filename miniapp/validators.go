package miniapp

import (
	"errors"
	"regexp"
)

var shortNameRe = regexp.MustCompile(`^\w{3,30}$`)

func ValidateShortName(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("must be a string")
	}
	if !shortNameRe.MatchString(s) {
		return errors.New("invalid short name format")
	}
	return nil
}
