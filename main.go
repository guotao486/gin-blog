/*
 * @Author: GG
 * @Date: 2022-09-13 16:47:36
 * @LastEditTime: 2022-09-15 17:17:50
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
	// global.DB.AutoMigrate(&models.Post{})
	initialize.Router()
}
