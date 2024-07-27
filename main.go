package main

import (
	"github.com/dev/test/server"
	"os"
)

// @title test
// @version 1.0
// @description this is an application web of interview assessment tests for interviewing out of the box .
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api

func main() {

	os.Setenv("key", "123456781234567812345678")
	server.Init()

}
