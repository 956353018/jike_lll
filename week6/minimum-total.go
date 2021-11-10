package main

import (
	"fmt"
	"math"
)

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标
相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么
下一步可以移动到下一行的下标 i 或 i + 1 。

 */
/*
思路：f[i][j]=min(f[i−1][j−1],f[i−1][j])+c[i][j]
当j=0，f[i][0]=f[i−1][0]+c[i][0] //对应j=0，只能从i-1行向下移动得到
当j=i,f[i][i]=f[i−1][i-1]+c[i][i] //对应第i行，只能从i-1的对角线移动得到
当i=0，f[0][0]=c[0][0]
 */
func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	//1.边界
	m := len(triangle)
	if m == 0 {
		return 0
	}

	//2.初始化同triangle大小的数组
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, len(triangle[i]))
	}

	f[0][0] = triangle[0][0]
	var n int
	for i := 1; i < m; i++ {
		n = len(triangle[i])
		for j := 0; j < n; j++ {
			if j == 0 {
				f[i][j] = f[i - 1][j] + triangle[i][j]
			} else if i == j {
				f[i][j] = f[i - 1][i - 1] + triangle[i][j]
			} else {
				f[i][j] = min(f[i - 1][j - 1], f[i - 1][j]) + triangle[i][j]
			}
		}
	}
	ans := math.MaxInt32
	for k := 0; k < len(triangle[m - 1]); k ++ {
		ans = min(ans, f[m - 1][k])
	}

	return ans
}

func main() {
	triangle := [][]int{{2}, {3,4}, {6,5,7}, {4,1,8,3}}
	res := minimumTotal(triangle)
	fmt.Println(res)

	triangle1 := [][]int{{-10}}
	res1 := minimumTotal(triangle1)
	fmt.Println(res1)

}