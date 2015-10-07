package main

import "fmt"
import "sync"
import "runtime"

var counter int = 0
var once sync.Once

func Count(lock *sync.Mutex) {
	once.Do(func () { fmt.Println("IN Count") })
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched() //出让时间片
		if c >= 10 {
			break
		}
	}
}
