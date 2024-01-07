package main

import (
	"strconv"
)

func main() {
	var a int = 2
	var b int = 3
	a, b = swap(a, b)
	println(a, b)
}

func convertToBin(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func swap(a, b int) (int, int) {
	return b, a
}
