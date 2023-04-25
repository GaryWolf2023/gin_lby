package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoleList(c *gin.Context) {
	fmt.Println("query参数", c.Query("name"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "获取接口权限列表",
	})
}
func CreateRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建接口权限",
	})
}
func ModifyRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "修改接口权限",
	})
}
func DeleteRole(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "删除接口权限",
	})
}
func SearchRole(c *gin.Context) {
	c.String(http.StatusOK, "检索接口权限")
}
