package main

import "fmt"

func main() {
	var str string // 声明一个字符串变量
	str = "Hello world" // 字符串赋值
	ch := str[0] // 取字符串的第一个字符
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	str = "Hello world" // 字符串也支持声明时进行初始化的做法
	// str[0] = 'X' // 编译错误

	str = "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ { // 以字节数组的方式遍历
	ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}

	for i, ch := range str { // 以Unicode字符遍历
		fmt.Println(i, ch)//ch的类型为rune
	}

}
