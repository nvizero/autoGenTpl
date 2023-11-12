package cache

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Conn2Redis() *redis.Client {
	// 创建一个Redis客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 如果有密码，请提供
		DB:       0,                // 默认的数据库
	})
	return client
}

func writeToRedis(client *redis.Client, data map[int]string, hashKey string) error {
	// 写入哈希表
	for key, value := range data {
		err := client.HSet(ctx, hashKey, fmt.Sprintf("%d", key), value).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetRedis(name string) bool {
	client := Conn2Redis()
	defer client.Close()
	d := client.Get(ctx, name)
	b, err := d.Bool()
	if err != nil {
		fmt.Println(err)
	}
	if b {
		return false
	}
	return false
}

func SetRedis(name string) {
	client := Conn2Redis()
	defer client.Close()
	client.Set(ctx, name, 1, 3600*time.Millisecond)
}

func readFromRedis(client *redis.Client, hashKey string) (map[string]string, error) {
	// 从哈希表中获取所有元素
	hashData, err := client.HGetAll(ctx, hashKey).Result()
	if err != nil {
		return nil, err
	}
	return hashData, nil
}

func ConnRedis() {
	client := Conn2Redis()
	defer client.Close()

	data := map[int]string{
		1: "aa",
		2: "bb",
		3: "dd",
	}

	hashKey := "myHash"

	err := writeToRedis(client, data, hashKey)
	if err != nil {
		log.Fatal(err)
	}

	hashData, err := readFromRedis(client, hashKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Elements in the hash table:")
	for key, value := range hashData {
		fmt.Printf("%s:%s\n", key, value)
	}

	var listData []string
	for _, value := range hashData {
		listData = append(listData, value)
	}

	fmt.Println("Elements in the list:")
	fmt.Println(strings.Join(listData, ", "))
}
