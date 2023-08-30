package main

import "example.com/server"

func main() {
	r := server.GetRouter()
	r.Run(":8080")
}
