package main

import (
	"fmt"
	"math"
)

func main() {
	var value2 int32
	value1 := 64 // value1将会被自动推导为int类型
	// value2 = value1 // 编译错误
	value2 = int32(value1) // 编译通过
	
	fmt.Printf("value1=%v\n", value1)
	fmt.Printf("value2=%v\n", value2)
	
	fmt.Println(5 % 3) // 结果为：2

	m, n := 1, 2
	if m == n {
		fmt.Println("m and n are equal.")
	}

	var i int32
	var j int64
	i, j = 1, 2
	// if i == j { // 编译错误，两个不同类型的整型数不能直接比较
	//	fmt.Println("i and j are equal.")
	// }

	if i == 1 || j == 2 { // 编译通过
		fmt.Println("i and j are equal.")
	}

	var fvalue1 float32
	fvalue1 = 12
	fvalue2 := 12.0 // 如果不加小数点，fvalue2会被推导为整型而不是浮点型
	// fvalue1 = fvalue2 // 编译错误，fvalue2被自动设为float64而非float32
	fvalue1 = float32(fvalue2)
	fmt.Printf("fvalue1=%v\n", fvalue1)
	fmt.Printf("fvalue2=%v\n", fvalue2)

	fmt.Println(IsEqual(float64(fvalue1), fvalue2, 0.00001))
}


// p为用户自定义的比较精度，比如0.00001
func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}
