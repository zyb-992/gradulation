package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowebsocket/common"
	"gowebsocket/controllers"
	"gowebsocket/lib/cache"
	"gowebsocket/models"
	"gowebsocket/servers/websocket"
	"strconv"
)

// 查看全部在线用户
// http://localhost:8080/user/list?appId=102
// 查询102号房间的用户
func List(c *gin.Context) {

	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 查看全部在线用户", appId)

	data := make(map[string]interface{})

	userList := websocket.UserList(appId)
	data["userList"] = userList
	data["userCount"] = len(userList)

	controllers.Response(c, common.OK, "", data)
}

// 查看用户是否在线
func Online(c *gin.Context) {

	userId := c.Query("userId")
	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 查看用户是否在线", userId, appIdStr)

	data := make(map[string]interface{})

	online := websocket.CheckUserOnline(appId, userId)
	data["userId"] = userId
	data["online"] = online

	controllers.Response(c, common.OK, "", data)
}

// 给全员发送消息
func SendMessageAll(c *gin.Context) {
	// 获取参数
	appIdStr := c.PostForm("appId")
	userId := c.PostForm("userId")
	msgId := c.PostForm("msgId")
	message := c.PostForm("message")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 给全体用户发送消息", appIdStr, userId, msgId, message)

	data := make(map[string]interface{})
	if cache.SeqDuplicates(msgId) {
		fmt.Println("给用户发送消息 重复提交:", msgId)
		controllers.Response(c, common.OK, "", data)

		return
	}

	sendResults, err := websocket.SendUserMessageAll(appId, userId, msgId, models.MessageCmdMsg, message)
	if err != nil {
		data["sendResultsErr"] = err.Error()

	}

	data["sendResults"] = sendResults

	controllers.Response(c, common.OK, "", data)

}
