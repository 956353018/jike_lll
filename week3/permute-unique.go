package main

import (
	"fmt"
	"sort"
)

/*
全排列
# 依次考虑0,1,...,n-1位置放哪个数
# 从还没用过的”数中选一个放在当前位置
*/
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var find func(depth int)
	var path []int

	used := make([]bool, len(nums))

	//1.先排序，这里是去重的关键逻辑，这样才能相邻判断
	sort.Ints(nums)

	find = func(depth int) {
		if len(nums) == depth {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			//2。去重，注意！！先要判断是否是相同元素
			if i > 0 && nums[i - 1] == nums[i] && used[i - 1] == false { //同一层去重效率更高：https://pic.leetcode-cn.com/1631608102-aDQDkD-file_1631608102194
				continue
			}

			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			find(depth + 1)
			path = path[:len(path) - 1]
			used[i] = false
		}
	}
	find(0)
	return res
}


func main() {
	nums := []int{3, 3, 0, 3}
	res := permuteUnique(nums)
	fmt.Printf("test_permute %d\n", res)
}
