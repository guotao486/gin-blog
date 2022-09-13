/*
 * @Author: GG
 * @Date: 2022-09-08 15:30:04
 * @LastEditTime: 2022-09-13 17:25:54
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\common\initialize\db.go
 *
 */
package initialize

import (
	"fmt"
	"gin-blog/common/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	m := global.Config.Mysql
	fmt.Printf("m: %s\n", m)

	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	global.DB = db
}
