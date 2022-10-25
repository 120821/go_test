package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
//使用open函数连接postgresql数据库，第一个参数填数据库类型，后面的参数分别是数据库地址，端口，数据库名称，密码，是否ssl模式连接，时区
	dsn := "host=127.0.0.1 port=5432 user=admin dbname=blogs password=88888888 sslmode=disable TimeZone=Asia/Shanghai"
  //连接数据库
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=db1 password=123456 sslmode=disable")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("连接成功！")
	}
	defer db.Close()
}

