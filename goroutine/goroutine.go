package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("Hello from goroutine %d \n", i)
				// 手动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
