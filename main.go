package main

import (
	"sync"
	"tpl/control"
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

	// fname := sql.NullString{String: "test", Valid: true}
	// tb, _ := db.DevQueries.CreateField(context.Background(), fname)
	// fmt.Println(tb)
	//初始化欄位
	//control.InitField()
	//建立假資料
	//control.InitData()
	//測試
	//control.TestGenLaravel()
	//清除測試資料
	//control.TrancateData()
	control.GetTbGenerateMigrateTable(21, "isb33")
}

func createdb() {
	db.CreateOrRecreateDB(dbname)
}

func rece() {
	var strs []string
	for msg := range status {
		//fmt.Println(msg)
		strs = append(strs, msg)
	}
}
