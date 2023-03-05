/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 09:59
 */

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
	http.ListenAndServe(":8080", routers.Router)

}

//func open() {
//
//	time.Sleep(1000 * time.Millisecond)
//
//	httpUrl := viper.GetString("app.httpUrl")
//	httpUrl = "http://" + httpUrl + "/home/index"
//
//	fmt.Println("访问页面体验:", httpUrl)
//
//	cmd := exec.Command("open", httpUrl)
//	cmd.Output()
//}
