package utils

import (
	"fmt"
	"log"
	"os/exec"
)

// 執行DOCKER 內命令
func RunCmd(params []interface{}, cmd string, statusChan chan string) {
	cmdStr := fmt.Sprintf(cmd, params...)
	fmt.Println("-- run --\n", cmdStr)
	//out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		log.Fatal("--", err)
	}
	//statusChan <- string(out)
	fmt.Printf("%s\n", out)
}

// 執行命令
func ExecCmd(cmd *exec.Cmd) string {
	// 获取命令的输出
	fmt.Println(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Command failed to run: %s\n", err)
		return ""
	}
	// 将命令的输出内容存储在变量中
	outputStr := string(output)
	fmt.Println(outputStr)
	return outputStr
}
