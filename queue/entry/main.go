package main

import (
	"awesomeProject/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)

	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

	//q.Push("abc")
	//fmt.Println(q)
}
