package utils

import "regexp"

func IsValidPassword(password string) bool {
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	if len(password) < LEN_MIN_PASSWORD {
		return false
	}
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case (char >= '!' && char <= '/') || (char >= ':' && char <= '@') || (char >= '[' && char <= '`') || (char >= '{' && char <= '~'):
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
