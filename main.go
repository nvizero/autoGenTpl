package main

import (
	"sync"
	"tpl/control"
)

var dbname string = "golang_gen_laravel"
var wg sync.WaitGroup
var status = make(chan string, 20) // 使用 make 函数初始化通道

func main() {
	//wg.Add(1)
	//go control.GenLaravel(status, &wg)
	control.CHttp()
	//取出資料
	//wg.Wait()
	//createdb()
	//control.TestAI()
	// DevQueries := db.ConnPq()
	// defer DevQueries.Close()
	//插入資料
	//    0.get data from web [not ya]
	//建立假資料
	// lara := control.InitFakeData()
	//建立laravel DB資料
	//control.GetTbGenerateMigrateTable(lara.ProjectID, lara.ProjectName)

	// control.GetTbGenerateMigrateTable(61, "isb32")
	//清除測試資料
	//control.TrancateData()
}
