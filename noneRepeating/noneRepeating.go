package main

var lastOccurred = make([]int, 0xffff)

func lengthOfNoneRepeatingSubStr(s string) int {
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		// rune 转为 字符串
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
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
