package main

func main() {
	stack := NewStack()

	if stack.IsEmpty() {
		stack.Push(1)
		stack.Push(2)

		println(stack.Peek())

		stack.Pop()

		println(stack.Peek())
	}
}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	lastIndex := len(s.data) - 1
	entry := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return entry, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	lastIndex := len(s.data) - 1
	entry := s.data[lastIndex]

	return entry, true
}

func (s *Stack) IsEmpty() bool {
	isEmpty := len(s.data) == 0
	return isEmpty
}
