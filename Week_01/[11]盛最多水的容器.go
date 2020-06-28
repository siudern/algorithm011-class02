package Week_01

import "math"

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器，且 n 的值至少为 2。
//
// 
//
//
//
//图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water


//时间复杂度O(n)
//空间复杂度O(1)
// 核心思想，双指针收敛
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	for left < right {
		minHeight := math.Min(float64(height[left]), float64(height[right]))
		area := (right - left) * int(minHeight)
		max = int(math.Max(float64(max), float64(area)))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}