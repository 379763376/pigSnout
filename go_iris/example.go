package main

import (
	"fmt"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

//1、sample
func main001() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
//2、使用不同的请求Get, Post, Put, Patch, Delete and Options
func main002() {
	// Creates an application with default middleware:
	// logger and recovery (crash-free) middleware.
	app := iris.Default()

	//app.Get("/someGet", getting)
	//app.Post("/somePost", posting)
	//app.Put("/somePut", putting)
	//app.Delete("/someDelete", deleting)
	//app.Patch("/somePatch", patching)
	//app.Head("/someHead", head)
	//app.Options("/someOptions", options)

	app.Run(iris.Addr(":8080"))
}
//3、URL中带参数、带函数校验
func main003() {
	app := iris.Default()
	app.Get("/users/{id:uint64}", func(ctx iris.Context){
		id := ctx.Params().GetUint64Default("id", 0)
		fmt.Println(id)
		// [...]
	})
	/*
	Params().Get  string
	Params().GetInt int 8 16 32 64
	Params().GetUint uint  8 16 32 64
	Params().GetBool bool	“1” or “t” or “T” or “TRUE” or “true” or “True” or “0” or “f” or “F” or “FALSE” or “false” or “False”
	:alphabetical	Params().Get string	小写或大写字母
	:file  Params().Get string  小写、大写字母、数字、underscore (_), dash (-), point (.)
	:path  Params().Get string  放在请求路由的最后  可以用斜线隔开的任何string
	*/
	app.Get("/profile/{name:alphabetical max(255)}", func(ctx iris.Context){
		name := ctx.Params().Get("name")
		// len(name) <=255 otherwise this route will fire 404 Not Found
		// and this handler will not be executed at all.
		fmt.Println(name)
	})
	/*
	regexp(expr string)
	prefix(prefix string)
	suffix(suffix string)
	contains(s string)
	min(minValue int、uint 8...、flat 32|64)
	max(...)
	range(min,max)
	*/
}