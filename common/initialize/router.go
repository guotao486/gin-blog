/*
 * @Author: GG
 * @Date: 2022-09-13 16:25:05
 * @LastEditTime: 2022-09-15 17:18:02
 * @LastEditors: GG
 * @Description: router 路由配置
 * @FilePath: \gin-blog\common\initialize\router.go
 *
 */
package initialize

import (
	"fmt"
	"gin-blog/common/global"
	"gin-blog/controller"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func Router() {
	engine := gin.Default()

	// 静态资源路径
	engine.Static("/assets", "./assets")

	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
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
	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
