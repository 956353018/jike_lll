package main

import (
	"fmt"
	"sort"
)

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
 */
func isAnagram(s string, t string) bool {
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

func isAnagram_2(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

func main() {
	s := "anagram"
	t := "nagaram"
	res := isAnagram(s ,t)
	fmt.Println(res)
}
