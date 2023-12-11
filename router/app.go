package router

import (
	"demo/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	r:=gin.Default()
	r.GET("/index",service.GetIndex)
	r.GET("/api/userlist",service.GetUserList)
	return r
}