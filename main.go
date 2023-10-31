package main

import (
	"fmt"
	"sync"
	"tpl/db"
)

var dbname string = "golang_gen_laravel"
var wg sync.WaitGroup
var status = make(chan string, 20) // 使用 make 函数初始化通道

func main() {
	//wg.Add(1)
	//go control.GenLaravel(status, &wg)
	// control.CHttp()
	//取出資料
	//	go rece()
	//wg.Wait()
	//createdb()
	//control.TestAI()
	// DevQueries := db.ConnPq()
	// defer DevQueries.Close()
	//插入資料
	//generate
	//    0.get data from web [not ya]
	//    1.init data to postgresql
	//    2.run laravel docker
	//    3.run laravel controller database model router sidemenu
	//建立假資料
	// lara := control.InitFakeData()
	//測試
	//control.TestGenLaravel()
	//建立laravel DB資料
	// control.GetTbGenerateMigrateTable(lara.ProjectID, lara.ProjectName)

	// control.GetTbGenerateMigrateTable(61, "isb32")
	//清除測試資料
	//	control.TrancateData()
	rece()
}

func createdb() {
	db.CreateOrRecreateDB(dbname)
}

type Animal struct {
	Field     string
	ShowName  string
	Migrate   string
	ModelType string
	Require   string
}

func rece() {
	data := map[string][]Animal{
		"cats": {
			{"name", "名字", "aa", "aa", "aa"},
			{"sex", "text", "aa", "aa", "aaaa"},
		},
		"cats1": {
			{"name", "text", "aa", "aa", "aa"},
			{"name", "text", "aaaa", "aa", "aa"},
		},
	}

	for a, b := range data {
		fmt.Println(a, b)
		for v, d := range b {
			fmt.Println(v, d.Field)
		}
	}
}
