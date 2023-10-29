package control

import (
	"database/sql"
	"fmt"
	"tpl/cache"
	db "tpl/db/sqlc"
	"tpl/utils"

	redis "github.com/go-redis/redis/v8"
)

// var 	utils.GenController(ControllerDir, ContName, TableName)

type GenerLaravel struct {
	Projects      []db.Project
	Pid           int32
	Pg            *db.Queries
	Redis         *redis.Client
	projectName   string
	projectDir    string
	ControllerDir string
}

// query project
func (p *GenerLaravel) QueryPj() {
	arg := db.ListProjectsParams{
		Limit:  10,
		Offset: 0,
	}
	project, err := p.Pg.ListProjects(ctx, arg)
	ChkErr(err)
	p.Projects = project
}

// 建立laravel Controller and migration
func (p *GenerLaravel) GenLaravelControllerMigration(tb string) {
	//首字母大寫
	contName := utils.FirstUpper(tb)
	//首字母小寫
	tbName := utils.FirstLower(tb)
	//用table name 產生 Controller
	utils.GenController(p.ControllerDir, contName, tbName)

	params = []interface{}{p.projectName, cmd_sh[1]}
	utils.RunCmd(params, docker_run_sh, statusChan)

	//migrate table
	params = []interface{}{p.projectName, tbName, tbName}
	utils.RunCmd(params, migration_table2, statusChan)
	// 取db
}

// query project's table
func (p *GenerLaravel) GenerateLaravelModel() {

	if p.Projects != nil {
		arg := db.WhereTbByPIDParams{
			Limit:     10,
			Offset:    0,
			ProjectID: sql.NullInt32{Int32: p.Projects[0].ID, Valid: true},
		}
		tbs, err := p.Pg.WhereTbByPID(ctx, arg)
		ChkErr(err)
		for i, tb := range tbs {
			fmt.Println(i, tb, "--------------- empty")
			tableid := sql.NullInt32{Int32: tb.ID, Valid: true}
			tfs, _ := p.Pg.GetTFByfID(ctx, tableid)
			//執行docker 中 migration
			p.GenLaravelControllerMigration(tb.Name.String)
			fmt.Println(tfs, "---empty")
		}
	}
}

func (p *GenerLaravel) DoTest() {
	p.QueryPj()
	p.GenerateLaravelModel()
	//query fields
	//query table ref fields
}

// 自動generate laravel
func TestGenLaravel() {
	gl := GenerLaravel{
		Pg:            db.ConnDev(),
		Redis:         cache.Conn2Redis(),
		projectName:   fmt.Sprintf("%s%d", project_name, No),
		projectDir:    localhostDir + "/" + fmt.Sprintf("%s%d", project_name, No),
		ControllerDir: localhostDir + "/" + fmt.Sprintf("%s%d", project_name, No) + ControllerDir,
	}
	gl.DoTest()
}

func TrancateData() {
	db.TrunateDB()
}
