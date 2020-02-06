<<<<<<< HEAD
package main
import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main() {

	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
	
=======
package main
import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main() {

	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
	
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
}