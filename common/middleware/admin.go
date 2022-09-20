/*
 * @Author: GG
 * @Date: 2022-09-20 15:49:20
 * @LastEditTime: 2022-09-20 17:41:02
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\middleware\admin.go
 *
 */
package middleware

import (
	"gin-blog/models"
	"gin-blog/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var UserLogService service.UserLogService

// 登录权限中间件
func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("userinfo")
		if user != nil {
			u, err := user.(models.User)
			if !err || u.ID == 0 {
				c.Redirect(http.StatusFound, "login")
				c.Abort()
			}
			c.Next()
		}
	}
}

// 登录日志记录中间件
func UserLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("userLogin")
		if !exists {
			c.Abort()
		}

		user, _ := value.(models.User)
		UserLogService.Add(user)
		c.Next()
	}
}
