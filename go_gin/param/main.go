package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

type Person struct {
	Name string `form:"name" binding:"required`
	Addr string `form:"addr"`
	Birthday time.Time `form:"birthday"  binding:"required,birthBefore" time_format:"2006-01-02"`
	Age int `form:"age" binding:"required,gt=10"`
	Tel int `form:"tel" binding:"required|gt=10"`
	Study time.Time `form:"study" binding:"required,gtfield=Birthday" time_format:"2006-01-02"`
}
func main() {
	r := gin.Default()
	//1.获取Get 参数
	//r.GET("/test", func(c *gin.Context) {
	//	name := c.Query("name")
	//	id := c.DefaultQuery("id","001")
	//	c.String(http.StatusOK,"%s,%s", name, id)
	//})
	//2.获取body请求参数
	//r.POST("/test", func(c *gin.Context) {
	//	//2.1使用ReadAll 之后c中的数据无法使用PostForm获取,需要回存
	//	//bts,err := ioutil.ReadAll(c.Request.Body)
	//	//if err != nil{
	//	//	c.String(http.StatusBadRequest,err.Error())
	//	//	c.Abort()
	//	//}
	//	//c.String(http.StatusOK,string(bts))
	//	//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bts))
	//	id := c.PostForm("id")
	//	name := c.DefaultPostForm("name","forever")
	//	c.String(http.StatusOK,"%s,%s", id, name)
	//})
	//3.参数绑定到结构体
	//r.GET("/test",testing)
	//r.POST("/test",testing)
	//4.验证请求参数
	//4.1 结构体验证 https://gopkg.in/go-playground/validator.v8 提供的参考规则
	//r.GET("/test", func(c *gin.Context) {
	//	var person Person
	//	if err:=c.ShouldBind(&person);err !=nil{
	//		c.String(http.StatusInternalServerError,"%v", err)
	//	}
	//	c.String(http.StatusOK,"%v", person)
	//})
	//4.2 自定义验证
	//validatorV8支持自定义结构体
	//步骤：定义binding标签、定义验证方法、方法注册到alidator验证器
	if v,ok :=binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("birthBefore",birthBefore)
	}else{
		fmt.Println("error")
	}
	r.GET("/test", func(c *gin.Context) {
		var person Person
		if err:=c.ShouldBind(&person);err !=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{"message":"OK!","person":person})
	})

	r.Run(":8081")
}

func testing(c *gin.Context) {
	var person *Person

	//会根据请求类型的content-type,binding不同的操作
	//Json需要在Content-Type:application/json指定
	//curl -H "Content-Type:application/json" -x POST "http://localhost:8080/test" -d '{"name":"wang"}'
	//curl -X POST 'localhost:8081/test' -d  'name=zs&addr=wh&birthday=2001-01-01'
	// curl -X GET 'localhost:8081/test?name=zs&addr=wh&birthday=2001-01-01'
	if err := c.ShouldBind(&person); err == nil{
		c.String(http.StatusOK,"%v", person)
		c.Abort()
		return
	}else{
		c.String(200, "person bind error:%v", err)
	}
}
func birthBefore(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date,ok := field.Interface().(time.Time);ok{
		today := time.Now()
		if date.Unix() <  today.Unix(){
			return true
		}
	}
	return false
}

