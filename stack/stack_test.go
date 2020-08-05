package stack

import "testing"

func TestPush(t *testing.T) {

	for tc := 0; tc < 100; tc++ {
		s := Stack{}
		for i := 0; i < tc; i++ {
			s.Push(i)
		}

		if s.Length() != tc {
			t.Errorf("Incorrect stack length for %v: Got: %d, Wanted: %d", s, tc, s.Length())
		}
	}
}

func buildStack(insert []int) *Stack {
	s := &Stack{}
	for _, el := range insert {
		s.Push(el)
	}
	return s
}

func TestPop(t *testing.T) {
	cases := []struct {
		insert    []int
		valid     bool
		expected  int
		remaining []int
	}{
		{
			insert:    []int{1, 2, 3},
			valid:     true,
			expected:  3,
			remaining: []int{1, 2},
		},
		{
			insert:    []int{},
			valid:     false,
			expected:  0,
			remaining: []int{},
		},
	}

	for _, tc := range cases {
		s := buildStack(tc.insert)

		first, ok := s.Pop()
		if ok != tc.valid {
			t.Errorf("Expected: %v, Got: %v", tc.valid, ok)
		}

		if first != tc.expected {
			t.Errorf("Expected: %v, Got: %v", tc.expected, first)
		}

		if len(tc.remaining) != s.Length() {
			for _, exp := range tc.remaining {
				pop, ok := s.Pop()
				if !ok || pop != exp {
					t.Errorf("Expected remaining: %v, Got: %v", tc.remaining, s)
					break
				}
			}
		}
	}
}

func TestPeek(t *testing.T) {
	cases := []struct {
		insert   []int
		valid    bool
		expected int
	}{
		{insert: []int{1, 2, 3}, valid: true, expected: 3},
		{insert: []int{}, valid: false, expected: 0},
	}

	for _, tc := range cases {
		s := buildStack(tc.insert)

		first, ok := s.Peek()
		if ok != tc.valid {
			t.Errorf("Expected: %v, Got: %v", tc.valid, ok)
		}

		if first != tc.expected {
			t.Errorf("Expected: %v, Got: %v", tc.expected, first)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		insert   []int
		expected string
	}{
		{
			insert:   []int{4, 8, 15, 16, 23, 42},
			expected: "4<=8<=15<=16<=23<=42",
		},
		{
			insert:   []int{},
			expected: "",
		},
		{
			insert:   []int{23},
			expected: "23",
		},
	}

	for _, tc := range cases {
		s := Stack{}
		for _, el := range tc.insert {
			s.Push(el)
		}

		if s.String() != tc.expected {
			t.Errorf("Got: %s, Expected: %s", s.String(), tc.expected)
		}
	}

}
