package main

import (
	"fmt"
	"math"
)
/*
给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
 */
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	//0/1背包
	//DP函数：所有的可能值落在区间1～n*n，f[i][j]前i个数字凑j，用的最少的数据
	//第i个数，选0个i，f[i][j]=f[i-1][j]
	//第i个数，选1个i，f[i][j]=f[i-1][j-1*i]+1
	//第i个数，选2个i，f[i][j]=f[i-1][j-2*i]+2

	//完全背包
	//f[i][j]-从前i个物品中选出总体积为j，物品总值最大。1）f[i][j]=f[i-1][j]尚未选过第i种物品 2）f[i][j]=f[i][j-v[i]]+w[i] j-v[i] >=0 从第i种物品种选一个

	if n == 1 {
		return 1
	}

	//1.初始化DP
	var f []int
	f = make([]int, n + 1)
	for i := 0; i <= n; i++ {
		f[i] = math.MaxInt32 //初始化为正无穷
	}

	f[0] = 0 //凑0，没有数的平方能凑

	//2.DP决策
	for i := 1; i <= n; i++ {
		for j := 0; j < n && i >= j * j; j++ {  //j的范围1，4，9，j*j<=n
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}

	//返回结果
	return f[n]
}

func main() {
	n := 12
	res := numSquares(n)
	fmt.Println(res)
}
