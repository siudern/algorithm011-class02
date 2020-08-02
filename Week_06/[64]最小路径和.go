package Week_06

//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
//  [1,3,1],
//[1,5,1],
//[4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 { //常见的 LeetCode 阴人输入，也是必要的防御式编程
		return 0
	}
	for index, row := range grid { //遍历一下网格
		for i, left, top := 0, 0, 0; i < len(row); i, left, top = i+1, 0, 0 { //因为 Walk 只能来自左、上，所以记录来自左、上的步数
			if i > 0 { //如果不是行的首列
				left = row[i] + row[i-1] //那么可能来自左侧
				if index == 0 {          //如果该行是首行
					grid[index][i] = left //那么只能来自左侧
					continue
				}
			}
			if index > 0 { //如果不是行的首行
				top = row[i] + grid[index-1][i] //那么可能来自顶部
				if i == 0 {                     //如果该行是首列
					grid[index][i] = top //那么只能来自顶部
					continue
				}
			}
			if i != 0 && index != 0 { //非首行首列的单元格可能来自左、上
				if grid[index][i] = left; left > top { //语法比较奇特，实际上是 grid[index][i] = min(left, top)
					grid[index][i] = top
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1] //漫步结束的最后一格即最优解
}
