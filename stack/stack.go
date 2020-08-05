// A Stack of type int.

package stack

import "strings"
import "strconv"

type Stack struct {
	store []int
}

func (s *Stack) Push(val int) {
	s.store = append(s.store, val)
}

func (s *Stack) Pop() (int, bool) {
	count := len(s.store)
	if count == 0 {
		return 0, false
	}
	popped := s.store[count-1]
	s.store = s.store[:count-1]
	return popped, true
}

func (s *Stack) Length() int {
	return len(s.store)
}

func (s *Stack) Peek() (int, bool) {
	count := len(s.store)
	if count == 0 {
		return 0, false
	}
	return s.store[count-1], true
}

func (s *Stack) String() string {
	builder := []string{}
	for _, el := range s.store {
		builder = append(builder, strconv.Itoa(el))
	}

	return strings.Join(builder, "<=")
}
