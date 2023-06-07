package cmd

import "fmt"

// 服务退出 优雅的关闭程序
func Close() {
	fmt.Println("服务停止，我们可以在这里做一些清理工作")
	clean()
}

func clean() {
	fmt.Println("清理工作")
}
