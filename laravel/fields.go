package laravel

import "fmt"

var fieldTypes = []string{
	"text",
	"file",
	"number",
	"select",
	"ckeditor",
	"system",
	"datetime|system",
}

func Fields() {
	// 输出字段类型
	for _, fieldType := range fieldTypes {
		fmt.Println(fieldType)
	}
}
