package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 建立ENV
func UpdateEnv(prodName, localDir, port, laravelPath string) {

	fmt.Println(laravelPath, "enenenv----")
	if err := os.Chdir(localDir); err != nil {
		fmt.Printf("Failed to change the working directory: %s\n", err)
		return
	}
	fmt.Println(laravelPath, "----")
	envPath := fmt.Sprintf("%s/initEnv", laravelPath)
	fmt.Println(envPath)
	// 打开配置文件
	file, err := os.Open(envPath)
	if err != nil {
		fmt.Println("无法打开配置文件:", err)
		return
	}
	defer file.Close()

	// 创建一个 map 用于存储配置变量
	config := make(map[string]string)

	// 逐行读取文件内容并解析配置变量
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			config[key] = value
		}
	}
	// 输出所有配置变量
	newFilePath := localDir + "/" + prodName + "/.env" // 新文件的完整路径
	fmt.Println(newFilePath)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("无法创建新配置文件:", err)
		return
	}
	defer newFile.Close()
	// 修改特定变量
	config["DB_DATABASE"] = prodName
	config["APP_URL"] = fmt.Sprintf("http://127.0.0.1:%s/", port)
	config["DB_HOST"] = fmt.Sprintf("172.17.0.5")
	config["APP_KEY"] = fmt.Sprintf("base64:yA/8JhlITdW/zo5/+A6XKrbLppTlA1bhTcspdeUjIDs=")
	// 创建一个新文件并将修改后的配置写入其中
	// 写入修改后的配置到新文件
	for key, value := range config {
		line := key + "=" + value + "\n"
		_, err := newFile.WriteString(line)
		if err != nil {
			fmt.Println("无法写入新配置文件:", err)
			return
		}
	}
	fmt.Printf("\n已保存修改后的配置到 %s .env", prodName)
}

// 產生Controller
func GenController(dir, cname, table string) {
	// 读取 PHP 文件内容
	// Check the value of cate to determine the file to read
	phpContent, err := ioutil.ReadFile(dir + "PostController.php")
	if err != nil {
		fmt.Println("无法读取 PHP 文件:", err)
		return
	}
	phpStr := string(phpContent)
	newPhpStr := strings.Replace(phpStr, "posts", table, -1)
	newPhpStr = strings.Replace(newPhpStr, "Post", cname, -1)
	// 将修改后的内容写回到文件
	err = ioutil.WriteFile(dir+cname+"Controller.php", []byte(newPhpStr), 0644)
	if err != nil {
		fmt.Println("无法写入修改后的 PHP 文件:", err)
		return
	}
	fmt.Printf("\n已成功将 PHP 文件中的 'Post' 修改为 '%s'\n", cname)
	fmt.Printf("\n已成功建立talbe '%s'\n", table)
}
