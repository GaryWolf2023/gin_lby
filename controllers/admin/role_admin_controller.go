package admin

import (
	"Gin_one/controllers/base"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleAdminController struct {
	base.BaseController
}

func (roleadmin RoleAdminController) List(c *gin.Context) {
	fmt.Println("query参数", c.Query("name"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "获取接口权限列表",
	})
}
func (roleadmin RoleAdminController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建接口权限",
	})
}
func (roleadmin RoleAdminController) Modify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "修改接口权限",
	})
}
func (roleadmin RoleAdminController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "删除接口权限",
	})
}
