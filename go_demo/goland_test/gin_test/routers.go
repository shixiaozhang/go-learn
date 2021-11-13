package main


// 导包 前面加名字为别名，
//加_为只引入，执行里面的init方法
//加.为导入所有，可直接用里面的函数
import (
	"fmt"
	gin"github.com/gin-gonic/gin"
	"goland_test/gin_test/routers"
)
// 路由拆分多个文件
func main() {
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(":3456"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}