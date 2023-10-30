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

var (
	dbpg           = db.ConnDev()
	router_web_txt string
	sidemenu_txt   string
	combinedTxt    string
)

func GetTbGenerateMigrateTable(projectid int32, projectName string) {
	arg := db.WhereTbByPIDParams{
		Limit:     10,
		Offset:    0,
		ProjectID: sql.NullInt32{Int32: projectid, Valid: true},
	}
	tables, _ := dbpg.WhereTbByPID(ctx, arg)
	for _, row := range tables {
		//MigrationTable(row.Name.String, projectName, row.ID)
		//CreateModel(row.Name.String, projectName, row.ID)
		//CollectRouter(row.Name.String)
		CollectSideMenu(row.Name.String)
	}
	//CreateRouter(projectName)
	CreateSideMenu(projectName)
}

// 收集ControllerName建立sidemenu
func CollectSideMenu(tablename string) {
	router_web_txt += fmt.Sprintf("                    <li><a href=\"{{ route('%s.index') }}\">%s</a></li>\n", tablename, tablename)
}

// 收集ControllerName建立router
func CollectRouter(tablename string) {
	firstname := utils.FirstLower(tablename)
	ControllerName := utils.RemoveS(utils.FirstUpper(tablename))
	router_web_txt += fmt.Sprintf("    Route::resource('%s', '%sController');\n", firstname, ControllerName)
}

// 建立 sidemenu.blade.php
func CreateSideMenu(projectName string) {
	directory := localhostDir + "/" + projectName + SideMenuDir
	combinedTxt = sidemenu_head + router_web_txt + sidemenu_footer
	saveName := fmt.Sprintf("sidemenu.blade.php")
	err := ioutil.WriteFile(directory+saveName, []byte(combinedTxt), 0644)
	ChkErr(err)
}

// 建立/router/web.php
func CreateRouter(projectName string) {
	directory := localhostDir + "/" + projectName + RouterDir
	combinedTxt = router_head + router_web_txt + router_footer
	saveName := fmt.Sprintf("web.php")
	err := ioutil.WriteFile(directory+saveName, []byte(combinedTxt), 0644)
	ChkErr(err)
}

// 建立Model .php
func CreateModel(tablename, projectName string, tableid int32) {
	var fields_txt string
	var field_setting string
	directory := localhostDir + "/" + projectName + ModelDir
	fileName := utils.RemoveS(utils.FirstUpper(tablename))
	className := fmt.Sprintf("class %s extends BaseModel", fileName)
	table := fmt.Sprintf("  protected $table = '%s';", utils.FirstLower(tablename))
	combinedTxt = model_1 + className + model_2 + table
	arg := sql.NullInt32{Int32: tableid, Valid: true}
	tbs, _ := dbpg.GetTFBytID(ctx, arg)
	for _, row := range tbs {
		fields_txt += fmt.Sprintf("'%s',", row.FieldName.String)
		field_setting += fmt.Sprintf("          '%s' => [\n", row.FieldName.String)
		field_setting += fmt.Sprintf("               'type' => '%s',\n", row.LaravelMap.String)
		field_setting += fmt.Sprintf("               'required' => true,\n")
		field_setting += fmt.Sprintf("               'search' => [\n")
		field_setting += fmt.Sprintf("                    'level' => 'like',\n")
		field_setting += fmt.Sprintf("               ]\n")
		field_setting += fmt.Sprintf("          ],\n")
	}
	combinedTxt += model_3 + "       " + fields_txt + model_4
	combinedTxt += model_5 + field_setting + model_6
	saveName := fmt.Sprintf("%s.php", fileName)
	err := ioutil.WriteFile(directory+saveName, []byte(combinedTxt), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 建立migration create table
func MigrationTable(rep_model, projectName string, tableid int32) {
	var txt string
	directory := localhostDir + "/" + projectName + DatabaseDir
	arg := sql.NullInt32{Int32: tableid, Valid: true}
	tbs, _ := dbpg.GetTFBytID(ctx, arg)
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
	// 将替换后的字符串写入文件  拼接文件名前缀和时间格式
	fileName := timeFormat + fmt.Sprintf("_create_%s.php", utils.FirstLower(rep_model))
	combinedMigration := migration_head + txt + migration_end
	err := ioutil.WriteFile(directory+fileName, []byte(combinedMigration), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
