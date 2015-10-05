package main

import "fmt"

func main() {

	value, exists := power("goku")
	if exists == true {
		println(value)
	}

	_, exists1 := power("goku1")
	if exists1 == false {
		println("goku1:9000")
	}

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234

	MyPrintf(v1, v2, v3, v4)

	var arr = []string{"1","2","3","4","5"}[:3]
	MyPrintf(arr[0], arr[2], 9)

	f := func(x, y int) int { return x + y }

	fmt.Println(f(1, 2))

	func1 := func(x, y int) int { return x + y }(3,4)
	fmt.Println(func1)
	fmt.Println(func(x, y int) int { return x + y }(3,4))

	var j int = 5
	a := func()(func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()

	j *= 2
	
	a()

}

func log(message string) {
}

func add(a, b int) int {
	return a+b
}

func power(name string) (int, bool) {
	return 9000, "goku" == name
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
