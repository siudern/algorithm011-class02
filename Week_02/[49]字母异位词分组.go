package Week_02

import (
	"sort"
)

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/group-anagrams


//思路1
//当排序后的字符串相等时，两个字符串是字母异位词，如： tab与bat，tab->abt == bat->abt ，相等
//首先对切片中的每个字符串元素进行排序，Go语言中对自定义数据的排序（sort.Sort）方法需要自己实现
//利用map，排序后的字符串作为key，如果匹配则把原字符串加入到value中，如果不匹配则把该排序字符串新增为key
//时间复杂度为O(NKlogK)，N为字符串切片的长度，K为字符串的最大长度，看源码sort.Sort的实现是快排，快排的时间复杂度为O(nlogn)
//得出该思路的时间复杂度为O(NKlogK)
//空间复杂度为O(NK)，即整个切片中全部的字符串信息

type byteStringSlice []byte

func (b byteStringSlice) Len() int {
	return len(b)
}

func (b byteStringSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b byteStringSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}


func groupAnagrams(strsSlice []string) [][]string {
	if len(strsSlice) == 0 {
		return nil
	}
	ret := [][]string{}
	m := map[string]int{}
	for _, s := range strsSlice {
		byteS := byteStringSlice(s) //要将字符串转换成byte类型的切片才能排序
		sort.Sort(byteS) // 这是个原地操作，因为切片是引用类型
		sortByteS := string(byteS) //再转换
		//Go的写法，可以通过该方法直接判断m[k]是否存在，!ok即无
		if retIndex, ok := m[sortByteS]; !ok {
			//排序后的字符串作为map的key，因为是没有出现过的新字符串，所以长度正好为ret的index
			//后续添加已存在的字符串，直接从该处取ret的index即可
			m[sortByteS] = len(ret)  //len(ret)为对应ret的一级下标，这里绝对不可能重复，因为只要添加未出现字符串之后，len(ret)则会加1
			ret = append(ret, []string{s}) // 当ret为空的时候，即append(ret[0], []string{s}
		} else {
			ret[retIndex] = append(ret[retIndex], s) //当ret不为空的时候，则为append(ret[m[sortByteS]], s)
		}
	}
	return ret
}

//思路2
//当一个字符串中每个字符的计数相等时，两个字符串是字母异位词，如area与raea，a二次，r一次，e一次，相等
//

func groupAnagrams2(strsSlice []string) [][]string {
	if len(strsSlice) == 0 {
		return nil
	}
	ret := [][]string{}
	m := map[string]int{}
	for i := range strsSlice {
		b := []byte{}
		count := make([]int, 26)
		for _, s := range strsSlice[i] {
			// 遇到哪个字母，对应的元素值就 +1
			count[s - 'a']++
		}
		for singleStr, strCount := range count {
			//把所有遇到的字符统计起来，并加上次数比如area  最后得出  a2e1r1
			if strCount != 0 {
				b = append(b, byte(singleStr)+'a', byte(strCount))
			}
		}
		s := string(b)
		if retIndex, ok := m[s]; !ok {
			m[s] = len(ret)
			ret = append(ret, []string{strsSlice[i]})
		} else {
			ret[retIndex] = append(ret[retIndex], strsSlice[i])
		}
	}
	return ret
}