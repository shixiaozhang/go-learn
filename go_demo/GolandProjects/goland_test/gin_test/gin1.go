package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	// 绑定路由规则
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		// gin.H 等同于 map[string]interface{}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go test()
	r.Run() // listen and serve on 0.0.0.0:8080

}

type User struct {
	ID   uint64
	Name string
}


func test() {
	fmt.Println("test")
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	// 普通get请求
	r.GET("/get", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})
	// 返回结构体json
	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	// 普通post请求
	r.POST("/post",func(context *gin.Context) {
		context.String(http.StatusOK, "POST!")
	})
	// api参数 :匹配一个 *匹配所有
	r.GET("user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")

		fmt.Println(action)
		//  截取/
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})
	// url参数
	r.GET("user2", func(context *gin.Context) { //指定默认值 //http://127.0.0.1/user
		name := context.DefaultQuery("name", "you")
		// c.String(200, c.Query("name"))
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name)) })
	// html 表单
	r.POST("/form", func(context *gin.Context) {
		types := context.DefaultPostForm("type", "post")
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, password, types))
	})
	// 单文件上传
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload1", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		context.SaveUploadedFile(file, file.Filename)
		context.String(http.StatusOK, file.Filename)
	})

	// 多文件上传
	r.POST("/upload2", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		//获取所有文件
		files := form.File["files"]
		//遍历所有文件
		for _, file := range files {
			if err := context.SaveUploadedFile(file, file.Filename); err != nil {
				context.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
	})
	// 创建路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	// 创建多个路由文件
	// 引入静态文件
	r.Static("/static", "./gin_test")
	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	//3.监听端口，默认8080
	r.Run(":8081")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}