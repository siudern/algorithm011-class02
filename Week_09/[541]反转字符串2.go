package Week_09


//给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔2k 个字符的前 k 个字符进行反转。
//
//如果剩余字符少于 k 个，则将剩余字符全部反转。
//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
//示例:
//
//输入: s = "abcdefg", k = 2
//输出: "bacdfeg"
//
//提示：
//
//该字符串只包含小写英文字母。
//给定字符串的长度和 k 在 [1, 10000] 范围内。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-string-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//1 字符串遍历，因为需要交换字符，go中字符串是不可以修改的，需要首先转成字节数组
//2 字符串翻转采用的首尾双指针的方式。首尾交换数据，然后指针同时移动
//3 需要唯一特殊处理的是，按照k 个字符移动，但是可能存在字符长度不够k个长度了， 这个时候需要设置尾指针在字符串尾部

func reverseStr(s string, k int) string {
	str :=[]byte(s)
	var i int
	for i<len(str){
		l:=i
		r:=i+k-1
		if r>=len(str){
			r =len(str)-1
		}
		for l<r&&l<len(str){
			str[l],str[r]=str[r],str[l]
			l++
			r--
		}
		i=i+2*k
	}
	return string(str)
}
