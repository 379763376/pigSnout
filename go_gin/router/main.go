package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	r:=gin.Default()
	//r.Static("/assets","./assets")
	//r.StaticFS("/static",http.Dir("static"))
	//r.StaticFile("favicon.ico","./favicon.ico")

	//4.请求路由 - 参数作为url
	//r.GET("/:name/:id", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"name":c.Param("name"),
	//		"id":c.Param("id"),
	//	})
	//})

	//5.请求路由 - 泛绑定
	//r.GET("/user/*action", func(c *gin.Context) {
	//	c.String(200,"hello world")
	//})
	r.Run()
}
