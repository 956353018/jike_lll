package main

import "fmt"

/*
给定两个字符串s和t，判断它们是否是同构的。

如果s中的字符可以按某种映射关系替换得到t，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
 */
func isIsomorphic(s string, t string) bool {
	map1 := make(map[byte]byte)
	map2 := make(map[byte]byte)

	//建立字符之间的映射关心
	for i := 0; i < len(s); i++ {
		if _, ok := map1[s[i]]; !ok { //s[i]->t[i]映射
			map1[s[i]] = t[i]
		}
		if _, ok := map2[t[i]]; !ok { //t[i]->s[i]映射
			map2[t[i]] = s[i]
		}

		if map1[s[i]] != t[i] || map2[t[i]] != s[i] {//字符互相不能映射，返回false
			return false
		}
	}

	return true
}

func main() {
	s := "egg"
	t := "add"
	res := isIsomorphic(s, t)
	fmt.Println(res)
}


