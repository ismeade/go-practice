// Package main provides ...
package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	fmt.Println(viper.Get("a"))
	fmt.Println(viper.Get("c.c1"))

	for {
		time.Sleep(3 * time.Second)
		fmt.Println(viper.Get("name"))
	}
}
