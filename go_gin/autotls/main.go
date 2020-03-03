package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
)

func main() {
	//自动配置证书
	//调用一下证书下载的名字即可
	r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200,"hello test")
	})

	//调用证书下载的包
	//流程autotls.Run先生成一个本地的密钥 把这个密钥发给一个证书颁发机构，获取一个私钥
	//拿到私钥 本地对私钥验证 没问题就保存起来，下次请求的时候用私钥进行加密
	autotls.Run(r, "www.itpp.tk")
}
