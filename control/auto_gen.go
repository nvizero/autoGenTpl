package control

import (
	"fmt"
	"tpl/db"
	"tpl/utils"
)

// 自動執行專案docker
// func GenLaravel(statusChan chan string, wg *sync.WaitGroup) {
func GenLaravel(statusChan chan string) {
	//defer wg.Done()
	var projectName string
	var dockerPort string
	var port string
	if Project.ProjectName != "" {
		projectName = fmt.Sprintf("%s", Project.ProjectName)
		dockerPort = fmt.Sprintf("%d:80", Project.DockerPort)
		port = fmt.Sprintf("%d", Project.DockerPort)
	} else {
		// projectName = fmt.Sprintf("%s%d", project_name, No)
		// dockerPort = fmt.Sprintf("%d%d:80", port_serial, No)
		// port = fmt.Sprintf("%d%d", port_serial, No)
	}
	visit := fmt.Sprintf("http://127.0.0.1:%s", port)
	projectDir := localhostDir + "/" + projectName
	dockerVPara := projectDir + ":" + dockerDir

	//建立DB
	db.CreateOrRecreateDB(projectName)
	statusChan <- "create db is ok"
	//git clone到本地
	GitCloneProject(gitRepo, projectName, localhostDir, statusChan)
	statusChan <- "Git Clone Completed"
	//設定 env
	utils.UpdateEnv(projectName, localhostDir, port, projectDir)
	statusChan <- "set env is Success"
	//建laravel 執行基本參數
	//	wg.Add(1)
	RunDockerCmd(projectName, dockerVPara, dockerPort, dockerHubImg, laravel_update, statusChan)

	//	建立額外laravel功能
	//statusChan <- "建立額外laravel功能"
	//	Laravel(projectName, ControllerDir, ControllerName, TableName, statusChan)

	//statusChan <- "推上git分支"
	//GitPushBranch(projectName, projectDir, statusChan)
	defer func() {
		statusChan <- "generate laravel docker success"
		statusChan <- fmt.Sprintf("建立專案%s完成!!", projectName)
		statusChan <- fmt.Sprintf("請訪問<a href=\"%s\" target=\"_blank\"  >%s</a>", visit, visit)
	}()
}
