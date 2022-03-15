package main

import (
	"fmt"
	"go-package/regexp/example"
	"regexp"
)

var (
	pattern = "[a-zA-Z0-9]"
	s       = "abc bac acc afc"
	sByte   = []byte{2, 0}
	expr     = "a.c"
	repl     = "repl"

)

func main() {
	// 一、正则匹配
	// 1.string 匹配
	regString, _ := regexp.MatchString(pattern, s)
	fmt.Println(regString)//true
	// 2.[]byte 匹配
	regByte, _ := regexp.Match(pattern, sByte)
	fmt.Println(regByte) //flase

	// 二、预编译匹配
	// Compile
	// 1.最左优先匹配，返回结果包含err字段
	regCompile, _ := regexp.Compile(expr)
	fmt.Println(example.MatchString(regCompile))//true
	// 2.同Compile；但简化错误返回，返回结果没有err字段，报错抛出panic
	regMustCompile := regexp.MustCompile(expr)
	fmt.Println(example.MatchString(regMustCompile))//true
	// CompilePOSIX
	// 3.POSIX ERE (egrep) 语法;最左最长匹配；返回结果包含err字段
	regCompilePOSIX, _ := regexp.CompilePOSIX(expr)
	fmt.Println(example.MatchString(regCompilePOSIX))//true
	// 4.同CompilePOSIX；简化错误返回
	regMustCompilePOSIX := regexp.MustCompilePOSIX(expr)
	fmt.Println(example.MatchString(regMustCompilePOSIX))//true

	// 三、正则查找--使用regexp.MustCompile方法做示例
	// todo String
	// 1.非贪婪；最左匹配第一个字符串
	findString := regMustCompile.FindString(s)
	fmt.Println(findString)//abc
	// 2.贪婪；n表示最多查询个数，为负数时查所有
	findAllString := regMustCompile.FindAllString(s, -1)
	fmt.Println(findAllString)//[abc acc afc]
	// Submatch
	// 3.非贪婪；分组、最左匹配第一个字符串，结果返回[]string数组
	findStringSubmatch := regMustCompile.FindStringSubmatch(s)
	fmt.Println(findStringSubmatch)//[abc]
	// 4.贪婪；分组、返回所有匹配的字符串，结果返回[][]string
	findAllStringSubmatch := regMustCompile.FindAllStringSubmatch(s, -1)
	fmt.Println(findAllStringSubmatch)//[[abc] [acc] [afc]]
	// Index
	// 5.非贪婪；最左匹配第一个字符串Index位置--slice格式
	findStringIndex := regMustCompile.FindStringIndex(s)
	fmt.Println(findStringIndex)//[0 3]
	// 6.贪婪；返回所有匹配字符串位置
	findAllStringIndex := regMustCompile.FindAllStringIndex(s, -1)
	fmt.Println(findAllStringIndex)//[[0 3] [8 11] [12 15]]
	// 7.非贪婪；分组、最左匹配第一个字符串Index位置--slice格式
	findStringSubmatchIndex := regMustCompile.FindStringSubmatchIndex(s)
	fmt.Println(findStringSubmatchIndex)//[[0 3] [8 11] [12 15]]
	// 8.贪婪；分组、返回所有匹配字符串位置
	findAllStringSubmatchIndex := regMustCompile.FindAllStringSubmatchIndex(s, -1)
	fmt.Println(findAllStringSubmatchIndex)//[[0 3] [8 11] [12 15]]
	// todo byte --函数方法基本同String
	// 1.非贪婪；最左匹配第一个字节
	find := regMustCompile.Find(sByte)
	fmt.Println(find)//[]

	// 四、正则替换--简述
	// todo String
	// 1.将符合正则表达式的字符串替换成repl, 在替换时，repl中的'$'($后接数字或字符)符号会按照Expand方法的规则进行解释和替换
	replace := regMustCompile.ReplaceAllString(s, repl)
	fmt.Println(replace)// repl bac repl repl
	// 2.将符合正则表达式的字符串替换成repl, 不会使用Expand进行扩展。
	replaceLiteral := regMustCompile.ReplaceAllLiteralString(s, repl)
	fmt.Println(replaceLiteral)// repl bac repl repl
	// 3.使用函数指定替换规则, 不会使用Expand进行扩展。
	replaceFunc := regMustCompile.ReplaceAllStringFunc(s, example.ReplFunc)
	fmt.Println(replaceFunc) //replFunc bac acc afc
	// todo byte --函数方法基本同String

}
