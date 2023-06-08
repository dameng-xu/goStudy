package devide

func HalfDevide(data []int, target int) int {
	if len(data) == 0 {
		return -1
	}

	left := 0
	right := len(data) - 1

	if data[left] == target {
		return 0
	}
	if data[right] == target {
		return right
	}

	for left <= right {
		mid := left + (right-left)>>1
		if data[mid] == target {
			return mid
		} else if target < data[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
