package control

import (
	"context"
	"database/sql"
	"fmt"
	db "tpl/db/sqlc"
	"tpl/utils"

	redis "github.com/go-redis/redis/v8"
)

var rkey string = "fields"
var ctx = context.Background()

type Pj struct {
	ProjectName string
	ProjectID   int32
	DockerPort  int32
	Tables      map[string][]LaraSetting
	FieldKey    map[string]string
	Pg          *db.Queries
	Redis       *redis.Client
}

func (p *Pj) ChkProjectName() bool {
	arg := sql.NullString{String: p.ProjectName, Valid: true}
	project, err := p.Pg.GetProjectByName(context.Background(), arg)
	ChkErr(err)
	if project.ID == 0 {
		return false
	}
	return true
}

func (p *Pj) GenProject() {
	arg := db.CreateProjectsParams{
		Name:  sql.NullString{String: p.ProjectName, Valid: true},
		Port:  sql.NullInt32{Int32: p.DockerPort, Valid: true},
		IsGen: sql.NullInt32{Int32: 0, Valid: true},
	}
	project, err := p.Pg.CreateProjects(context.Background(), arg)
	ChkErr(err)
	p.ProjectID = project.ID
}

func (p *Pj) GenTable(table string) int32 {
	arg := db.CreateTbParams{
		Name:      sql.NullString{String: table, Valid: true},
		ProjectID: sql.NullInt32{Int32: p.ProjectID, Valid: true},
		Describe:  sql.NullString{String: utils.RandomString(20), Valid: true},
	}
	t, err := p.Pg.CreateTb(context.Background(), arg)
	ChkErr(err)
	return t.ID
}

func (p *Pj) FieldExistsInHash(client *redis.Client, hashKey, field string) (bool, error) {
	exists, err := client.HExists(ctx, hashKey, field).Result()
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (p *Pj) ReadFromRedis(hashKey string) map[string]string {
	// 从哈希表中获取所有元素
	hashData, err := p.Redis.HGetAll(ctx, hashKey).Result()
	ChkErr(err)
	return hashData
}

func (p *Pj) CheckFieldsRedis(hashKey string) bool {
	var b = false
	fieldToCheck := "string" // 要检查的字段
	exists, err := p.FieldExistsInHash(p.Redis, hashKey, fieldToCheck)
	ChkErr(err)
	if exists {
		b = true
	}
	return b
}

// 写入哈希表
func (p *Pj) WriteToRedis(data map[string]string, hashKey string) {
	for key, value := range data {
		fmt.Println(fmt.Sprintf("%s", value), "---")
		err := p.Redis.HSet(ctx, hashKey, fmt.Sprintf("%s", key), value).Err()
		ChkErr(err)
	}
}

func (p *Pj) Parse2TableField() {
	for table, fields := range p.Tables {
		tableId := p.GenTable(table)
		for _, v := range fields {
			arg := db.CreateTbFieldParams{
				TableID:   sql.NullInt32{Int32: tableId, Valid: true},
				FieldName: sql.NullString{String: v.Field, Valid: true},
				Migration: sql.NullString{String: v.Migration, Valid: true},
				ShowName:  sql.NullString{String: v.ShowName, Valid: true},
				ModelType: sql.NullString{String: v.ModelType, Valid: true},
				IsRequire: sql.NullInt32{Int32: v.IsRequire, Valid: true},
			}
			_, err := p.Pg.CreateTbField(context.Background(), arg)
			ChkErr(err)
			//fmt.Println(tf)

		}
	}
}

func (p *Pj) Controller() {
	//是否有相同名稱
	if !p.ChkProjectName() {
		p.GenProject()
		p.Parse2TableField()
	}
}

func InitFakeData() Pj {
	fdata := map[string][]LaraSetting{
		"cats": {
			{"name", "名字", "string", "text", 1},
			{"sex", "性別", "text", "ckeditor", 1},
		},
		"dogs": {
			{"name", "名字", "string", "text", 1},
			{"sex", "性別", "text", "ckeditor", 1},
		},
	}
	Project = Pj{
		Pg:          db.ConnDev(),
		ProjectName: "isb20",
		DockerPort:  2020,
		Tables:      fdata,
	}
	Project.Controller()
	GenLaravel(statusChan)
	return Project
}

func ChkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
