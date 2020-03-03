package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	r:=gin.Default()
	//1.测试请求类型
	//r.GET("/get", func(c *gin.Context) {
	//	c.String(200,"get")
	//})
	//r.POST("/post", func(c *gin.Context) {
	//	c.String(200,"post")
	//})
	//r.DELETE("/delete", func(c *gin.Context) {
	//	c.String(200,"delete")
	//})
	//r.Handle("DELETE","/delete2", func(c *gin.Context) {
	//	c.String(200,"delete")
	//})
	//r.Any("/any", func(c *gin.Context) {
	//	c.String(200,"any")
	//})
	//2.测试静态文件夹
	//r.Static("/assets","./assets")
	//r.StaticFS("/static",http.Dir("static"))
	//r.StaticFile("favicon.ico","./favicon.ico")

	//3.请求路由 - 参数作为url
	//r.GET("/:name/:id", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"name":c.Param("name"),
	//		"id":c.Param("id"),
	//	})
	//})

	//4.请求路由 - 泛绑定
	//r.GET("/user/*action", func(c *gin.Context) {
	//	c.String(200,"hello world")
	//})
	r.Run()
}
