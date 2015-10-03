package main

func main() {

	value, exists := power("goku")
	if exists == true {
		println(value)
	}

	_, exists1 := power("goku1")
	if exists1 == false {
		println("goku1:9000")
	}
}

func log(message string) {
}

func add(a, b int) int {
	return a+b
}

func power(name string) (int, bool) {
	return 9000, "goku" == name
}
