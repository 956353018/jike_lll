package main

import (
	"fmt"
	"math"
)

func min1(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
给你一个非负整数数组nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
 */

func jump(nums []int) int {
	//DP从j到i（j < i）需要的最小跳跃数
	n := len(nums)

	//DP init
	var f []int
	f = make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = math.MaxInt32
	}

	//DP
	f[0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if j + nums[j] >= i{
				f[i] = min1(f[i], f[j] + 1) // 从j跳跃到i的最小值
			}
		}
	}

	return f[n - 1]
}

func main() {
	nums := []int{2,3,1,1,4}
	res := jump(nums)
	fmt.Println(res)
}


