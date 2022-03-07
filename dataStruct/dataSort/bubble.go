package dataSort

// 冒泡排序
// 空间复杂度 1
// 时间复杂度 n^2

func BubbleSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++{
			if s[j] > s[j+1]{
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
