package main

import "math"

func triangle() {
	var a, b int64 = 3, 4
	println(calcTriangle(a, b))
}
func calcTriangle(a, b int64) int64 {
	return int64(math.Sqrt(float64(a*a + b*b)))
}
func main() {
	triangle()
}
