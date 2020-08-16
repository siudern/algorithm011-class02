package Week_08


//冒泡排序
func BubbleSort(arr []int) []int {
	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//插入排序
func insertSort(arr []int) []int{
	for i := 1; i < len(arr); i++ {
		for v := 0; v < i; v++ {
			if arr[v] > arr[i] {
				arr[v], arr[i] = arr[i], arr[v]
			}
		}
	}
	return arr
}

//选择排序
func selectSort(arr []int) {
	var index int
	for i := 0; i < len(arr)-1; i++ {
		index = i
		for j := i + 1; j < len(arr); j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}

}