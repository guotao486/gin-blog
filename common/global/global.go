/*
 * @Author: GG
 * @Date: 2022-09-08 15:12:21
 * @LastEditTime: 2022-09-08 15:13:12
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\global\global.go
 *
 */
package global

import (
	"gin-blog/common/config"

	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
)
