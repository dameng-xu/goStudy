package heap

func (h *MinHeap) Sort() {
	for i := len(*h) / 2; i >= 0; i-- {
		h.shitDown(i, len(*h)-1)
	}
	for i := 0; i < len(*h); i++ {
		(*h)[0], (*h)[len(*h)-1-i] = (*h)[len(*h)-1-i], (*h)[0]
		h.shitDown(0, len(*h)-2-i)
	}
}
