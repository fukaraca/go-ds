package stacks

import (
	"sync"
)

type Stack struct {
	stack []interface{}
	lock  *sync.Mutex
	len   int
}

//creates a new stack structure
func New() *Stack {
	nev := Stack{}
	nev.lock = &sync.Mutex{}
	return &nev
}

//pushes element to top
func (s *Stack) Push(elem interface{}) {
	s.lock.Lock()
	s.stack = append(s.stack, elem)
	s.len++
	s.lock.Unlock()
}

//pops element at the top
func (s *Stack) Pop() interface{} {
	if s.len > 0 {
		s.lock.Lock()
		popped := s.stack[s.len-1]
		s.stack = s.stack[:s.len-1]
		s.len--
		s.lock.Unlock()
		return popped
	}
	return nil
}

//returns element at the top
func (s *Stack) Top() interface{} {
	if s.len > 0 {
		s.lock.Lock()
		top := s.stack[s.len-1]
		s.lock.Unlock()
		return top
	}
	return nil
}

//returns true if length of the stack is zero
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

//returns length of the stack
func (s *Stack) Len() int {
	return s.len
}
