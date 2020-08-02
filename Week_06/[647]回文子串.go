package Week_06

//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。
//
//示例 1:
//
//输入: "abc"
//输出: 3
//解释: 三个回文子串: "a", "b", "c".
//示例 2:
//
//输入: "aaa"
//输出: 6
//说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
//注意:
//
//输入的字符串长度不会超过1000。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindromic-substrings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func countSubstrings(s string) int {
	result:=0
	for i:=0;i<len(s);i++{
		r1:=dp(i,i,s)
		r2:=dp(i,i+1,s)
		result=result+r1+r2
	}
	return result
}

func dp(left int, right int,s string )int{
	result:=0
	for left>=0 && right<len(s){
		if s[left]==s[right]{
			result++
			left--
			right++
		}else{
			return result
		}
	}
	return result
}

