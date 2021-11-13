package common

import (
	"ginEssential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义秘钥
var jwtKey=[]byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}
// 登录后发放token
func ReleaseToken(user model.User) (string,error) {
	// 设置token的有效时间
	expirationTime:=time.Now().Add(7*24*time.Hour)
	claims:=&Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "hello",// tkoen发放者
			Subject: "user token",// token主题
		},
	}
	// 选择加密方式
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(jwtKey)
	if err!=nil{
		return "",err
	}
	return tokenString,nil
}

func ParseToken(tokenString string)(*jwt.Token,*Claims,error){
	claims:=&Claims{}
	
	token,err :=jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey,nil
	})
	return token,claims,err
}

