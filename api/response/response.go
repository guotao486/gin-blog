/*
 * @Author: GG
 * @Date: 2022-09-16 10:22:53
 * @LastEditTime: 2022-09-16 10:23:17
 * @LastEditors: GG
 * @Description: api response
 * @FilePath: \gin-blog\api\response\response.go
 *
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// api返回数据对象结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 请求成功返回
func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}
