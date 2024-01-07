package main

import (
	"strconv"
)

func main() {

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

func updateSlices(s []int) {
	s[0] = 100
}
