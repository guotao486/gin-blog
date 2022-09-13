/*
 * @Author: GG
 * @Date: 2022-09-13 16:47:36
 * @LastEditTime: 2022-09-13 17:27:33
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\main.go
 *
 */
package main

import (
	"fmt"
	"gin-blog/common/global"
	"gin-blog/common/initialize"
	"gin-blog/models"
)

func main() {
	fmt.Println("hello!")
	initialize.LoadConfig()
	initialize.Mysql()
	global.DB.AutoMigrate(&models.Channel{})
	initialize.Router()
}
