package control

import (
	"context"
	"database/sql"
	"fmt"
	"tpl/cache"
	db "tpl/db/sqlc"
	"tpl/utils"

	redis "github.com/go-redis/redis/v8"
)

var rkey string = "fields"
var ctx = context.Background()

type Pj struct {
	Pname    string
	Pid      int32
	Tables   map[string]map[string]string
	FieldKey map[string]string
	Pg       *db.Queries
	Redis    *redis.Client
}

func (p *Pj) ChkProjectName() bool {
	arg := sql.NullString{String: p.Pname, Valid: true}
	project, err := p.Pg.GetProjectByName(context.Background(), arg)
	ChkErr(err)
	if project.ID == 0 {
		return false
	}
	return true
}

func (p *Pj) GenProject() {
	arg := db.CreateProjectsParams{
		Name:  sql.NullString{String: p.Pname, Valid: true},
		IsGen: sql.NullInt32{Int32: 0, Valid: true},
	}
	project, err := p.Pg.CreateProjects(context.Background(), arg)
	ChkErr(err)
	p.Pid = project.ID
}

func (p *Pj) GenTable(table string) int32 {
	arg := db.CreateTbParams{
		Name:      sql.NullString{String: table, Valid: true},
		ProjectID: sql.NullInt32{Int32: p.Pid, Valid: true},
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
		for k, v := range fields {
			fmt.Println("====", k, v)
			arg := db.CreateTbFieldParams{
				TableID:    sql.NullInt32{Int32: tableId, Valid: true},
				FieldName:  sql.NullString{String: k, Valid: true},
				LaravelMap: sql.NullString{String: v, Valid: true},
			}
			tf, err := p.Pg.CreateTbField(context.Background(), arg)
			ChkErr(err)
			fmt.Println(tf)

		}
	}
}

// check has redis
func (p *Pj) ChkFieldRedis() {
	// if !p.CheckFieldsRedis(rkey) {
	// 	var tmp map[string]string
	// 	tmp = make(map[string]string) // 初始化 tmp 映射
	// 	for _, row := range fls {
	// 		tmp[row.Name.String] = fmt.Sprintf("%d", row.ID)

	// 	}
	// 	p.WriteToRedis(tmp, rkey)
	// 	p.FieldKey = tmp
	// } else {
	// 	p.FieldKey = p.ReadFromRedis(rkey)
	// }
}

func (p *Pj) Controller() {
	//是否有相同名稱
	// if !p.ChkProjectName() {
	p.GenProject()
	//查看field是否在redis
	// p.ChkFieldRedis()
	p.Parse2TableField()
	// }
}

func InitData() {
	table := map[string]map[string]string{
		"qbbts": {
			"title": "string",
			"body":  "text",
		},
		"gaews": {
			"title": "string",
			"local": "string",
			"body":  "ckeditor",
			"cover": "string",
		},
	}
	opj := Pj{
		Pg:     db.ConnDev(),
		Redis:  cache.Conn2Redis(),
		Pname:  "isb2",
		Tables: table,
	}
	opj.Controller()
}

func ChkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
