package control

import (
	"fmt"
	"strings"
	"testing"
)

// test list project
func TestPData(t *testing.T) {
	var postData string = "tablename_1=cat&table1_fieldname_1=test1&table1_showName_1=bb&table1_migration_1=aa&table1_modelType_1=ww&table1_isRequire_1=1&table1_fieldname_2=test2&table1_showName_2=qweqwe&table1_migration_2=qqasd&table1_modelType_2=kjkjk&table1_isRequire_2=1&tablename_2=dogs&table2_fieldname_1=qqq1&table2_showName_1=iui&table2_migration_1=iuiu&table2_modelType_1=iuiu&table2_isRequire_1=1&table2_fieldname_2=www2&table2_showName_2=iuiu&table2_migration_2=jhj&table2_modelType_2=jhjh&table2_isRequire_2=1&table2_fieldname_3=ddd3&table2_showName_3=jhjh&table2_migration_3=jhjh&table2_modelType_3=jh&table2_isRequire_3=1"
	//require.NoError(t, err)
	tables := ParseTable(postData)
	for i, data := range tables {
		fmt.Println(i, data)
		parts := strings.Split(i, "_")
		key := fmt.Sprintf("table%s", parts[1])
		// tables := ParseTableFields(postData)
		// 	fmt.Println(i, data)
		// 	// 	ParseRow(data)
		fmt.Println(key)
		// dd := ParseRow(data)
	}
}
