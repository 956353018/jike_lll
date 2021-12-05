package main

import "fmt"

/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
1 <= s.length <= 104
s 仅有英文字母和空格 ' ' 组成
s 中至少存在一个单词
 */
func lengthOfLastWord(s string) int {
	ans := 0
	right := len(s) - 1

	for right >= 0 && s[right] == ' ' {
		right--
	}

	for right >= 0 && s[right] != ' '{
		right--
		ans++
	}

	return ans
}

func main() {
	s := "   fly me   to   the moon  "
	res := lengthOfLastWord(s)
	fmt.Println(res)
}

