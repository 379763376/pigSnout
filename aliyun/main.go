package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Payload struct {
	page       int
	size       int
	categoryId string
	orderNum   string
	_csrf      string
}

func main() {
	url := "https://www.ica-alliance.org/common/resource/download?ownerType=3&id="
	seen := make([]int,50)
	InitTempIds()
	for _, tempType := range TempTypes {
		if check, err := isFileExists("./aliyun-data-"+tempType); check == false && err == nil { //判断不存在
			if err := os.Mkdir("./aliyun-data-"+tempType,0777); err != nil { //创建失败
				fmt.Println("创建失败", err)
				return
			}
		} else if check == false && err != nil { // 函数调用错误
			fmt.Println("检测目录出错", err)
			return
		}

		for _,id := range TempIds[tempType] {
			if isExistSlice(id,seen){
				continue
			}
			seen = append(seen, id)
			resp,err := http.Get(url+strconv.Itoa(id))
			if err != nil{
				fmt.Fprintf(os.Stderr, "获取错误: %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()
			f,err := os.Create("./aliyun-data-"+tempType+"/"+strconv.Itoa(id))
			if err != nil {
				fmt.Println("创建文件失败：",id)
				return
			}
			if _,err = io.Copy(f,resp.Body);err != nil {
				fmt.Println("复制文件出错：", err)
			}
			if err := f.Close(); err != nil {
				fmt.Println("关闭文件出差：", err)
			}
		}


	}


}

func isFileExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func isExistSlice(id int, slise []int) bool  {
	for _,v := range slise {
		if id == v {
			return true
		}
	}
	return false
}