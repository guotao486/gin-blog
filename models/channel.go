/*
 * @Author: GG
 * @Date: 2022-09-09 14:56:20
 * @LastEditTime: 2022-09-13 17:27:13
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\models\channel.go
 *
 */
package models

type Channel struct {
	Id      uint64 `form:"id" gorm:"primaryKey"`
	Title   string `form:"title" gorm:"title"`
	Slug    string `form:"slug" gorm:"slug"`
	Content string `form:"content" gorm:"content"`
	Status  int    `form:"status" gorm:"status"`
	Weight  int    `form:"weight" grom:"weight"`
}
