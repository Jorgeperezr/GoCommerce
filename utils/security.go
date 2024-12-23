package utils

import "regexp"

func ValidarEntrada(input string) bool {
	regex := `^[a-zA-Z0-9_@.-]+$`
	match, _ := regexp.MatchString(regex, input)
	return match
}
