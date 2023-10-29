package utils

import "strings"

func ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 去除結尾s
func RemoveS(str string) string {
	inputString := str

	// 检查字符串是否以 "s" 结尾
	if len(inputString) > 0 && inputString[len(inputString)-1] == 's' {
		inputString = inputString[:len(inputString)-1]
	}
	return inputString
}
