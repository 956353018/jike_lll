package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := len(numbers) - 1; i < j && numbers[i] + numbers[j] >= target; j-- {
			if numbers[i] + numbers[j] == target {
				return []int{i, j} //返回的数组下标从0开始
			}
		}
	}
	return []int{}
}
/*
给你一个整数数组 nums 和一个整数k ，请你统计并返回该数组中和为k的连续子数组的个数。
示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2
*/
func subarraySum(nums []int, k int) int {
	//1.前缀和
	n := len(nums)
	s := make([]int, n+1)
	ans := 0

	s[0] = 0
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	//2.
	for i := 1; i <= n; i++ {
		if s[i] == k {
			ans++
		}
		for j := 1; j < i; j++ {
			if s[i] - s[j] == k {
				ans++
			}
		}
	}
	return ans
}

func main() {
	nums := []int{1,1,1}
	k := 2
	//nums := []int{1,2,3}
	//k := 3
	res := subarraySum(nums, k)
	fmt.Println(res)
}
