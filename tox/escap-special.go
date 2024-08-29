package tox

import "fmt"

func EscapingSpecial(es interface{}) string {

	s := fmt.Sprintf(`%v`, es)

	specials := []string{"[", "]", "^", "\\\\", ".", "?", "|", "$", "*", "(", ")", "-", "{", "}"}
	escaped := ``
	for _, char := range s {
		isSpecial := false
		for _, v := range specials {
			if string(char) == v {
				isSpecial = true
				break
			}
		}
		if isSpecial {
			escaped += "\\"
		}
		escaped += string(char)
	}
	return escaped
}
