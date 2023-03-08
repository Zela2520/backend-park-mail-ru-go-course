package src

type Stack struct {
	value []interface{}
}

func (s *Stack) Size() int {
	return len(s.value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.value) == 0
}

func (s *Stack) Top() interface{} {
	return s.value[len(s.value)-1]
}

func (s *Stack) Pop() interface{} {
	deletedValue := s.value[len(s.value)-1]

	s.value = s.value[:(len(s.value) - 1)]

	return deletedValue
}

func (s *Stack) Push(value interface{}) {
	s.value = append(s.value, value)
}
