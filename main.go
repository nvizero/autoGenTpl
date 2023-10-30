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
	//generate
	//    0.get data from web [not ya]
	//    1.init data to postgresql
	//    2.run laravel docker
	//    3.run laravel controller database model router sidemenu
	//建立假資料
	lara := control.InitFakeData()
	//測試
	//control.TestGenLaravel()
	//建立laravel DB資料
	control.GetTbGenerateMigrateTable(lara.ProjectID, lara.ProjectName)
	// control.GetTbGenerateMigrateTable(61, "isb32")
	//清除測試資料
	//	control.TrancateData()
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
