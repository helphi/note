package main

import "fmt"

func main() {
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(myArray), cap(myArray))

	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))

	mySlice = myArray[5:]
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))

	mySlice = myArray[3:8]
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))

	// 创建一个初始元素个数为5的数组切片，元素初始值为0：
	mySlice = make([]int, 5)
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))

	// 创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间：
	mySlice = make([]int, 5, 10)
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))

	// 直接创建并初始化包含5个元素的数组切片：
	mySlice = []int{1, 2, 3, 4, 5}
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 1, 2, 3)
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))



	mySlice2 := []int{8, 9, 10}
	// 给mySlice后面添加另一个数组切片
	mySlice = append(mySlice, mySlice2...)
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(mySlice), cap(mySlice))

	newSlice := mySlice[:15] // 基于oldSlice的前15个元素构建新数组切片
	for _, v := range newSlice {
		fmt.Print(v, " ")
	}
	fmt.Printf("len=%v,cap=%v\n", len(newSlice), cap(newSlice))

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	for _, v := range slice1 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	for _, v := range slice2 {
		fmt.Print(v, " ")
	}
	
	fmt.Print("\nmyArray:")
	fmt.Println(myArray)
	mySlice = myArray[:5]
	fmt.Print("mySlice:")
	fmt.Println(mySlice)
	mySlice[4] = 100 // 改变数据切片的值也会引向数组
	fmt.Print("myArray:")
	fmt.Println(myArray)
}


