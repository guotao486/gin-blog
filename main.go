/*
 * @Author: GG
 * @Date: 2022-09-13 16:47:36
 * @LastEditTime: 2022-09-20 16:27:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\main.go
 *
 */
package main

import (
	"fmt"
	"gin-blog/common/initialize"
)

func main() {
	fmt.Println("hello!")
	initialize.LoadConfig()
	initialize.Mysql()
	// global.DB.AutoMigrate(&models.User{}, &models.UserLog{})
	initialize.Router()
}
