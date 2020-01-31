package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
标准库
encoding/json
encoding/xml
encoding/ans1
google的Protocol Buffers 库github.com/golang/protobuf

json可以有效表示  基础数据类型 数组 slice 结构体 map等聚合数据类型
*/
/*
	"released": 1967,
	"color": true,
 */
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}
func main() {
	data,error := json.Marshal(movies)// 输出

	if error != nil{
		log.Fatal("JSON marshaling failed: %s", error)
	}
	fmt.Printf("%s\n",data)
	data2,error2 := json.MarshalIndent(movies,"","	")
	if error2 != nil{
		log.Fatal("JSON marshaling failed: %s", error2)
	}

	fmt.Printf("%s\n",data2)
}
