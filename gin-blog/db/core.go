package core

import (
	"Gorm"
  "gorm.io/driver/postgres"
	"Gorm"
	"GORM"
	"GORM"
)

// 定义db全局变量
var Db *gorm.DB

func init() {
	var err error
  dsn := "host=localhost user=admin password=88888888 dbname=blogs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger: logger.Default.LogMode(logger.Info),// 打印sql语句
		DisableAutomaticPing: false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		panic("Connecting database failed: " + err.Error())
	}
}

//GetDB
func GetDB() *gorm.DB {
	return Db
}
