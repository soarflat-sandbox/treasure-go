package stack

type Stack struct {
	s     []interface{}
	limit int
}

func (s *Stack) Pop() interface{} {
	if len(s.s) == 0 {
		return nil
	}
	ss := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return ss
}

func (s *Stack) Push(i interface{}) {
	s.s = append(s.s, i)
	if len(s.s) > s.limit {
		s.s = s.s[len(s.s)-s.limit:]
	}
}
