package main

import (
	"gowebsocket/init"
	"gowebsocket/routers"
	"gowebsocket/servers/task"
	"gowebsocket/servers/websocket"
	"net/http"
)

func main() {
	initres.Init()

	// 初始化路由
	routers.InitRouter()

	// 定时任务
	task.Init()
	// 服务注册
	task.ServerInit()

	go websocket.StartWebSocket()

	//go open()

	// config的httpPort
	err := http.ListenAndServe(":8080", routers.Router)
	if err != nil {
		panic("Http Listen Error")
	}
}
