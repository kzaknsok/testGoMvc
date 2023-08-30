package main

import "example.com/server"

func main() {
	// sever/router.goからfuncを参照
	r := server.GetRouter()
	r.Run(":8080")
	// 3回目のgit pushのために追加したコメント
}
