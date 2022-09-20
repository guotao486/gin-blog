/*
 * @Author: GG
 * @Date: 2022-09-20 16:48:40
 * @LastEditTime: 2022-09-20 16:59:06
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\middleware\session.go
 *
 */
package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 配置
func sessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "gin_blog"
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path:   "/admin",
	})

	return store
}

// session 中间件
func Session(key string) gin.HandlerFunc {
	store := sessionConfig()
	return sessions.Sessions(key, store)
}
