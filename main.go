package main

import (
	"sync"
	"tpl/control"
)

var dbname string = "golang_gen_laravel"
var wg sync.WaitGroup
var status = make(chan string, 20) // 使用 make 函数初始化通道

func main() {
	control.CHttp()
	//取出資料
	//wg.Wait()
	//control.TestAI()
	//建立假資料
	// lara := control.InitFakeData()
	//建立laravel DB資料
	//control.GetTbGenerateMigrateTable(lara.ProjectID, lara.ProjectName)
	//清除測試資料
	//control.TrancateData()
}
