package heap

import "fmt"

type Heap []int

func (h *Heap) shitUp(pos, ePos int) {
	item := (*h)[pos]
	for pos > ePos {
		pPos := (pos - 1) >> 1
		if pPos >= ePos && (*h)[pPos] < item {
			(*h)[pos] = (*h)[pPos]
			pos = pPos
		} else {
			break
		}
	}
	(*h)[pos] = item
}

func (h *Heap) shitDown(pos, ePos int) {
	item := (*h)[pos]
	for pos < ePos {
		cPos := pos<<1 + 1
		rcPos := cPos + 1
		if rcPos <= ePos && (*h)[cPos] < (*h)[rcPos] {
			cPos = rcPos
		}
		if cPos <= ePos && (*h)[cPos] > item {
			(*h)[pos] = (*h)[cPos]
			pos = cPos
		} else {
			break
		}
	}
	(*h)[pos] = item
}

func (h *Heap) Push(item int) {
	*h = append(*h, item)
	h.shitUp(len(*h)-1, 0)
}

func (h *Heap) Pop() int {
	if len(*h) == 0 {
		return 0
	}
	item := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.shitDown(0, len(*h)-1)
	return item
}

func (h *Heap) List() {
	fmt.Println(*h)
}
