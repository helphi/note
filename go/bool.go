package main

import "fmt"

func main() {
	var v1 bool
	v1 = true
	v2 := (1 == 2) // v2也会被推导为bool类型
	fmt.Printf("v1=%v,v2=%v\n", v1, v2) 


	var b bool
	// b = 1 // 编译错误
	// b = bool(1) // 编译错误
	fmt.Println("b=", b) 

	b = (1!=0) // 编译正确
	fmt.Println("b=", b) 

}
