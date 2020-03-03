package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1.优雅关停
	//httpServer构建一个server 不再使用Run
	//收到os.Signal会把这个信号传导shutdown方法里面
	//设置超时时间 关闭这个时间段进来的请求 处理接收到的请求
	r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {

	})
	srv ：= &http.Server{
		Addr:              "",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
}
