package utils

import (
	"regexp"
	"unicode"
)

func IsValidPassword(password string) bool {
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	if len(password) < LEN_MIN_PASSWORD {
		return false
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		default:
			return false
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func IsValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailPattern, email)
	return matched
}
