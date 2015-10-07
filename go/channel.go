package main

import "fmt"
import "time"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func read(ch <-chan int) {
	i := 0
	for value := range ch {
		if 5==i {
			break
		}
		fmt.Println("read value ", value)
		i++
	}
}

func write(ch chan<- int) {
	fmt.Println(ch)
	for i:=0; i<10; i++ {
		ch <- i
		fmt.Println("write value ", i)
	}
}

func main() {
	ch := make(chan int, 1024)

	// 首先，我们实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待1秒钟
		timeout <- true
	}()

	// 然后我们把timeout这个channel利用起来
	select {
	case <-ch:
		// 从ch中读取到数据
		for i := range ch {
			fmt.Println(i)
		}
	case <-timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
		fmt.Println("timeout")
	}

	ch = make(chan int, 10)
	chin := chan<- int(ch)
	fmt.Println("cap=", cap(ch))
	fmt.Println(ch)
	fmt.Println(chin)
	write(chin)
	read(ch)

	ch = make(chan int, 1)
	for i:=0; i<10; i++ {
		select {
		case ch <- 0:
			fmt.Println("0 write")
		case ch <- 1:
			fmt.Println("1 write")
	}
	i := <-ch
	fmt.Println("Value received:", i)
	}

	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int, 1)
		go Count(chs[i])
	}
	for _, ch := range(chs) {
		v, ok := <-ch
		if ok {
			close(ch)
			fmt.Println("chan closed")
		}
		fmt.Println(v)
	}

	ch9 := make(chan int, 10)
	ch9 <- 9
	ch10 := chan<- int(ch9)
	ch11 := <-chan int(ch9)
	ch10 <- 10
	fmt.Println(ch9)
	fmt.Println(ch10)
	fmt.Println(ch11)
	fmt.Println(<-ch9)
	//fmt.Println(<-ch10) //error
	//fmt.Println(<-ch11) //error
}
