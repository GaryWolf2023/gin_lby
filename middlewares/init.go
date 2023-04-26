package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Init(ctx *gin.Context) {
	startTime := time.Now()
	ctx.Set("username", "lby")   //利用中间件传值
	fmt.Println(time.Now())      // 打印当前时间
	fmt.Println(ctx.Request.URL) // 打印当前地址
	req_header := ctx.Request.Header
	// fmt.Println(req_header)
	// for k, v := range req_header {
	// 	fmt.Printf("%v : %v\n", k, v)
	// }
	if req_header["Token"] != nil {
		fmt.Println("Token:", req_header["Token"])
	} else {
		fmt.Println("Token不存在")
	}
	ctx.Next() // 执行下一个中间件
	endTime := time.Now()
	fmt.Printf("程序用时：%v", endTime.Sub(startTime))
	// ctx.Abort() // 阻止程序继续执行，不会执行接下来的所有处理程序
}
