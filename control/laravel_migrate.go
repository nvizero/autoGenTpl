package control

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
	db "tpl/db/sqlc"
	"tpl/utils"
)

var dbpg = db.ConnDev()

func GetTbGenerateMigrateTable(projectid int32, projectName string) {
	arg := db.WhereTbByPIDParams{
		Limit:     10,
		Offset:    0,
		ProjectID: sql.NullInt32{Int32: projectid, Valid: true},
	}
	tables, _ := dbpg.WhereTbByPID(ctx, arg)
	for _, row := range tables {
		MigrationTable(row.Name.String, projectName, row.ID)
	}
}

func MigrationTable(rep_model, projectName string, tableid int32) {
	var txt string
	directory := localhostDir + "/" + projectName + DatabaseDir
	arg := sql.NullInt32{Int32: tableid, Valid: true}
	tbs, _ := dbpg.GetTFBytID(ctx, arg)
	fmt.Println(tbs)
	for _, row := range tbs {
		txt += fmt.Sprintf("            $table->%s(\"%s\");\n", row.LaravelMap.String, row.FieldName.String)
	}
	// 不区分大小写替换 "xnews" 为 "abc"
	re := regexp.MustCompile(`(?i)xnews`)
	migration_head = re.ReplaceAllString(migration_head, rep_model)
	migration_end = re.ReplaceAllString(migration_end, rep_model)
	// 获取当前时间
	currentTime := time.Now()

	// 格式化时间为 "2006_01_02_150405"
	timeFormat := currentTime.Format("2006_01_02_150405")
	// 将替换后的字符串写入文件
	// 拼接文件名前缀和时间格式
	fileName := timeFormat + fmt.Sprintf("_create_%s.php", utils.FirstLower(rep_model))
	combinedMigration := migration_head + txt + migration_end
	err := ioutil.WriteFile(directory+fileName, []byte(combinedMigration), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n替换完成，并已将结果保存到文件 %s_migration.php", utils.FirstLower(rep_model))
}
