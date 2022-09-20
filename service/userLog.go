/*
 * @Author: GG
 * @Date: 2022-09-20 16:34:11
 * @LastEditTime: 2022-09-20 16:45:53
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\service\userLog.go
 *
 */
package service

import (
	"gin-blog/common/global"
	"gin-blog/models"
)

type UserLogService struct{}

func (u *UserLogService) Add(user models.User) int64 {
	var userLog models.UserLog
	userLog.Uid = int(user.ID)
	return global.DB.Create(&userLog).RowsAffected
}

func (u *UserLogService) List() (list []models.UserLog) {
	return
}

func (u *UserLogService) View(id int) (userLog models.UserLog) {
	return
}
