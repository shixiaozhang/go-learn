package common

import (
	"fmt"
	"ginEssential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 连接mysql
func InitDB() *gorm.DB {
	//driverName := "mysql"
	//driverName:=viper.GetString("datasource.driverName")
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := ""
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args),&gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	// 自动创建数据表
	db.AutoMigrate(&model.User{})
	DB=db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

