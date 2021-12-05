package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
 */
func longestCommonPrefix(strs []string) string {
	//1.找到长度最短的str
	minLength := len(strs[0])
	minStr := strs[0]
	for _, str := range strs{
		if len(str) < minLength {
			minLength = len(str)
			minStr = str
		}
	}

	//2.遍历所有的strs，与minStr比较，不断缩小最小的字符串minStr
	minStrRune := []rune(minStr)
	for _, str := range strs {
		strRune := []rune(str)
		index := 0
		for index < minLength {
			if minStrRune[index] != strRune[index] {
				minStrRune = minStrRune[:index] //go语言不包括index
				minLength = index
				break
			}
			index++
		}
	}
	if len(minStrRune) ==  0 {
		return ""
	}

	return string(minStrRune)
}

func main() {
	strs := []string{"flower","flow","flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)

	strs1 := []string{"dog","racecar","car"}
	res1 := longestCommonPrefix(strs1)
	fmt.Println(res1)

	strs2 := []string{"flower","fkow"}
	res2 := longestCommonPrefix(strs2)
	fmt.Println(res2)

	strs3 := []string{"abca","aba","aaab"}
	res3 := longestCommonPrefix(strs3)
	fmt.Println(res3)
}

