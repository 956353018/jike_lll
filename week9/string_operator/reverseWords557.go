package main

import "fmt"

/*
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。s
 */
func reverseWords557(s string) string {
	n := len(s)
	start := 0
	end := 0

	var newS []rune
	newS = []rune(s)
	for start < n {
		//查找每个单词的start和end位置
		for end < n && s[end] != ' ' {
			end++
		}
		//反转每个单词
		reverse_each(newS, start, end - 1)

		//更新下一个单词的start和end
		start = end + 1
		end++ //start=end
	}

	return string(newS)
}

func reverse_each(s []rune, l int, r int)  {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	s := "Let's take LeetCode contest"
	res := reverseWords557(s)
	fmt.Println(res)
}
