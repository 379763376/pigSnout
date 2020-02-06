package main

import "fmt"

//5cock 3hen 1/3chicken
func main() {
	//方法一：
	//cock 0~20
	//hen 0~33
	count1 := 0 //累加次数
	for cock:=0;cock<=20 ;cock++  {
		for hen:=0;hen<33 ;hen++  {
			count1 += 1
			chicken := 100-cock-hen
			if 5*cock+3*hen+chicken/3 == 100 &&
				chicken % 3 == 0{
				fmt.Println(cock,hen,chicken)
			}
		}
	}

	//方法二：
	count2 := 0
	for cock:=0;cock<=20;cock++{
		for hen:=0;hen<=33 ;hen++  {
			for chicken:=0;chicken<=100 ;chicken+=3  {
				count2+=1
				if 5*cock+3*hen+chicken/3 == 100 &&cock+hen+chicken==100 {
					fmt.Println(cock,hen,chicken)
				}
			}
		}
	}

}