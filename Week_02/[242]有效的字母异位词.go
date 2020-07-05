package Week_02


//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//异位词：字母出现的次数都是一样的，但是顺序不一样
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//说明:
//你可以假设字符串只包含小写字母。
//
//进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-anagram


//
//1、暴力解法 利用sort排序之后，对sorted_str进行比较是否相等？时间复杂度O(NlogN）
//2、利用hash的方式，字母作为key，遇到则加1，第二个字符串字符串-1，最后判断为0
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[uint8]int{}
	for i:=0; i<len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, elem := range m {
		if elem != 0 {
			return false
		}
	}
	return true
}

//3、利用数组的方法
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := [26]int{}
	for i:=0; i<len(s); i++ {
		//利用Go语言的字符串存储特性
		//s := 'a'
		//fmt.Printf("%T, %d, %s",s, s, string(s))
		//int32, 97, a
		//相当于: 如果s[i] = 'a' 则 'a'-'a'，即 97-97=0，正好为数组的起始下标，a-z，对应26字母
		a[s[i]-'a']++
	}
	for i:=0; i<len(s); i++ {
		a[t[i]-'a']--
		//如果小于0，则表示在t字符串出现了s字符串中没有的，直接返回false
		if a[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}