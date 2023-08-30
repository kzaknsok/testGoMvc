package server

import (
	"example.com/controller/top"
	"example.com/controller/user"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")

	r.GET("/", top.IndexDisplayAction)
	r.GET("/user", user.UserListDisplayAction)
	r.GET("/user/add", user.UserAddDisplayAction)
	r.GET("/user/edit/:id", user.UserEditDisplayAction)

	return r
}
