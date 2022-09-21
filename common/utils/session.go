/*
 * @Author: GG
 * @Date: 2022-09-21 09:45:49
 * @LastEditTime: 2022-09-21 10:54:30
 * @LastEditors: GG
 * @Description: session util
 * @FilePath: \gin-blog\common\utils\session.go
 *
 */
package utils

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type SessionUtils struct{}

const MAX_AGE = 60 * 60               // 60 * 60 1hour
const MAX_AGE_REMEMBER = 60 * 60 * 24 // 60 * 60 * 24 1day

// session 中间件 默认配置
func (s *SessionUtils) SessionDefaultConfig() sessions.Store {
	sessionMaxAge := MAX_AGE
	sessionSecret := "gin_blog"
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path:   "/admin",
	})

	return store
}

// 修改sessionMaxAge
func (s *SessionUtils) SetMaxAge(session sessions.Session, maxage int) sessions.Session {
	session.Options(sessions.Options{
		MaxAge: maxage,
	})
	return session
}

// 存储session
func (s *SessionUtils) Set(ctx *gin.Context, name string, data interface{}, structBool bool, maxage int) {
	if structBool {
		gob.Register(data)
	}
	session := sessions.Default(ctx)

	if maxage != 0 {
		session.Options(sessions.Options{
			MaxAge: maxage,
		})
	}
	session.Set(name, data)
	session.Save()
}
