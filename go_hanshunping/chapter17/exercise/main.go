<<<<<<< HEAD
package main
import (
	"fmt"
	"reflect"
)
func main() {
	var str string = "tom"   //ok
	fs := reflect.ValueOf(&str) //ok fs -> string
	fs.Elem().SetString("jack") //ok
	fmt.Printf("%v\n", str) // jack
}
=======
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "tom"      //ok
	fs := reflect.ValueOf(&str) //ok fs -> string
	fs.Elem().SetString("jack") //ok
	fmt.Printf("%v\n", str)     // jack
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
