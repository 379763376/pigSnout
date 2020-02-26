package main

import "fmt"

func main() {
	var rmdirs []func()
	mp := make(map[string]string)
	for k, _ := range map[string]string{"aa": "", "bb": "", "cc": ""} {
		//dir := k // NOTE: necessary!
		mp[k] = k // creates parent directories too
		println(k)
		rmdirs = append(rmdirs, func() {
			print(k) //cccccc
			delete(mp, k)
		})
	}
	for _, f := range rmdirs {
		f()
	}

	fmt.Print("\n")
	fmt.Print(mp) //map[aa:aa bb:bb]
}
