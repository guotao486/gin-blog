/*
 * @Author: GG
 * @Date: 2022-09-20 15:00:53
 * @LastEditTime: 2022-09-20 17:38:18
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\controller\login.go
 *
 */
package controller

import (
	"encoding/gob"
	"fmt"
	"gin-blog/common/global"
	"gin-blog/models"
	"gin-blog/service"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var UserService service.UserService

func Login(c *gin.Context) {
	method := c.Request.Method
	fmt.Printf("method: %v\n", method)

	if method == http.MethodPost {
		username := c.PostForm("username")
		password := c.PostForm("password")

		user, exist := UserService.Login(username, password)
		if !exist {
			gintemplate.HTML(c, http.StatusOK, "login.html", gin.H{"error": "用户不存在"})
			return
		}

		// session
		gob.Register(user)
		session := sessions.Default(c)
		session.Set("userinfo", user)
		session.Set("userinfo1", models.User{})
		session.Save()

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
	user.Password = "password"

	global.DB.Create(&user)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()

	c.Redirect(http.StatusFound, "login")
}