package functions

import "strings"

func FrenchLowerCase(word string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(word), "é", "e"), "è", "e"), "ê", "e")
}
