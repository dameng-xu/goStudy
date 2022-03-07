package dataSort

func MergeSort(s []int, lft, rht int){
	if lft == rht{
		return
	}

	mid := lft + (rht - lft) >> 1
	MergeSort(s, lft, mid)
	MergeSort(s, mid+1, rht)

	merge(s, lft, mid, rht)
}

func merge(s []int, lft, mid, rht int) []int{
	help := make([]int, rht-lft+1)
	i := 0
	P1 := lft
	P2 := mid + 1

	for P1 <= mid && P2 <= rht{
		if s[P1] <= s[P2]{
			help[i] = s[P1]
			P1 ++
			i ++
		}else {
			help[i] = s[P2]
			P2 ++
			i ++
		}
	}
	for P1 <= mid{
		help[i] = s[P1]
		P1 ++
		i ++
	}
	for P2 <= rht{
		help[i] = s[P2]
		P2 ++
		i ++
	}

	for j := 0; j < len(help); j++ {
		s[lft+j] = help[j]
	}
	return s
}