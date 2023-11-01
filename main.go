package main

import (
	"github.com/x14n/x14n-gin-api/initliza"
	"github.com/x14n/x14n-gin-api/server"
)

func main() {
	initliza.InitServer()

	server.RunServer()

}
