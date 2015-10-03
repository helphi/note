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

	// power2 := 9002
	// COMPILER ERROR:
	// no new variables on left side of :=

	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)

	fmt.Printf("power1:%d,power2:%d\n", power1, power2)
	power1, power2 = power2, power1
	fmt.Printf("power1:%d,power2:%d\n", power1, power2)
}
