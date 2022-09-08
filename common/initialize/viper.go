/*
 * @Author: GG
 * @Date: 2022-09-08 15:04:25
 * @LastEditTime: 2022-09-08 15:22:24
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\initialize\viper.go
 *
 */
package initialize

import (
	"fmt"
	"gin-blog/common/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error resources file: %w \n", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct %w \n", err))
	}
}
