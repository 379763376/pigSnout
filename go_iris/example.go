package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"gopkg.in/go-playground/validator.v9"
	"mime/multipart"
	"os"
	"regexp"
	"strings"
	"time"

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

//3、Parameters in path URL中带参数、内置函数校验、自定义函数校验
func main003() {
	//URL中带参数
	app := iris.Default()
	app.Get("/users/{id:uint64}", func(ctx iris.Context) {
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
	//内置函数校验
	app.Get("/profile/{name:alphabetical max(255)}", func(ctx iris.Context) {
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
	//自定义函数校验  func(string) bool
	latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
	latLonRegex, _ := regexp.Compile(latLonExpr)
	// Register your custom argument-less macro function to the :string param type.
	// MatchString is a type of func(string) bool, so we use it as it is.
	app.Macros().Get("string").RegisterFunc("coordinate", latLonRegex.MatchString)
	app.Get("/coordinates/{lat:string coordinate()}/{lon:string coordinate()}", func(ctx iris.Context) {
		ctx.Writef("Lat: %s | Lon: %s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
	})
	//自定义函数接受两个变量
	app.Macros().Get("string").RegisterFunc("range", func(minLength, maxLength int) func(string) bool {
		return func(paramValue string) bool {
			return len(paramValue) >= minLength && len(paramValue) <= maxLength
		}
	})
	app.Get("/limitchar/{name:string range(1,200) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length
    otherwise this handler will not be executed`, name)
	})
	//自定义函数接收slice
	app.Macros().Get("string").RegisterFunc("has", func(validNames []string) func(string) bool {
		return func(paramValue string) bool {
			for _, validName := range validNames {
				if validName == paramValue {
					return true
				}
			}
			return false
		}
	})

	app.Get("/static_validation/{name:string has([kataras,gerasimos,maropoulos])}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be "kataras" or "gerasimos" or "maropoulos"
    otherwise this handler will not be executed`, name)
	})

	// However, this one will match /user/john/send and also /user/john/everything/else/here
	// but will not match /user/john neither /user/john/.
	app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})
	//{name:string}和{name}相同的
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

//4、Querystring parameters 参数拼接在url的get请求
func main004() {
	app := iris.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe.
	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		// shortcut for ctx.Request().URL.Query().Get("lastname").
		lastname := ctx.URLParam("lastname")

		ctx.Writef("Hello %s %s", firstname, lastname)
	})

	app.Run(iris.Addr(":8080"))
}

//5、Multipart/Urlencoded Form post提交form表单
func main005() {
	app := iris.Default()

	app.Post("/form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick := ctx.FormValueDefault("nick", "anonymous")

		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	app.Run(iris.Addr(":8080"))
}

//6、Another example: query + post form 参数在url的post
/*
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
*/
func main006() {
	app := iris.Default()

	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		// or `ctx.PostValue` for POST, PUT & PATCH-only HTTP Methods.

		app.Logger().Infof("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	app.Run(iris.Addr(":8080"))
}

//7、提取refer
//curl http://localhost:8080?referer=https://twitter.com/Xinterio/status/1023566830974251008
//curl http://localhost:8080?referer=https://www.google.com/search?q=Top+6+golang+web+frameworks&oq=Top+6+golang+web+frameworks
func main007() {
	app := iris.New()

	app.Get("/", func(ctx context.Context) /* or iris.Context, it's the same for Go 1.9+. */ {

		// request header "referer" or url parameter "referer".
		r := ctx.GetReferrer()
		switch r.Type {
		case context.ReferrerSearch:
			ctx.Writef("Search %s: %s\n", r.Label, r.Query)
			ctx.Writef("Google: %s\n", r.GoogleType)
		case context.ReferrerSocial:
			ctx.Writef("Social %s\n", r.Label)
		case context.ReferrerIndirect:
			ctx.Writef("Indirect: %s\n", r.URL)
		}
	})

	app.Run(iris.Addr(":8080"))
}

//8、上传文件
//但文件上传
//多文件上传
const maxSize1 = 5 << 20 // 5MB
func main008() {
	app := iris.Default()
	app.Post("/upload", iris.LimitRequestBodySize(maxSize1), func(ctx iris.Context) {
		//
		// UploadFormFiles
		// uploads any number of incoming files ("multiple" property on the form input).
		//

		// The second, optional, argument
		// can be used to change a file's name based on the request,
		// at this example we will showcase how to use it
		// by prefixing the uploaded file with the current user's ip.
		ctx.UploadFormFiles("./uploads", beforeSave)
	})

	app.Run(iris.Addr(":8080"))
}
func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=	time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}

//9.Grouping routes 路由组
func main009() {
	app := iris.Default()

	// Simple group: v1.
	//v1 := app.Party("/v1")
	//{
	//	v1.Post("/login", loginEndpoint)
	//	v1.Post("/submit", submitEndpoint)
	//	v1.Post("/read", readEndpoint)
	//}

	// Simple group: v2.
	//v2 := app.Party("/v2")
	//{
	//	v2.Post("/login", loginEndpoint)
	//	v2.Post("/submit", submitEndpoint)
	//	v2.Post("/read", readEndpoint)
	//}

	app.Run(iris.Addr(":8080"))
}

//10.使用blank iris 而不是默认的default
//默认的default包含 Logger and Recovery middleware
//自定义中间件加入到app
//路由加中加入中间件,加入自定义中间件
//路由组加入中间件
func main10() {
	// Creates an application without any middleware by default.
	app := iris.New()

	// Recover middleware recovers from any panics and writes a 500 if there was one.
	app.Use(recover.New())

	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	// Per route middleware, you can add as many as you desire.
	//app.Get("/benchmark",/* MyBenchLogger(),*/ benchEndpoint)

	// Authorization party /user.
	// authorized := app.Party("/user", AuthRequired())
	// exactly the same as:
	//authorized := app.Party("/user")
	// per party middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group/party.
	//authorized.Use(AuthRequired())
	//{
	//authorized.Post("/login", loginEndpoint)
	//authorized.Post("/submit", submitEndpoint)
	//authorized.Post("/read", readEndpoint)
	//
	//// nested group: /user/testing
	//testing := authorized.Party("/testing")
	//testing.Get("/analytics", analyticsEndpoint)

	//}

	// Listen and serve on http://0.0.0.0:8080
	app.Run(iris.Addr(":8080"))
}

//11.程序日志文件
// Get a filename based on the date, just for the sugar.
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func main11() {
	f := newLogFile()
	defer f.Close()

	app := iris.New()
	// Attach the file as logger, remember, iris' app logger is just an io.Writer.
	// Use the following code if you need to write the logs to file and console at the same time.
	// app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	app.Logger().SetOutput(f)

	app.Get("/ping", func(ctx iris.Context) {
		// for the sake of simplicity, in order see the logs at the ./_today_.txt
		ctx.Application().Logger().Infof("Request path: %s", ctx.Path())
		ctx.WriteString("pong")
	})

	// Navigate to http://localhost:8080/ping
	// and open the ./logs{TODAY}.txt file.
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}

//12、模型绑定&校验
//使用go-playground/validator.v9
//eg 绑定到json json:"fieldname"
// User contains user information.
type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
	Age            uint8      `json:"age" validate:"gte=0,lte=130"`
	Email          string     `json:"email" validate:"required,email"`
	FavouriteColor string     `json:"favColor" validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `json:"addresses" validate:"required,dive,required"`
}

// Address houses a users address information.
type Address struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	Planet string `json:"planet" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}

// Use a single instance of Validate, it caches struct info.
var validate *validator.Validate

func main12() {
	validate = validator.New()

	// Register validation for 'User'
	// NOTE: only have to register a non-pointer type for 'User', validator
	// internally dereferences during it's type checks.
	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	app := iris.New()
	app.Post("/user", func(ctx iris.Context) {
		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			// Handle error.
		}

		// Returns InvalidValidationError for bad validation input,
		// nil or ValidationErrors ( []FieldError )
		err := validate.Struct(user)
		if err != nil {

			// This check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.
			if _, ok := err.(*validator.InvalidValidationError); ok {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.WriteString(err.Error())
				return
			}

			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println()
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			}

			return
		}
		fmt.Println("save user to database")
		// save user to database.
	})

	app.Run(iris.Addr(":8080"))
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}
}
//13、cookies
//http会话
//给指定的路由设置cookie
func newApp13() *iris.Application {
	app := iris.New()

	// Set A Cookie.
	app.Get("/cookies/{name}/{value}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		value := ctx.Params().Get("value")

		ctx.SetCookieKV(name, value)
		//ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))//指定路径
		//ctx.SetCookieKV(name, value, iris.CookieCleanPath /* or iris.CookiePath("") */)//只是当前路径
		//ctx.Request().Cookie(name)

		//iris.CookieExpires(time.Duration)//时效
		iris.CookieHTTPOnly(false)//只是http

		ctx.Writef("cookie added: %s = %s", name, value)
	})

	// Retrieve A Cookie.
	app.Get("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		value := ctx.GetCookie(name)

		ctx.WriteString(value)
	})

	// Delete A Cookie.
	app.Delete("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name)

		ctx.Writef("cookie %s removed", name)
	})

	return app
}

func main13() {
	app := newApp13()

	// GET:    http://localhost:8080/cookies/my_name/my_value
	// GET:    http://localhost:8080/cookies/my_name
	// DELETE: http://localhost:8080/cookies/my_name
	app.Run(iris.Addr(":8080"))
}
//14、web 测试框架 httpexpect
//包iris/httptest提供Iris + httpexpect支持
//15、Basic Authentication
//使用 iris/httptest 测试 Basic Authentication.
func newApp15() *iris.Application {
	app := iris.New()

	authConfig := basicauth.Config{
		Users: map[string]string{"myusername": "mypassword"},
	}

	authentication := basicauth.New(authConfig)

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", h)
		// http://localhost:8080/admin/profile
		needAuth.Get("/profile", h)

		// http://localhost:8080/admin/settings
		needAuth.Get("/settings", h)
	}

	return app
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	// third parameter  ^ will be always true because the middleware
	// makes sure for that, otherwise this handler will not be executed.

	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

func main() {
	app := newApp15()
	app.Run(iris.Addr(":8080"))
}
//16.go test -v
//17.测试cookies
//go test -v -run=TestCookiesBasic$