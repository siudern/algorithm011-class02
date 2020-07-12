package Week_03
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 
//
//示例：
//
//输入：n = 3
//输出：[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/generate-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


//利用递归的思想
//两个合法条件判断：
//必须要现有左括号才有右括号
//左括号的最大数量小于MAX括号数
//时间复杂度O(n)
//空间复杂度O(n)，取决于栈递归的深度
func generateParenthesis(n int) []string {
	output := make([]string, 0)
	_generate(0, 0, n, "", &output)
	return output
}

func _generate(left, right, max int, s string, output *[]string) {
	if left == max && right == max {
		*output = append(*output, s)
		return
	}

	if left < max {
		_generate(left+1, right, max, s+"(", output)
	}
	if right < left {
		_generate(left, right+1, max, s+")", output)
	}
}
