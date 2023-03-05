// Package init
// @Author: Dummy
// @File: init.go
// @Date: 2023/2/2 17:53

package initres

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gowebsocket/lib/redislib"
	"io"
	"os"
)

func Init() {
	initConfig()
	initRedis()
	initFile()

}

func initConfig() {

	viper.SetConfigName("config/app")
	viper.AddConfigPath(".") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config redis:", viper.Get("redis"))

}

// 初始化日志
func initFile() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	// 注释
	//gin.DisableConsoleColor()

	// Logging to a file.
	// log/gin.log
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}

func initRedis() {
	redislib.ExampleNewClient()
}
