/*
 * @Author: GG
 * @Date: 2022-09-20 15:38:44
 * @LastEditTime: 2022-09-20 15:40:23
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\models\user.go
 *
 */
package models

import "gorm.io/gorm"

type User struct {
	Username string
	Password string

	gorm.Model
}

type UserLog struct {
	Uid int
	gorm.Model
}
