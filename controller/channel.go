/*
 * @Author: GG
 * @Date: 2022-09-13 15:27:49
 * @LastEditTime: 2022-09-13 17:45:51
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\controller\channel.go
 *
 */
package controller

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/service"
	"net/http"
	"strconv"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "index", gin.H{})
}
func AdminIndex(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "index", gin.H{})
}

var channel service.ChannelService

func ListChannel(c *gin.Context) {
	list := channel.GetChannelList()
	gintemplate.HTML(c, http.StatusOK, "channel/list", list)
}

func ViewChannel(c *gin.Context) {
	sid, r := c.GetQuery("id")
	var detail models.Channel
	if r {
		id, _ := strconv.Atoi(sid)
		detail = channel.GetChannel(id)
	}
	gintemplate.HTML(c, http.StatusOK, "channel/view", detail)
}

func DeleteChannel(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	channel.DelChannel(id)
	c.Redirect(http.StatusFound, "/admin/channel/list")
}

func SaveChannel(c *gin.Context) {
	var chann models.Channel
	if err := c.ShouldBind(&chann); err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	chann.Status, _ = strconv.Atoi(c.PostForm("status"))

	id := c.PostForm("id")
	if id != "0" {
		channel.UpdateChannel(chann)
	} else {
		channel.AddChannel(chann)
	}
	c.Redirect(http.StatusFound, "/admin/channel/list")
}
