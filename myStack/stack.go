package myStack

type Stack struct {
	Items []rune
}

func (s *Stack) Push(r rune) {
	s.Items = append(s.Items, r)
}

func (s *Stack) Pop() rune {
	length := len(s.Items)

	if length > 0 {
		lastElement := s.Items[length-1]
		s.Items = s.Items[:length-1]
		return lastElement
	} else {
		return 0
	}
}
