/*
 * @Author: GG
 * @Date: 2022-09-08 15:03:53
 * @LastEditTime: 2022-09-13 17:06:31
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\config\config.go
 *
 */
package config

type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructrue:"mysql"`
}
type Server struct {
	Post string `mapstructure:"post"`
}
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}
