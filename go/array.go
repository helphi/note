package main

import "fmt"

// 在Go语言中数组是一个值类型（value type）。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。如果将数组作为函数的参数类型，则在函数调用时该参数将发生数据复制。因此，在函数体中无法修改传入的数组的内容，因为函数内操作的只是所传入数组的一个副本。
func modify(array [5]int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}

func main() {
	array := [5]int{1,2,3,4,5} // 定义并初始化一个数组
	modify(array) // 传递给一个函数，并试图在函数体内修改这个数组内容
	fmt.Println("In main(), array values:", array)
}
