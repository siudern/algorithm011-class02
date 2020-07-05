package Week_02

// 看了hash table之后，为了加强对算法的理解，对两数之和题目再利用hash table做一遍
// hash table  本质是利用 key value的形式去实现，在GO语言中是map

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

//暴力法是循环两边，时间复杂度时候O(n^2)
//通过hash的方式，时间复杂度降为O(n)，n取决于数组的长度，空间复杂度也是O(n)
//实现思路，将a + b = target 转换位target - a = b的形式
//判断条件则是：target - a的值，是否能在map中的key找到，如果没有找到，则把a放入到map，当下一个来查询的时候，则直接找到这个key是否在map中
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, value := range(nums) {
		// existIndex 代表如果能找到，则可以返回找到数的索引
		//hasKey 代表能否找到m的key，如果找不到会返回false（成功状态）
		existIndex, hasKey := m[target-value]
		//如果成功了，说明m中（target-value）形成的key是存在的，满足了a = target - b的条件
		if hasKey {
			return []int{existIndex, index}
		}
		//如果失败，则在这个m中加入相应的key，value，数组的值为key，数组的index为value
		m[value] = index
	}
	return nil
}