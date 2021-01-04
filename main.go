package main

import (
	"github.com/josephshih13/short-url/redis"
	"github.com/josephshih13/short-url/server"
)

func main() {
	redis.ClientInit()
	server.InitServer()

}
