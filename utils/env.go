package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(env string) string {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		fmt.Println("无法加载 .env 文件")
	}
	// 读取名为 "MY_ENV_VAR" 的环境变量的值
	myEnvVar := os.Getenv(env)

	// 检查环境变量是否存在
	if myEnvVar == "" {
		fmt.Println("环境变量 MY_ENV_VAR 未设置或为空")
	}
	return myEnvVar
}
