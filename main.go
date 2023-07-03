package main

import (
	"Gin_one/cmd"
	"Gin_one/utils"
	"fmt"
)

// @title Gin_one API文档
// @description Gin_one 学习记录关于go/gin的
// @version 0.0.1
func main() {

	defer cmd.Close()
	cmd.Start()

	token, _ := utils.CreateToken(1, "lby")
	fmt.Println(token)

	iJwtCustClaims, _ := utils.ParseToken(token)
	fmt.Println(iJwtCustClaims)
}
