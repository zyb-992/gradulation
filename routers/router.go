// Package routers
// @Author: Dummy
// @File: router.go
// @Date: 2023/2/2 18:12

package routers

import "github.com/gin-gonic/gin"

var Router = gin.Default()

func InitRouter() {
	Init(Router)
	WebsocketInit()
}
