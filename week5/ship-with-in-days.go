package main

import "fmt"

/*
传送带上的包裹必须在D天内从一个港口运送到另一个港口。
传送带上的第 i个包裹的重量为weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

示例 1：
输入：weights = [1,2,3,4,5,6,7,8,9,10], D = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10
请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。
 */

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

//判定问题
func validate(weights []int, dayTotal int, weightMax int) bool { //参数依次是num:包裹的重量数组，dayTotal:运包裹的总天数 ，weightMax是每天能承担的包裹的总重量
	weightTotal := 0 //每一天的包裹重量计数
	dayCount := 1 //开1天运载

	for _, weight := range weights { //遍历包裹的重量数组
		if weightTotal + weight <= weightMax { //1.一天里放一个weight还未超过weightMax的限制，就把weight放入这一天,这里面是小于等于，只要这一天能容的下就可以，不一定是等号
			weightTotal += weight
		} else {
			dayCount++ //2.放不下，开新的一天运载包裹
			if weight > weightMax { //3.如果weight太大，超过每天能承担的包裹重量，肯定放不下，直接无解
				return false
			}
			weightTotal = weight //2.把新来的包裹重量放进去
		}
	}
	return dayCount <= dayTotal //如果需要的天数不超过给定的天数，就是合法的，返回true
}

func shipWithinDays(weights []int, days int) int {
	//1.二分进行查找，left是包裹的最小值，right所有重量的和
	var left int
	var right int
	for _, weight:= range weights {
		left = max(left, weight)
		right += weight
	}

	//二分模版
	for left < right {
		mid := (left + right) / 2
		if validate(weights, days, mid) { //weights去分段，分成days，每段大小不超过mid
			right = mid //要小的那边
		} else {
			left = mid + 1
		}
	}

	return right
}

func main() {
	weights := []int{3,2,2,4,1,4}
	D := 3
	res := shipWithinDays(weights, D)
	fmt.Println(res)
}

