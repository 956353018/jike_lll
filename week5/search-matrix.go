package main

import "fmt"

var m int
var n int

//1.将二维数组转化为一维数组
func num(i int, j int) int{
	return i * n + j
}

/*
编写一个高效的算法来判断m*n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
 */
func searchMatrix(matrix [][]int, target int) bool {
	m = len(matrix)
	if m == 0 {
		return false
	}
	n = len(matrix[0])
	if n == 0 {
		return false
	}

	left := 0
	right := m * n - 1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid / n][mid % n] == target {  //参考将二维数组转化为一维数组，i * n + j，可知i是商，j是余数
			return true
		} else if matrix[mid / n][mid % n] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1,3,5,7}, {10,11,16,20}, {23,30,34,60}}
	target := 3
	res := searchMatrix(matrix, target)
	fmt.Printf("test search matrix %v", res)
}
