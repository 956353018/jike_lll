package main

import (
	"fmt"
)

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
bacdfeg
 */
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func reverseString341(s []rune, l int, r int)  {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
func reverseStr(s string, k int) string {
	sAns := []rune(s)
	n := len(s)
	for l := 0; l < n; l = l + 2 * k{
		r := l + k - 1
		reverseString341(sAns, l, min(r, n - 1))
	}
	return string(sAns)
}

func main() {
	s := "abcdefg"
	k := 2
	res := reverseStr(s, k)
	fmt.Println(res)
}
