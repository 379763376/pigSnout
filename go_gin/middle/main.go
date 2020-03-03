package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//1.gin 中间件 default默认会实现两个中间件
	//1.1设置log输出路径
	//f,_ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter = io.MultiWriter(f)
	//r := gin.New()
	//r.Use(gin.Logger(),gin.Recovery()) //gin.Logger()请求的时候会打印url耗时
	//1.2设置recovery 如果没有 这个报错会传到main 整个服务宕掉
	//r.GET("/test", func(c *gin.Context) {
	//	name := c.DefaultQuery("name","d_n")
	//	panic("test panic")
	//	c.String(200,"%s",name)
	//})
	//2.自定义中间件 Ip白名单
	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name","d_n")
		c.String(200,"%s",name)
	})

	r.Run()
}

func IPAuthMiddleware()  gin.HandlerFunc  {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP()
		for _,host := range ipList{
			if clientIP == host {
				flag = true
				break
			}
		}
		if !flag{
			c.String(401,"%s, not in iplist", clientIP)
			c.Abort()
		}
	}
}