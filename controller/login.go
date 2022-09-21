/*
 * @Author: GG
 * @Date: 2022-09-20 15:00:53
 * @LastEditTime: 2022-09-21 17:53:18
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\controller\login.go
 *
 */
package controller

import (
	"fmt"
	"gin-blog/common/global"
	"gin-blog/common/utils"
	"gin-blog/models"
	"gin-blog/service"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var UserService service.UserService
var SessionUtil utils.SessionUtils

func Login(c *gin.Context) {
	method := c.Request.Method
	fmt.Printf("method: %v\n", method)

	if method == http.MethodPost {
		username := c.PostForm("username")
		password := c.PostForm("password")

		user, exist := UserService.GetUserByUsername(username)
		if !exist {
			gintemplate.HTML(c, http.StatusOK, "login.html", gin.H{"error": "用户不存在"})
			return
		}

		if !UserService.ComparePwd(user.Password, password) {
			gintemplate.HTML(c, http.StatusOK, "login.html", gin.H{"error": "密码不正确"})
			return
		}

		// 记住登录，加长session 时间
		remember := c.PostForm("remember")
		maxage := 0
		fmt.Printf("remember: %v\n", remember)
		if remember != "" {
			maxage = utils.MAX_AGE_REMEMBER
		}

		// session
		SessionUtil.Set(c, "userinfo", user, true, maxage)

		// 放入context上下文,方便中间件调用
		c.Set("userLogin", user)
		// c.Next()
		c.Redirect(http.StatusFound, "post/list")
	}

	if method == http.MethodGet {
		gintemplate.HTML(c, http.StatusOK, "login.html", gin.H{})
	}
}

func Regiest(c *gin.Context) {
	var user models.User
	user.Username = "amdin"
	user.Password = UserService.GenPwd("password")

	r := global.DB.Table("users").Where("id = ?", 1).Update("password", user.Password)
	fmt.Printf("r.RowsAffected: %v\n", r.RowsAffected)
	fmt.Printf("r.Error: %v\n", r.Error)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()

	c.Redirect(http.StatusFound, "login")
}
