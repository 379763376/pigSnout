package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer //接口
	w = os.Stdout //w为*File类型,具体实现类型
	fmt.Printf("w++++T:%T,V:%V,addr:%v\n",w,w,&w)//T:*os.File
	f := w.(*os.File) //f接口值 T:os.Stdout V:
	fmt.Printf("w++++T:%T,V:%V,addr:%v\n",w,w,&w)//T:*os.File
	fmt.Printf("f++++T:%T,V:%V,addr:%v\n",f,f,&f)//T:*os.File


	rw := w.(io.ReadWriter)//断言类型为接口，w为具体实现，rw的类型为
	fmt.Printf("rw++++T:%T,V:%V,addr:%v\n",rw,rw,&rw)//T:*os.File

	var rw2  io.ReadWriter //rw的接口值是nil T:nil V:nil
	w=rw2
	fmt.Printf("rw2++++T:%T,V:%V,addr:%v\n",rw2,rw2,&rw2)//nil
	fmt.Printf("w++++T:%T,V:%V,addr:%v\n",w,w,&w)//nil
	//w = rw2.(io.Writer) //panic 断言操作的对象是一个nil接口值
	//fmt.Printf("rw2++++T:%T,V:%V,addr:%v\n",rw2,rw2,&rw2)

	w=os.Stdout
	b,ok := w.(*bytes.Buffer)
	if !ok{
		fmt.Printf("%V\n", b)
	}
	if c,ok := w.(*os.File);ok{
		fmt.Printf("%V\n",c)
	}
}
