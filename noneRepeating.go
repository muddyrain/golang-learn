package main

func lengthOfNoneRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
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
	println(lengthOfNoneRepeatingSubStr(""))
	println(lengthOfNoneRepeatingSubStr("b"))
	println(lengthOfNoneRepeatingSubStr("abcdef"))
	println(lengthOfNoneRepeatingSubStr("这里是慕课网"))
	println(lengthOfNoneRepeatingSubStr("一二三二一"))
}
