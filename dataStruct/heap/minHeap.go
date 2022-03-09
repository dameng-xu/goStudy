package heap

type Heap []int


func _siftDown(h Heap, startPos, pos int){
	newItem := h[pos]
	for pos > startPos {
		parentPos := (pos - 1) >> 1
		parent := h[parentPos]
		if newItem < parent{
			h[pos] = parent
			pos = parentPos
			continue
		}else {
			break
		}
	}
	h[pos] = newItem
}


func _siftUp(h Heap, pos int){
	endPos := len(h)
	startPos := pos
	newItem := h[pos]

	childPos := 2 * pos + 1

	for childPos < endPos{
		rightPos := childPos + 1
		if rightPos < endPos && h[rightPos] > h[childPos]{
			childPos = rightPos
		}
		h[pos] = h[childPos]
		pos = childPos
		childPos = 2 * pos + 1
	}

	h[pos] = newItem
	_siftDown(h, startPos, pos)
}


func MyHeapPush(h Heap, n int) Heap {
	h = append(h, n)
	_siftDown(h, 0, len(h)-1)
	return h
}

func MyHeapPop(h Heap) int {
	var returnItem int

	lastItem := h[len(h)-1]
	h = h[:len(h)-1]

	if len(h) > 0{
		returnItem = h[0]
		h[0] = lastItem
		_siftUp(h, 0)
	}else {
		returnItem = lastItem
	}
	return returnItem
}
