package utils

import (
	"bufio"
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
		fmt.Println(cmdStr)
		log.Fatal("錯--", err)
	}
	statusChan <- string(out)
	// 在这里将消息发送到 WebSocket 客户端
	fmt.Printf("%s\n", out)

	// 在这里将消息发送到 WebSocket 客户端
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

func ExecCmd2(cmdStr string, statusChan chan string) {
	command := exec.Command("/bin/sh", "-c", cmdStr)
	stdout, err := command.StdoutPipe()
	if err != nil {
		fmt.Println(err, cmdStr)
	}
	stderr, err := command.StderrPipe()
	if err != nil {
		fmt.Println(err, cmdStr)
	}

	if err := command.Start(); err != nil {
		log.Fatal("-錯-", err)
	}

	go func() {
		defer command.Wait()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			statusChan <- scanner.Text()
		}
	}()

	go func() {
		defer command.Wait()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			statusChan <- scanner.Text()
		}
	}()

	// 等待命令执行完成
	if err := command.Wait(); err != nil {
		fmt.Println("命令执行失败:", err)
	}

	//statusChan <- string(out)
}
