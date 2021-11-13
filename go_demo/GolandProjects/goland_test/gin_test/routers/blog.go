package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBlog(e *gin.Engine) {
	e.GET("/post", postHandler)
	e.GET("/comment", commentHandler)
}
func postHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello blog!")
}
func commentHandler(c *gin.Context) {}


