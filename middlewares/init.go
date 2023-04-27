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
	if req_header["Token"] != nil {
		fmt.Println("Token:", req_header["Token"])
	} else {
		fmt.Println("Token不存在")
	}
	// 因为主程序是永远不会退出的，所以说不用担心子协程会被提前退出
	cCp := ctx.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("我是子协程")
		fmt.Println(cCp.Request.URL)
		fmt.Println(cCp.Request.Header)
		fmt.Println(cCp.Get("username"))
		fmt.Println("子协程执行结束")
	}()
	ctx.Next() // 执行下一个中间件
	endTime := time.Now()
	fmt.Printf("程序用时：%v", endTime.Sub(startTime))
	// ctx.Abort() // 阻止程序继续执行，不会执行接下来的所有处理程序
}
