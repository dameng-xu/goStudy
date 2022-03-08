package dataSort

import (
	"math/rand"
	"time"
)

// 随机快排
// 时间复杂度 nlog(n)
// 空间复杂度 nlog(n)

func Partition(s []int, n int) {
	minIdx := -1
	for i := 0; i < len(s); i++ {
		if s[i] <= n{
			s[minIdx+1], s[i] = s[i], s[minIdx+1]
			minIdx ++
		}
	}
}

func PartitionV2(s []int, n int){
	minIdx := -1
	maxIdx := len(s)

	for i := 0; i < maxIdx; {
		if s[i] == n{
			i ++
		}else if s[i] < n{
			s[minIdx+1], s[i] = s[i], s[minIdx+1]
			minIdx ++
			i ++
		}else {
			s[maxIdx-1], s[i] = s[i], s[maxIdx-1]
			maxIdx --
		}
	}
}

func PartitionV3(s []int, L, R int) []int {
	minIdx := L - 1
	maxIdx := R

	for i := L; i < maxIdx; {
		if s[i] == s[R]{
			i ++
		}else if s[i] < s[R]{
			s[minIdx+1], s[i] = s[i], s[minIdx+1]
			minIdx ++
			i ++
		}else{
			s[maxIdx-1], s[i] = s[i], s[maxIdx-1]
			maxIdx --
		}
	}

	s[maxIdx], s[R] = s[R], s[maxIdx]
	return []int{minIdx+1, maxIdx}
}

func QuickSort(s []int, L, R int){
	if L >= R{
		return
	}

	// 随机
	rand.Seed(time.Now().UnixNano())
	random := L + (R - L) * rand.Intn(99) / 100
	s[random], s[R] = s[R], s[random]

	// 快排
	mid := PartitionV3(s, L, R)
	QuickSort(s, L, mid[0]-1)
	QuickSort(s, mid[1]+1, R)
}
