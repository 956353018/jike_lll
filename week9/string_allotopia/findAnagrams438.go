package main

import (
	"fmt"
)

/*【会超时】
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
 */
func findAnagrams(s string, p string) []int {
	//滑动窗口区间[start,end]
	start := 0
	end := 0
	sLen := len(s)
	pLen := len(p)

	ans := make([]int, 0)
	for end < sLen {
		//找到滑动窗口的start和end，判断是否是异位词
		if end - start + 1 == pLen {
			if isAnagram_ok(s[start:end+1], p) {
				ans = append(ans, start)
			}
			start++
		}

		end++
	}
	return ans
}

func isAnagram_ok(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//统计每个字符出现的次数存入hashmap
	map1 := make(map[rune]int)

	for _, ch := range s {
		map1[ch]++
	}

	for _, ch := range t {
		map1[ch]--
		if map1[ch] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)

	s = "abab"
	p = "ab"
	res = findAnagrams(s, p)
	fmt.Println(res)

	s = "abacbabc"
	p = "abc"
	res = findAnagrams(s, p)
	fmt.Println(res)
}


