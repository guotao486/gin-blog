/*
 * @Author: GG
 * @Date: 2022-09-20 16:48:40
 * @LastEditTime: 2022-09-21 09:54:57
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\middleware\session.go
 *
 */
package middleware

import (
	"gin-blog/common/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var sessionUtil utils.SessionUtils

// session 中间件
func Session(key string) gin.HandlerFunc {
	store := sessionUtil.SessionDefaultConfig()
	return sessions.Sessions(key, store)
}
