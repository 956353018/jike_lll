package main

import "fmt"

/*
给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
说明：
输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
翻转后单词间应当仅用一个空格分隔。
翻转后的字符串中不应包含额外的空格。
 */
func reverseWords(s string) string {
	//1-1).移除头部和尾部的空格
	i := 0
	j := len(s) - 1
	for i <= j && s[i] == ' ' {
		i++
	}
	for j >= 0 && s[j] == ' ' {
		j--
	}
	//1-2).移除字符中间的多余空格，多余1个
	newS := make([]rune, 0)
	for i <= j {
		if i + 1 < j && s[i + 1] == s[i] && s[i] == ' '{
			i++
			continue
		}
		newS = append(newS, rune(s[i]))
		i++
	}
	fmt.Printf("erase space [%s]\n", newS)

	//2.反转所有字符串
	reverse(newS, 0 , len(newS) - 1)

	//3.在反转每个单词
	reverse_each_word(newS)

	return string(newS)
}

func reverse(s []rune, l int, r int)  {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	fmt.Printf("reverse all string [%s]\n", string(s))
}

func reverse_each_word(s []rune) {
	n := len(s)
	start := 0
	end := 0

	for start < n {
		//查找每个单词的start和end位置
		for end < n && s[end] != ' ' {
			end++
		}
		//反转每个单词
		reverse(s, start, end - 1)
		//更新下一个单词的start和end
		start = end + 1
		end++ //start=end
	}
	fmt.Println(string(s)[start-1:end-1])
}

func main() {
	s := "a good   example"
	res := reverseWords(s)
	fmt.Println(res)
}
