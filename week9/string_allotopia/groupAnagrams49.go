package main

import (
	"fmt"
	"sort"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
 */
func groupAnagrams(strs []string) [][]string {
	hashMapArray := make(map[string][]string)
	for _, str := range strs {
		//1.字符排序
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

		//2.相同key的字符串放入相同数组
		key := string(s)
		hashMapArray[key] = append(hashMapArray[key], str)
	}

	//3.输出到答案
	ans := [][]string{}
	for _, strArray := range hashMapArray {
		ans = append(ans, strArray)
	}

	return  ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
