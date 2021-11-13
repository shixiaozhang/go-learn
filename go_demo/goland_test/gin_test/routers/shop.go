package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadShop(e *gin.Engine)  {
	e.GET("/shop/hello", helloHandler)
	e.GET("/shop/goods", goodsHandler)
	e.GET("/shop/checkout", checkoutHandler)
}
func helloHandler(c *gin.Context)  {
	c.String(http.StatusOK, "Hello shop!")
}
func goodsHandler(c*gin.Context)  {}
func checkoutHandler(c*gin.Context)  {}