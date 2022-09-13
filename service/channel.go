/*
 * @Author: GG
 * @Date: 2022-09-13 14:37:08
 * @LastEditTime: 2022-09-13 15:00:55
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\service\channel.go
 *
 */
package service

import (
	"gin-blog/common/global"
	"gin-blog/models"
)

type ChannelService struct {
}

func (u ChannelService) AddChannel(channel models.Channel) int64 {
	return global.DB.Create(&channel).RowsAffected
}

func (u ChannelService) DelChannel(id int) int64 {
	return global.DB.Delete(&models.Channel{}, id).RowsAffected
}

func (u ChannelService) UpdateChannel(channel models.Channel) int64 {
	return global.DB.Updates(&channel).RowsAffected
}

func (u ChannelService) GetChannel(id int) models.Channel {
	var channel models.Channel
	global.DB.First(&channel, id)
	return channel
}

func (u ChannelService) GetChannelList() []models.Channel {
	var channels []models.Channel
	global.DB.Find(&channels)
	return channels
}
