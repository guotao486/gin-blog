/*
 * @Author: GG
 * @Date: 2022-09-21 10:28:20
 * @LastEditTime: 2022-09-21 10:33:02
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\controller\response.go
 *
 */
package controller

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HtmlResponse(ctx *gin.Context, code int, name string, data interface{}) {
	session := sessions.Default(ctx)
	user := session.Get("userinfo")
	gintemplate.HTML(ctx, code, name, gin.H{"user": user, "data": data})
}
