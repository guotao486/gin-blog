/*
 * @Author: GG
 * @Date: 2022-09-13 16:25:05
 * @LastEditTime: 2022-09-16 10:02:47
 * @LastEditors: GG
 * @Description: router 路由配置
 * @FilePath: \gin-blog\common\initialize\router.go
 *
 */
package initialize

import (
	"fmt"
	"gin-blog/common/global"
	"gin-blog/common/routers"

	"github.com/gin-gonic/gin"
)

var webRouter routers.WebRouter
var apiRouter routers.ApiRouter

func Router() {
	engine := gin.Default()
	webRouter.Router(engine)
	apiRouter.Router(engine)

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
