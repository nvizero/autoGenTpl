package control

import (
	"fmt"
	"net/url"
	"strings"
)

type FieldsRow struct {
	FieldName string
	ModleType string
	ShowName  string
	Migration string
	IsRequire string
}

type TableInfo struct {
	TableName string
	Fields    []FieldsRow
}

func ParseTableFields(postData string) {
	values, err := url.ParseQuery(postData)
	ChkErr(err)
	dataMap := make(map[string]map[string]string)
	for key, values := range values {
		parts := strings.Split(key, "_")
		fmt.Println(key, values, parts)
		if len(parts) >= 3 {
			tableName := parts[0]
			fieldName := parts[1]
			fieldIndex := parts[len(parts)-1]
			if _, found := dataMap[tableName]; !found {
				dataMap[tableName] = make(map[string]string)
			}
			dataMap[tableName][fieldName+"_"+fieldIndex] = values[0]
		}
	}
}

func ParseTable(postData string) map[string]string {
	values, err := url.ParseQuery(postData)
	ChkErr(err)
	tables := make(map[string]string)
	for key, values := range values {
		parts := strings.Split(key, "_")
		if len(parts) == 2 {
			tables[key] = values[0]
		}
	}
	return tables
}
