package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %d \n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{in: make(chan int), done: func() { wg.Done() }}
	go doWorker(id, w)
	return w
}
func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg) // 为每个worker创建一个channel
	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	// wait for all of them
	wg.Wait()
}

func main() {
	channelDemo()
}
