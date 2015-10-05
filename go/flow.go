package main

import "fmt"

func main() {
	fmt.Println(compare(3,4))

	var n int = 99
	switch n {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough // 继续执行下一个case
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}

	fmt.Println()

	Num := 5
	switch {
	case 0 <= Num && Num <= 3:
		fmt.Printf("0-3")
	case 4 <= Num && Num <= 6:
		fmt.Printf("4-6")
	case 7 <= Num && Num <= 9:
		fmt.Printf("7-9")
	}

	fmt.Println()

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum=%d\n", sum)

	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Printf("sum=%d\n", sum)

	a := []int{1, 2, 3, 4, 5, 6}
	for j, i := len(a) - 1, 0; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)

	for j := 0; j < 5; j++ {
		L1:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break L1
			}
			fmt.Printf("i=%d\n", i)
		}
	}


	j := 0
FLAG:

	j++
	if j < 10 {
		fmt.Printf("j=%d\n", j)
		goto FLAG
	}
}

func compare(m, n int) bool {
	if m < n {
		return false
	} else {
		return true
	}
}
