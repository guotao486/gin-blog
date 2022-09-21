/*
 * @Author: GG
 * @Date: 2022-09-21 09:45:49
 * @LastEditTime: 2022-09-21 10:42:02
 * @LastEditors: GG
 * @Description: session util
 * @FilePath: \gin-blog\common\utils\session.go
 *
 */
package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

type SessionUtils struct{}

const MAX_AGE = 60 * 60               // 60 * 60 1hour
const MAX_AGE_REMEMBER = 60 * 60 * 24 // 60 * 60 * 24 1day

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
