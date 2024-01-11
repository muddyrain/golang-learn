package main

import (
	"awesomeProject/functional/fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writerFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//writerFile("fib.txt")
	tryDefer()
}
