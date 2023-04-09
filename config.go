package main

import (
	"fmt"
	"log"
	"strconv"
	"sync/atomic"

	"github.com/spf13/viper"
)

// Config 配置内容
type Config struct {
	Key   string
	Port  int
	Token string
}

var gConfig atomic.Value

// LoadConfig 加载配置
func LoadConfig() {
	// 读取配置文件
	viper.SetConfigFile(".env")
	//viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}
	viper.AutomaticEnv()
	// 解析配置文件
	var config Config
	config.Key = fmt.Sprint(viper.Get("OPENAI_API_KEY"))
	config.Port, _ = strconv.Atoi(fmt.Sprint(viper.Get("HTTP_PORT")))
	config.Token = fmt.Sprint(viper.Get("WECHAT_TOKEN"))
	//if err := viper.Unmarshal(&config); err != nil {
	//	panic(fmt.Errorf("解析配置文件失败: %s", err))
	//}
	gConfig.Store(config)
}

// GetConfig 获取配置
func GetConfig() Config {
	return gConfig.Load().(Config)
}

// ShowConfig 展示配置信息
func ShowConfig() {
	config := GetConfig()
	log.Println("====================================================")
	log.Printf("OpenAIKey=%v\n", config.Key)
	log.Printf("Token =%v\n", config.Token)
	log.Printf("port=%v\n", config.Port)
	log.Println("====================================================")
}
