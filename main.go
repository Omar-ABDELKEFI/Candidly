package main

import (
	"github.com/tekab-dev/tekab-test/server"
)

// @title tekab-test
// @version 1.0
// @description this is an application web of interview assessment tests for interviewing out of the box .
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	server.Init()

}
