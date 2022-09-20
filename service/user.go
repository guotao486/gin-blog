/*
 * @Author: GG
 * @Date: 2022-09-20 16:10:22
 * @LastEditTime: 2022-09-20 17:36:07
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\service\user.go
 *
 */
package service

import (
	"gin-blog/common/global"
	"gin-blog/models"
)

type UserService struct{}

func (u *UserService) Login(username string, password string) (user models.User, exist bool) {
	exist = true
	r := global.DB.Debug().Where("username = ?", username).Where("password =?", password).First(&user)
	if r.RowsAffected == 0 {
		exist = false
	}
	return
}
