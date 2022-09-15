/*
 * @Author: GG
 * @Date: 2022-09-15 11:42:14
 * @LastEditTime: 2022-09-15 16:49:39
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\controller\post.go
 *
 */
package controller

import (
	"fmt"
	"gin-blog/common/stripmd"
	"gin-blog/models"
	"gin-blog/service"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

var postService service.PostService
var channelService service.ChannelService

func ListPost(c *gin.Context) {
	plist := postService.GetPostList()
	gintemplate.HTML(c, http.StatusOK, "post/list", gin.H{"plist": plist})
}

func DetailPost(c *gin.Context) {
	sid := c.Param("id")
	id := cast.ToInt(sid)
	post := postService.GetPost(id)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	md := []byte(post.Content)
	md = markdown.NormalizeNewlines(md)
	html := markdown.ToHTML(md, parser, nil)

	gintemplate.HTML(c, http.StatusOK, "post/detail", gin.H{
		"post":    post,
		"content": template.HTML(html),
	})
}

// 后台
func ViewPost(c *gin.Context) {
	sid, r := c.GetQuery("id")
	var post models.Post
	if r {
		id, _ := strconv.Atoi(sid)
		post = postService.GetPost(id)
	}
	clist := channelService.GetChannelList()
	gintemplate.HTML(c, http.StatusOK, "post/view", gin.H{
		"post":  post,
		"clist": clist,
	})
}

func DeletePost(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	postService.DelPost(id)
	c.Redirect(http.StatusFound, "/admin/post/list")
}

// 上传封面
func UploadThumbnails(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "请上传文件!"})
		return
	}

	// 获取文件扩展名
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.NewString() + extension

	pwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/static/thumbnails/%s", pwd, newFileName)
	relativeFilePath := fmt.Sprintf("/static/thumbnails/%s", newFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		fmt.Printf("err: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messsage": "文件保存失败"})
		return
	}

	c.String(http.StatusOK, relativeFilePath)
}

func SavePost(c *gin.Context) {
	var post models.Post

	title := c.PostForm("title")
	thumbnail := c.PostForm("thumbnail")
	tags := c.PostForm("tags")
	content := c.PostForm("content")
	channelId := c.PostForm("channelId")

	summary := stripmd.Strip(content)
	l := len(summary)
	if l > 200 {
		summary = summary[0:200]
	} else {
		summary = summary[0:l]
	}

	post.Title = title
	post.Thumbnail = thumbnail
	post.Tags = tags
	post.Content = content
	post.Summary = summary
	post.ChannelId = cast.ToInt(channelId)
	post.AuthorId = 1

	if err := c.ShouldBind(&post); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	id := cast.ToInt(c.PostForm("id"))

	if id != 0 {
		postService.UpdatePost(post)
	} else {
		postService.AddPost(post)
	}
	c.Redirect(http.StatusFound, "/admin/post/list")
}
