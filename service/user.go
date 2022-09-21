/*
 * @Author: GG
 * @Date: 2022-09-20 16:10:22
 * @LastEditTime: 2022-09-21 17:52:59
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\service\user.go
 *
 */
package service

import (
	"gin-blog/common/global"
	"gin-blog/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (u *UserService) GetUserByUsername(username string) (user models.User, exist bool) {
	exist = true
	r := global.DB.Debug().Where("username = ?", username).First(&user)
	if r.RowsAffected == 0 {
		exist = false
	}
	return
}

// 获取加密密码
func (u *UserService) GenPwd(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)
}

// 密码比较
func (u *UserService) ComparePwd(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	} else {
		return true
	}
}
