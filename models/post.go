/*
 * @Author: GG
 * @Date: 2022-09-15 10:37:45
 * @LastEditTime: 2022-09-15 10:40:38
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\models\post.go
 *
 */
package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string `gorm:"slug"`
	Thumbnail string `gorm:"thumbnail"`
	Summary   string `gorm:"summary"`
	Content   string `gorm:"content"`
	AuthorId  int    `gorm:"author_id"`
	ChannelId int    `gorm:"channel_id"`
	Comments  int    `gorm:"comments"`
	Favors    int    `gorm:"favors"`
	Featured  int    `gorm:"featured"`
	Status    int    `gorm:"status"`
	Tags      string `gorm:"tags"`
	Views     int    `gorm:"views"`
	Weight    int    `gorm:"weight"`
	Url       string `gorm:"url"`
}
