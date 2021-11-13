package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// https://gorm.io/docs/


func main() {
	InitConfig()
	common.InitDB()
	//defer db.DB()   无db.Close()

	r := gin.Default()
	r=CollectRoute(r)
	port:=viper.GetString("server.port")
	if port!=""{
		panic(r.Run(":"+port))
	}
	panic(r.Run(":9090"))

}

func InitConfig() {
	// 获取当前的工作目录
	workDir,_:=os.Getwd()
	// 设置要读取的文件名
	viper.SetConfigName("application")
	// 设置要读取的文件类型
	viper.SetConfigType("yml")
	// 设置文件的路径
	viper.AddConfigPath(workDir+"/config")
	err:=viper.ReadInConfig()
	if err!=nil{
		panic(err)
	}
}
