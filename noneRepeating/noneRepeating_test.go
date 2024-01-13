package main

import "testing"

func TestNoneRepeating(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {
		res := lengthOfNoneRepeatingSubStr(tt.s)
		if res != tt.ans {
			t.Errorf("Got %d for input %s; expected %d", res, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8
	res := lengthOfNoneRepeatingSubStr(s)
	for i := 0; i < b.N; i++ {
		if res != ans {
			b.Errorf("Got %d for input %s; expected %d", res, s, ans)
		}
	}
}
