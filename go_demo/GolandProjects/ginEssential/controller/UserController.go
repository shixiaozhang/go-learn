package controller

import (
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB:=common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证 手机号11位、密文至少6位
	if len(telephone) != 11 {
		// 修改响应代码
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB,telephone){
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "该号码已注册",
		})
		return
	}
	// 创建用户
	// 密码用密文存入数据库
	hasedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"加密错误"})
		return
	}

	newUser:=model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(200, gin.H{
		"code":200,
		"msg": "注册成功",
	})
	//response.Success(ctx,nil,"注册成功")
}

func Login(ctx *gin.Context) {
	DB:=common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone=?",telephone).First(&user)
	if user.ID==0{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号不存在",
		})
		return
	}
	// 判断密码是否正确
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	// 发放token
	token,err:=common.ReleaseToken(user)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		log.Printf("token generate error : %v",err)
		return
	}
	// 返回结果
	ctx.JSON(200,gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"msg":"登录成功",
	})
}

func Info(ctx *gin.Context)  {
	// 从上下文中获取用户信息
	user,_:=ctx.Get("user")
	// interface{}类型转换 （接口断言）类型断言
	ctx.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user":dto.ToUserDto(user.(model.User))}})

}


// https://golang.org/pkg/net/rpc

// 判断手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID!=0{
		return true
	}
	return false
}