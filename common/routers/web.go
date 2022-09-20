/*
 * @Author: GG
 * @Date: 2022-09-16 09:28:19
 * @LastEditTime: 2022-09-20 17:40:23
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\routers\web.go
 *
 */
package routers

import (
	"gin-blog/common/middleware"
	"gin-blog/controller"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

type WebRouter struct {
}

func (r *WebRouter) Router(engine *gin.Engine) {
	// session 中间件
	engine.Use(middleware.Session("admin"))
	// 静态资源路径
	engine.Static("/assets", "./assets")

	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})
	engine.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "404",
		})
	})
	engine.GET("/", controller.Index)
	engine.GET("/post/:id", controller.DetailPost)

	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:         "templates/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})

	admin := engine.Group("/admin", mw)
	{
		admin.GET("/login", controller.Login)
		admin.POST("/login", controller.Login, middleware.UserLog())
		admin.GET("/regiest", controller.Regiest)
		admin.GET("/logout", controller.Logout)
	}
	{
		// 权限判断
		admin.Use(middleware.IsLogin())
		//index
		admin.GET("/", controller.AdminIndex)
	}
	{
		admin.GET("/channel/list", controller.ListChannel)
		admin.GET("/channel/view", controller.ViewChannel)
		admin.DELETE("/channel/delete", controller.DeleteChannel)
		admin.POST("/channel/save", controller.SaveChannel)
	}

	{
		admin.GET("/post/list", controller.ListPost)
		admin.GET("/post/view", controller.ViewPost)
		admin.DELETE("/post/delete", controller.DeletePost)
		admin.POST("/post/save", controller.SavePost)

		admin.POST("/post/upload", controller.UploadThumbnails)
	}

}
