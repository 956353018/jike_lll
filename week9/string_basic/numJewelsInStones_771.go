package main

import "fmt"

/*
给你一个字符串 jewels代表石头中宝石的类型，另有一个字符串 stones 代表你拥有的石头。stones中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
字母区分大小写，因此 "a" 和 "A" 是不同类型的石头。
 */
func numJewelsInStones(jewels string, stones string) int {
	//1.hashmap
	jewelsMap := make(map[rune]int)
	jewelsRune := []rune(jewels) //打散string到char
	stonesRune := []rune(stones)

	//2.计算每个宝石出现的次数
	for _, v := range jewelsRune{
		jewelsMap[v]++
	}

	//3.遍历石头
	ans := 0
	for _, v := range stonesRune {
		if _, ok := jewelsMap[v]; ok {
			ans++
		}
	}

	return ans
}

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	res := numJewelsInStones(jewels, stones)
	fmt.Println(res)
}
