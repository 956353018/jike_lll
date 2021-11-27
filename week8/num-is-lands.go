package main

import "fmt"

/*
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
 */
var m int
var n int
var fa []int
func numIslands(grid [][]byte) int {
	m = len(grid)
	n = len(grid[0])
	dx := [4]int{-1, 0, 0 ,1}
	dy := [4]int{0, -1, 1, 0}

	//1.并查集
	fa = make([]int, m * n)
	for i := 0; i < m * n; i++ {
		fa[i] = i
	}

	//2.对于每条边，把两个集合合并。
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ { //从[i,j]为顶点，合并邻接点
			if grid[i][j] == '0' {
				continue
			}
			/*
				if i > 0 && grid[i - 1][j] == '1' {
					unionSet(num(i, j), num(i - 1, j))
				}

				if j > 0 && grid[i][j - 1] == '1' {
					unionSet(num(i, j), num(i, j - 1))
				}
			*/
			for k := 0; k < 4; k++ {
				ni := i + dx[k]
				nj := j + dy[k]
				if ni < 0 || nj < 0 || ni >= m || nj >= n {
					continue
				}
				if grid[ni][nj] == '1' { //需要把"临近点"合并
					unionSet(num(i, j), num(ni, nj))
				}
			}
		}
	}

	//3.有几颗树（就是有几颗根）,就是有几个岛屿
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			if find(num(i,j)) == num(i,j) {
				ans++
			}
		}
	}

	return ans
}

func num(i int, j int) int {
	return i * n + j
}

func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func unionSet(x int, y int) {
	x = find(x)
	y = find(y)
	if x != y {
		fa[x] = y
	}
}

func main() {
	grid := [][]byte{{'1','1','1','1','0'}, {'1','1','0','1','0'}, {'1','1','0','0','0'}, {'0','0','0','0','0'}}
	res := numIslands(grid)
	fmt.Println(res)

	grid = [][]byte{{'0'},{'0'}}
	res = numIslands(grid)
	fmt.Println(res)
}



