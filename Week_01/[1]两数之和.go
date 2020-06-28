package Week_01


//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
// 
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum

//暴力解析法
//时间复杂度：O(n^2)
//空间复杂度O(1)
func twoSum(nums []int, target int) []int {
	for i :=0; i < len(nums); i++ {
		for j := len(nums) -1; j > i; j-- {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//map法
//时间复杂度O(n)
//空间复杂度O(n)
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	//遍历数组的下标和值，把值作为map的key存储，把下标作为map的value存储
	for i, v := range nums {
		k, ok := m[target-v]  //当目标减去数组的值，这个值又是map的key，即在map中找得到时，则返回map的值（数组下标）
		if ok {
			return []int{i, k}
		} else {
			m[v] = i
		}
	}
	return nil
}

