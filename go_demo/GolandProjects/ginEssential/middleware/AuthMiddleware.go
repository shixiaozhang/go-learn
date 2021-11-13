package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


// gin的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString:=ctx.GetHeader("Authorization")

		// validate token formate
		// 如果token为空或者不是以Bearer开头
		if tokenString==""||!strings.HasPrefix(tokenString,"Bearer"){
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()	// 将这次请求抛弃掉，并返回
			return
		}
		// 除去Bearer+空格的有效部分
		tokenString=tokenString[7:]
		// 解析token
		token,claims,err:=common.ParseToken(tokenString)
		// 如果解析失败或token无效
		if err!=nil||!token.Valid{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		// 验证通过后获取claim中的userid
		userId:=claims.UserId
		DB:=common.GetDB()
		var user model.User
		DB.First(&user,userId)

		// 用户
		if user.ID==0{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		// 用户存在 将user的信息写入上下文
		ctx.Set("user",user)
		ctx.Next()

	}
}