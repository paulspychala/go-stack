// Copyright 2014 Pawel Spychala

package stack

type Element struct {
	Value interface{}
	prev  *Element
}

type Stack struct {
	top *Element
	Len int
}

func New() *Stack {
	return new(Stack)
}

func (s *Stack) Push(val interface{}) {
	n := &Element{val, s.top}
	s.top = n
	s.Len++
}

func (s *Stack) Pop() interface{} {
	if s.Len > 0 {
		ret := s.top
		s.top = s.top.prev
		s.Len--
		return ret.Value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	return s.top.Value
}

func (s *Stack) Clear() {
	s.Len = 0
	s.top = nil
}
