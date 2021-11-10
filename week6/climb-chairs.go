package main

import (
	"fmt"
)

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
 */

func climbStairs(n int) int {
	//当前级i都可以是先到i-1或者i-2两种走法到i级
	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	n := 2
	res := climbStairs(n)
	fmt.Println(res)
}


