package main

import "fmt"

func lengthOfNoneRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		// rune 转为 字符串
		fmt.Println(i, ch)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		//fmt.Println(start)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	println(lengthOfNoneRepeatingSubStr("abcabcbb"))
	println(lengthOfNoneRepeatingSubStr("bbbbb"))
	println(lengthOfNoneRepeatingSubStr("pwwkew"))

}
