package user

import "github.com/gin-gonic/gin"

func UserEditDisplayAction(c *gin.Context) {
	// urlからidを取得し振り分け
	id := c.Param("id")
	c.HTML(200, "user-edit.html", gin.H{
		// htmlにidの値を渡す
		"id": id,
	})
}
