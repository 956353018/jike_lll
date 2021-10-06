package main

import (
	"fmt"
)

type rect struct {
	width int
	height int
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	//3.最后的答案
	ans := 0
	//1.定义1个stack
	stack := make([]rect, 0)
	//4.保证栈最后被弹空，需要将最后1个元素设置为0
	heights = append(heights, 0)

	//1.遍历每根柱子的高度
	for _, height := range heights {
		accumulateWidth := 0
		//2.条件：栈顶高度大于当前高度，单调性破坏，1）需要更新宽度；2)计算面积，更新最大值；3）删除栈顶
		for len(stack) > 0 && stack[len(stack) - 1].height >= height {
			accumulateWidth += stack[len(stack) - 1].width //1）需要更新宽度
			ans = max(ans, stack[len(stack) - 1].height * accumulateWidth) //2)计算面积
			stack = stack[:len(stack) - 1] //3）删除栈顶
		}
		stack = append(stack, rect{accumulateWidth + 1, height})
	}

	return ans
}

func maximalRectangle(matrix [][]byte) int {
	ans := 0

	row := len(matrix)//行数,空行判断
	if row == 0 {
		return 0
	}

	col := len(matrix[0])//列数，只有一行
	if col == 0 {
		return 0
	}

	//遍历列，计算列出现1的个数，如果列出现0，整个heiht为0
	heights := make([]int, col) //定义一个数组存储每次计算列的连续长度
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, largestRectangleArea(heights))
	}
	return ans
}

func main() {
	//matrix := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	//matrix := [][]byte{}
	matrix := [][]byte{{'1','0'}}
	res := maximalRectangle(matrix)
	fmt.Printf("test maximalRectangle[%d]", res)
}
