package main

import "fmt"

//定义方向数组
var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}

//全局变量
var m int
var n int

//定义dfs
func dfs(board [][]byte, x, y int) {
	//边界问题
	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}
	//递归更新出现'O'的坐标
	if board[x][y] == 'O' {
		board[x][y] = '*'
		for k := 0; k < 4; k++ {
			nx := x + dx[k]
			ny := y + dy[k]
			dfs(board, nx, ny)
		}
	}
}

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
*/
func solve(board [][]byte) {
	//1.初始化,单行单列不存在围绕区域
	m = len(board)
	if m == 0 {
		return
	}

	n = len(board[0])
	if n == 0 {
		return
	}

	//1.遍历每个坐标点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//如果矩阵中元素等于'O'且是边界，就标记该坐标为'*'
			if i == 0 || i == m - 1 || j == 0 || j == n - 1 {
				if board[i][j] == 'O' {
					dfs(board, i, j)
				}
			}
		}
	}

	//2.二次扫描，把标记为'*'还原为'O'，未标记为'*'的'O'设置为'X'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X','O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Printf("test slove %c", board)
}