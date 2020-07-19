package Week_04

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
// 
//
//示例 1:
//
//输入:
//[
//['1','1','1','1','0'],
//['1','1','0','1','0'],
//['1','1','0','0','0'],
//['0','0','0','0','0']
//]
//输出: 1
//示例 2:
//
//输入:
//[
//['1','1','0','0','0'],
//['1','1','0','0','0'],
//['0','0','1','0','0'],
//['0','0','0','1','1']
//]
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-islands
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



//深度优先搜索

func numIslands(grid [][]byte) int {
	count := 0
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] != '0' {
				dfs(grid, line, column)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, line int, column int) {
	if line < 0 || line >= len(grid) || column < 0 || column >= len(grid[line]) {
		return
	}
	if grid[line][column] == '0' {
		return
	}
	grid[line][column] = '0'
	dfs(grid, line+1, column)
	dfs(grid, line-1, column)
	dfs(grid, line, column+1)
	dfs(grid, line, column-1)
}

