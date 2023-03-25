package utils

import "regexp"

type Authentication interface {
	ValidatePassword(password string) bool
	ValidateUsername(username string) bool
}

type Auth struct{}

func (a *Auth) ValidatePassword(password string) bool {
	// one digit from 1 to 9, one lowercase letter, one uppercase letter, one special character, no space, and it must be 8-16 characters long
	regex := "/^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*\\W)(?!.* ).{8,16}$/"

	ok, _ := regexp.MatchString(regex, password)

	return ok
}

func (a *Auth) ValidateUsername(username string) bool {
	// https://stackoverflow.com/questions/12018245/regular-expression-to-validate-username
	regex := "^(?=.{8,20}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$"

	ok, _ := regexp.MatchString(regex, username)

	return ok
}
