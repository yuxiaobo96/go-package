package example

import "strings"

func ReplFunc(s string) string {
	if strings.Contains(s, "abc") {
		return "replFunc"
	}
	return s
}