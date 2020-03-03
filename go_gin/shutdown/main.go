package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//1.优雅关停
	//gin.Run是阻塞方式
	//http.Server构建一个server,server.ListenAndServe执行 是不阻塞的
	//os.Signal阻塞进程，监听关闭信号 ，获取到信号 把超时的上下文传到server.shutdown
	//从收到信号到关闭进程这段时间 拒绝链接请求 处理接收到的请求
	r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(10*time.Second)
		c.String(200,"hello test")
	})
	srv := &http.Server{Addr:":8085",Handler:r}
	go func() {
		if err:=srv.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{
			log.Fatal("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT, //ctrl+c
	syscall.SIGTERM,//kill -15
	//kill -9 无法捕获
	)
	<-quit
	log.Printf("shutdown server ...")

	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx);err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting")
}
