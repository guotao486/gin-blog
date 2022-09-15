/*
 * @Author: GG
 * @Date: 2022-09-15 10:44:28
 * @LastEditTime: 2022-09-15 11:05:10
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\service\post.go
 *
 */
package service

import (
	"gin-blog/common/global"
	"gin-blog/models"
)

type PostService struct {
}

func (p *PostService) AddPost(post models.Post) int64 {
	return global.DB.Create(&post).RowsAffected
}

func (p *PostService) DelPost(id int) int64 {
	return global.DB.Delete(&models.Post{}, id).RowsAffected
}

func (p *PostService) UpdatePost(post models.Post) int64 {
	return global.DB.Updates(&post).RowsAffected
}

func (p *PostService) GetPostList() []models.Post {
	var postList []models.Post
	global.DB.Find(&postList)
	return postList
}

func (p *PostService) GetPost(id int) (post models.Post) {
	global.DB.First(&post, id)
	return post
}
