package linked_list

import "fmt"
import "strings"

type SLLNode struct {
	data string
	next *SLLNode
}

// Singly linked list
type SLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

var InvalidIndexError = fmt.Errorf("Invalid Index")

// Insert adds an item to the list at index, or an error if invalid index
func (s *SLinkedList) Insert(data string, index int) error {
	if s.head == nil {
		s.head = &SLLNode{data: data}
		s.tail = s.head
		return nil
	}

	if index == 0 {
		s.head = &SLLNode{data: data, next: s.head}
		return nil
	}

	prev := s.head
	cur := s.head

	for i := 0; i < index; i++ {
		if cur == nil {
			return InvalidIndexError
		}
		prev = cur
		cur = cur.next
	}
	// insert at tail!
	if cur == nil {
		prev.next = &SLLNode{data: data}
		s.tail = prev.next
	} else {
		prev.next = &SLLNode{data: data, next: cur}
	}

	return nil
}

// Get item at index, or returns error
func (s *SLinkedList) Get(index int) (string, error) {
	return "", nil
}

// Delete item at index, or returns error
func (s *SLinkedList) Delete(index int) error {
	return nil
}

func (s *SLinkedList) String() string {
	items := []string{}
	cur := s.head
	for cur != nil {
		items = append(items, cur.data)
		cur = cur.next
	}

	return fmt.Sprintf("[%s]", strings.Join(items, ","))
}
