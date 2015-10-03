package main

import (
	"fmt"
)

func main() {
	var power int
	power = 9000
	fmt.Printf("It's over %d\n", power)

	var power1 int = 9001
	fmt.Printf("It's over %d\n", power1)

	power2 := 9002
	fmt.Printf("It's over %d\n", power2)

	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}
