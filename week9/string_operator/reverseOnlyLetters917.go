package main

import "fmt"

/*
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
 */
func reverseOnlyLetters(s string) string {
	l := 0
	r := len(s) - 1
	newS := []rune(s)
	for l < r {
		if isDigitOrLetter(newS[l]) && isDigitOrLetter(newS[r]) {
			newS[l], newS[r] = newS[r], newS[l]
			l++
			r--
		} else if !isDigitOrLetter(newS[r]){
			r--
		} else if !isDigitOrLetter(newS[l]) {
			l++
		}
	}
	return string(newS)
}

func isDigitOrLetter(ch rune) bool{
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func main() {
	s := "ab-cd"
	res := reverseOnlyLetters(s)
	fmt.Println(res)
}
