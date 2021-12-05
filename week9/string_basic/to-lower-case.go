package main
/*
给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。
1。大写字母变小写字母时加上32
 */
func toLowerCase(s string) string {
	index := 0
	ans := []rune(s)
	diff := 'a' - 'A'

	for index < len(s) {
		if s[index] >= 'A' && s[index] <= 'Z' {
			ans[index] = ans[index] + diff
		}
		index++
	}
	return string(ans)
}




