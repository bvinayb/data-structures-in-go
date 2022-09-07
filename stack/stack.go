package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop will remove the value at the end and return removed value
func (s *Stack) Pop() int {
	lenOfStack := len(s.items) - 1
	toRemove := s.items[lenOfStack]
	s.items = s.items[:lenOfStack]
	return toRemove
}
