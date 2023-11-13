package control

import (
	"fmt"
	"tpl/utils"

	"github.com/gorilla/websocket"
)

// 建立額外laravel功能
func Laravel(projectName, ControllerDir, ContName, TableName string, statusChan chan string, ws *websocket.Conn) {

	//建立Controller
	utils.GenController(ControllerDir, ContName, TableName)

	//建立model
	params = []interface{}{projectName, ModelName}
	utils.RunCmd(params, gen_model, statusChan)

	//migrate table
	params = []interface{}{projectName, CreateTableName}
	utils.RunCmd(params, migration_table, statusChan)
}

// git clone 下來
func GitCloneProject(gitRepo, projectName, localhostDir string, statusChan chan string) {
	params = []interface{}{localhostDir, gitRepo, projectName}
	utils.RunCmd(params, git_clone, statusChan)
}

// 推到分支
func GitPushBranch(projectName, projectDir string, statusChan chan string) {
	//建立分支
	statusChan <- "建立分支"
	params = []interface{}{projectDir, projectName}
	utils.RunCmd(params, git_branch, statusChan)

	//切到分支
	statusChan <- "切到分支"
	params = []interface{}{projectDir, projectName}
	utils.RunCmd(params, git_checkout, statusChan)

	//分支加入修改檔案
	params = []interface{}{projectDir}
	utils.RunCmd(params, git_add, statusChan)

	//分支記錄
	params = []interface{}{projectDir, projectName}
	utils.RunCmd(params, git_commit, statusChan)

	//推上分支
	statusChan <- "推上分支"
	params = []interface{}{projectDir, projectName}
	utils.RunCmd(params, git_push, statusChan)
}

// 跑docker 內的指令
// 1.建立本機docker
func RunDockerCmd(projectName, dockerVPara, dockerPort, dockerHubImg, laravel_update_cmd string, statusChan chan string) {
	//	defer wg.Done()
	//建立本機docker
	statusChan <- "建立本機docker"
	params = []interface{}{dockerPort, projectName, dockerVPara, dockerHubImg}
	utils.RunCmd(params, gen_docker_cmd, statusChan)

	//init 參數
	statusChan <- "init 參數"
	params = []interface{}{projectName, cmd_sh[0]}
	utils.RunCmd(params, docker_run_sh, statusChan)

	//composer update
	statusChan <- "composer update..."
	params = []interface{}{projectName}
	utils.RunCmd(params, laravel_update_cmd, statusChan)

	//建立 基本 database 執行php artisan migrate
	// statusChan <- "建立 基本 database 執行php artisan migrate"
	// params = []interface{}{projectName, cmd_sh[1]}
	// utils.RunCmd(params, docker_run_sh, statusChan)
}

var txt = `
  hi
  hwllo world
  `

func TestAI() {
	env := "OPEN_AI_KEY"
	key := utils.GetEnv(env)
	q := fmt.Sprintf("參考%s,用%s 寫出php 代碼", txt, "name type=text,sex type=text,")
	response, err := utils.Openai(key, q)
	if err != nil {
		fmt.Println("Openai 函数执行时出错:", err)
		return
	}
	// 处理解析后的响应数据
	fmt.Printf(response.Choices[0].Message.Content)
}
