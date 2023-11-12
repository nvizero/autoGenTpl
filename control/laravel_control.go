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
)

func RunMigration(projectName string) {
	statusChan <- "2建立 基本 database 執行php artisan migrate"
	params = []interface{}{projectName, cmd_sh[1]}
	utils.RunCmd(params, docker_run_sh, statusChan)
}

// 讀最後一筆並建立LARAVEL
func GenerateLaravelByLest() {
	data, err := dbpg.LatestOne(ctx)
	ChkErr(err)
	Project = Pj{
		ProjectName: data.Name.String,
		ProjectID:   data.ID,
		DockerPort:  data.Port.Int32,
	}
	GenLaravel(statusChan)
	GetTbGenerateMigrateTable(data.ID, data.Name.String)
}

func GetTbGenerateMigrateTable(projectid int32, projectName string) {
	arg := db.WhereTbByPIDParams{
		Limit:     10,
		Offset:    0,
		ProjectID: sql.NullInt32{Int32: projectid, Valid: true},
	}
	tables, _ := dbpg.WhereTbByPID(ctx, arg)
	for _, row := range tables {
		MigrationTable(row.Name.String, projectName, row.ID)
		CreateModel(row.Name.String, projectName, row.ID)
		CollectRouter(row.Name.String)
		CreateController(row.Name.String, projectName, row.ID)
		CollectSideMenu(row.Name.String)
	}
	CreateRouter(projectName)
	CreateSideMenu(projectName)
	RunMigration(projectName)
}

func CreateController(tablename, projectName string, tableid int32) {
	var controller_txt string
	directory := localhostDir + "/" + projectName + ControllerDir
	fileName := utils.RemoveS(utils.FirstUpper(tablename))
	comb1 := fmt.Sprintf("use App\\Models\\%s;\n", utils.RemoveS(utils.FirstUpper(tablename)))
	comb2 := fmt.Sprintf("class %sController extends TemplateController\n{\n", utils.RemoveS(utils.FirstUpper(tablename)))
	comb3 := fmt.Sprintf("\n    public string $main = '%s';\n", utils.FirstLower(tablename))
	comb4 := fmt.Sprintf("\n    function __construct(Request $request, %s $%s, RequestService $requestService)\n    {\n", utils.RemoveS(utils.FirstUpper(tablename)), utils.FirstLower(tablename))
	comb5 := fmt.Sprintf("\n        $this->entity = $%s;", utils.FirstLower(tablename))
	controller_txt = controller_0 + comb1 + comb2
	controller_txt += comb3 + comb4 + comb5 + controller_1
	saveName := fmt.Sprintf("%sController.php", fileName)
	err := ioutil.WriteFile(directory+saveName, []byte(controller_txt), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}

// 收集ControllerName建立sidemenu
func CollectSideMenu(tablename string) {
	sidemenu_txt += fmt.Sprintf("                    <li><a href=\"{{ route('%s.index') }}\">%s</a></li>\n", tablename, tablename)
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
	sidemenu := sidemenu_head + sidemenu_txt + sidemenu_footer
	saveName := fmt.Sprintf("sidemenu.blade.php")
	err := ioutil.WriteFile(directory+saveName, []byte(sidemenu), 0644)
	ChkErr(err)
}

// 建立/router/web.php
func CreateRouter(projectName string) {
	var router_txt string
	directory := localhostDir + "/" + projectName + RouterDir
	router_txt = router_head + router_web_txt + router_footer
	saveName := fmt.Sprintf("web.php")
	err := ioutil.WriteFile(directory+saveName, []byte(router_txt), 0644)
	ChkErr(err)
}

// 建立Model .php
func CreateModel(tablename, projectName string, tableid int32) {
	field_setting := ""
	fields_txt := ""
	directory := localhostDir + "/" + projectName + ModelDir
	fileName := utils.RemoveS(utils.FirstUpper(tablename))
	className := fmt.Sprintf("class %s extends BaseModel", fileName)
	table := fmt.Sprintf("  protected $table = '%s';", utils.FirstLower(tablename))
	arg := sql.NullInt32{Int32: tableid, Valid: true}
	tbs, _ := dbpg.GetTFBytID(ctx, arg)
	for _, row := range tbs {
		fields_txt += fmt.Sprintf("'%s',", row.FieldName.String)
		field_setting += fmt.Sprintf("          '%s' => [\n", row.FieldName.String)
		field_setting += fmt.Sprintf("               'type' => '%s',\n", row.ModelType.String)
		field_setting += fmt.Sprintf("               'required' => %d,\n", row.IsRequire.Int32)
		field_setting += fmt.Sprintf("               'search' => [\n")
		field_setting += fmt.Sprintf("                    'level' => 'like',\n")
		field_setting += fmt.Sprintf("               ]\n")
		field_setting += fmt.Sprintf("          ],\n")
	}
	model_txt := model_1 + className + model_2 + table
	model_txt += model_3 + "       " + fields_txt + model_4
	model_txt += model_5 + field_setting + model_6
	saveName := fmt.Sprintf("%s.php", fileName)
	err := ioutil.WriteFile(directory+saveName, []byte(model_txt), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 建立migration create table
func MigrationTable(rep_model, projectName string, tableid int32) {
	var txt string
	var combinedMigration string
	directory := localhostDir + "/" + projectName + DatabaseDir
	arg := sql.NullInt32{Int32: tableid, Valid: true}
	tbs, _ := dbpg.GetTFBytID(ctx, arg)
	for _, row := range tbs {
		txt += fmt.Sprintf("            $table->%s(\"%s\");\n", row.Migration.String, row.FieldName.String)
	}
	// 不区分大小写替换 "xnews" 为 "abc"
	re := regexp.MustCompile(`(?i)xnews`)
	migration_head = re.ReplaceAllString(migration_head, rep_model)
	migration_end = re.ReplaceAllString(migration_end, rep_model)
	// 获取当前时间
	currentTime := time.Now()
	// 格式化时间为 "2006_01_02_150405"
	migration_head1 := fmt.Sprintf("class Create%sTable extends Migration", utils.FirstUpper(rep_model))
	migration_head3 := fmt.Sprintf("\n    Schema::create('%s', function (Blueprint $table) {", utils.FirstLower(rep_model))
	migration_latest := fmt.Sprintf("\n     Schema::dropIfExists('%s');", utils.FirstUpper(rep_model))
	timeFormat := currentTime.Format("2006_01_02_150405")
	// 将替换后的字符串写入文件  拼接文件名前缀和时间格式
	fileName := timeFormat + fmt.Sprintf("_create_%s_table.php", utils.FirstLower(rep_model))
	combinedMigration = migration_head + migration_head1 + migration_head2
	combinedMigration += migration_head3 + migration_head4 + txt + migration_end + migration_latest + migration_end1
	err := ioutil.WriteFile(directory+fileName, []byte(combinedMigration), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
