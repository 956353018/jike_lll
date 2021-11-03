package main

import "fmt"

/*
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个可能存在 重复元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
 */
func findMin(nums []int) int {
	/* nums[i] <= nums[n - 1],分成两个数组，左数组和右数组
	4 5 6 7 0 1 4
	0 0 0 0 1 1 1

	2 2 2 0 1
	0 0 0 1 1
	 */
	n := len(nums)
	left := 0
	right := n - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] { //mid小，可能有比它还小的，向左找
			right = mid
		} else if nums[mid] > nums[right] { //mid大，向右数组找
			left = mid + 1
		} else { //mid相等情况，也是向左，最终答案是left边界
			right = right - 1
		}
	}

	return nums[left]

}

func main() {

	nums := []int{2,2,2,0,1}
	res := findMin(nums)
	fmt.Println(res)

	nums1 := []int{4,5,6,7,0,1,4}
	res1 := findMin(nums1)
	fmt.Println(res1)

	nums2 := []int{3,1,3}
	res2 := findMin(nums2)
	fmt.Println(res2)

	nums3 := []int{1,3,3}
	res3 := findMin(nums3)
	fmt.Println(res3)

}


