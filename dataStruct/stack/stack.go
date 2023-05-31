package stack

type Stack struct {
	Top     int
	DataLen int
	Data    []interface{}
}

func NewStack(dataLen int) *Stack {
	return &Stack{
		Top:     -1,
		DataLen: dataLen,
		Data:    make([]interface{}, dataLen),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) IsFull() bool {
	return s.Top == s.DataLen-1
}

func (s *Stack) Push(data interface{}) bool {
	if s.IsFull() {
		return false
	}
	s.Top += 1
	s.Data[s.Top] = data
	return true
}

func (s *Stack) Pop() (data interface{}, ok bool) {
	if s.IsEmpty() {
		return 0, false
	}
	data = s.Data[s.Top]
	s.Top -= 1
	return data, true
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.Data[s.Top]
}
