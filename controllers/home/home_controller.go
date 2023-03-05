package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gowebsocket/servers/websocket"
	"net/http"
	"strconv"
)

// 聊天页面
func Index(c *gin.Context) {
	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)
	if !websocket.InAppIds(appId) {
		appId = websocket.GetDefaultAppId()
	}

	fmt.Println(" 聊天首页: ", appId)

	data := gin.H{
		"title":        "聊天首页",
		"appId":        appId,
		"httpUrl":      viper.GetString("app.httpUrl"),
		"webSocketUrl": viper.GetString("app.webSocketUrl"),
	}
	c.HTML(http.StatusOK, "index.tpl", data)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
