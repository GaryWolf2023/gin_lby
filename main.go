package main

import (
	"Gin_one/cmd"
)

// @title Gin_one API文档
// @description Gin_one 学习记录关于go/gin的
// @version 0.0.1
func main() {

	defer cmd.Close()
	cmd.Start()

}
