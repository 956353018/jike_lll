package main

import (
	"fmt"
)

/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
 */
func canJump(nums []int) bool {
	//DP从j到i（j < i）需要的最小跳跃数,这种方法所依据的核心特性：如果一个位置能够到达，那么这个位置左侧所有位置都能到达
	n := len(nums)

	//DP init
	var f []bool
	f = make([]bool, n)

	//DP
	f[0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j + nums[j] >= i && f[j] == true {
				f[i] = true // 从j跳跃到i
				break
			}
		}
	}

	return f[n - 1]
}

func main() {
	nums := []int{2,3,1,1,4}
	res := canJump(nums)
	fmt.Println(res)

	nums = []int{3,2,1,0,4}
	res = canJump(nums)
	fmt.Println(res)

	nums = []int{0,2,3}
	res = canJump(nums)
	fmt.Println(res)

	nums = []int{0}
	res = canJump(nums)
	fmt.Println(res)
}


