package Week_09

//给定一个非空字符串s，最多删除一个字符。判断是否能成为回文字符串。
//
//示例 1:
//
//输入: "aba"
//输出: True
//示例 2:
//
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
//注意:
//
//字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-palindrome-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



//时间复杂度：O(n)，其中 n 是字符串的长度。判断整个字符串是否是回文字符串的时间复杂度是O(n)，遇到不同字符时，判断两个子串是否是回文字符串的时间复杂度也都是O(n)。
//空间复杂度：O(1)。只需要维护有限的常量空间

func validPalindrome(s string) bool {
	low, high := 0, len(s) - 1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high - 1; i < j; i, j = i + 1, j - 1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low + 1, high; i < j; i, j = i + 1, j - 1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
