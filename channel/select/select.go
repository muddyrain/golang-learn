package main

import (
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	var c1, c2 = generator(), generator()

	select {
	case n1 := <-c1:
		println("received", n1, "from c1")
	case n2 := <-c2:
		println("received", n2, "from c2")
		//default:
		//	println("no communication")
	}

}
