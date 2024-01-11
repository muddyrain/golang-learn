package main

import (
	"awesomeProject/functional/fib"
	"bufio"
	"errors"
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
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error")
	if err != nil {
		pathError, ok := err.(*os.PathError)
		if !ok {
			panic(err)
		} else {
			fmt.Println("pathError", pathError.Path, pathError.Op, pathError.Err)
		}
		return
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
	writerFile("fib.txt")
}
