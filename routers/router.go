package routers

import "github.com/gin-gonic/gin"

var Router = gin.Default()

func InitRouter() {
	Init(Router)
	WebsocketInit()
}
