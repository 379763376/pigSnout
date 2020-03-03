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
	r.Static("/assets","./assets")
	r.StaticFS("/static")
	r.Run()
}
