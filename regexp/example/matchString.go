package example

import "regexp"

// 判断是否匹配
func MatchString(reg *regexp.Regexp) bool {
	return reg.MatchString("abc bac acc")
}
