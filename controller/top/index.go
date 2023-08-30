// topディレクトリの機能をパッケージ化
package top

import "github.com/gin-gonic/gin"

// 他のパッケージからの読み込みに対応する為、頭文字を大文字にする
func IndexDisplayAction(c *gin.Context) {
	// トップ画面を返すindex.html
	c.HTML(200, "index.html", gin.H{})
}
