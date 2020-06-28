package Week_01

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0，请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum


// 核心思想，排序+双指针的收敛
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) <= 2 {
		return nil
	}
	ans := make([][]int, 0)
	for t:=0; t < len(nums)-2; t++ {
		// 去重判断
		if t > 0 && nums[t] == nums[t-1] {
			continue
		}
		i, j := t + 1, len(nums) -1
		for i < j {
			sum := nums[i] + nums[j] + nums[t]
			// 双指针移动过程中的去重
			if i > t+1 && nums[i] == nums[i-1] {
				i++
				continue
			}
			if j < len(nums) - 1 && nums[j] == nums[j+1] {
				j--
				continue
			}
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[t]})
				i++
				j--
			} else if sum < 0 {
				i++
			} else {
				j--
			}
		}
	}
	return ans
}