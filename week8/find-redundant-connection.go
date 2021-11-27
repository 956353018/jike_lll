package main

import "fmt"

/*
树可以看成是一个连通且 无环的无向图。
给定往一棵n 个节点 (节点值1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges，edges[i] = [ai, bi]表示图中在 ai 和 bi 之间存在一条边。
请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组edges中最后出现的边。
 */

/*
[[1,2],[1,3],[2,4],[7,8],[3,4]]
定义一个数组：int[] a,不管它是干啥的
首先，拿到[1,2]这条边，那么代表1能到2，记a[1] = 2
在拿[1,3],那么代表1能到3, 2也能到到达3吧， a[1] = 2(1能到达2) a[1] = 3 (1能到达3)
在拿[2,4],那么代表2能到4, 1,3也能到到达4吧, a[1] =2 a[2]=3 a[2]= 4

 */
var fa1 []int

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	//初始化并查集
	fa1 = make([]int, n + 1)
	for i := 1;i <= n; i++{
		fa1[i] = i
	}

	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		//如果两个元素的根是一个，说明已经成环了，直接返回即可
		if Find(x) == Find(y) {
			return []int{x, y}
		} else{
			UnionSet(x, y)
		}
	}
	return []int{}
}

//并查集函数Find
func Find(x int) int {
	if x == fa1[x] {
		return x
	}
	fa1[x] = Find(fa1[x])
	return fa1[x]
}

func UnionSet(x int, y int) {
	x = Find(x)
	y = Find(y)
	if x != y {
		fa1[x] = y
	}
}

func main() {
	edges := [][]int{{1,2}, {1,3}, {2,3}}
	res := findRedundantConnection(edges)
	fmt.Println(res)
}

