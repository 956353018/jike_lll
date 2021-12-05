package main

import "fmt"

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 */
func firstUniqChar(s string) int {
	sMap := make(map[rune]int)
	sRune := []rune(s)

	for _, v := range sRune {
		sMap[v]++ //字符对应出现的次数
	}

	for idx, v := range sRune {
		if sMap[v] == 1 {
			return idx
		}
	}

	return -1
}

func main() {
	s := "loveleetcode"
	res := firstUniqChar(s)
	fmt.Println(res)

	s = "aaa"
	res = firstUniqChar(s)
	fmt.Println(res)
}


