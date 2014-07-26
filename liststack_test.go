// Copyright 2014 Pawel Spychala

package stack

import "testing"

func TestStack(t *testing.T) {
	s1 := New()

	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	checkStack(t, s1, []interface{}{3, 2, 1})

	s1.Push(4)
	s1.Push(5)

	checkStack(t, s1, []interface{}{5, 4, 3, 2, 1})

	s1.Pop()
	s1.Pop()
	s1.Pop()

	checkStack(t, s1, []interface{}{2, 1})

	s1.Pop()
	s1.Pop()

	checkStack(t, s1, []interface{}{})

	s1.Pop()

	checkStack(t, s1, []interface{}{})

	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	checkStack(t, s1, []interface{}{3, 2, 1})

	s1.Clear()

	checkStack(t, s1, []interface{}{})
}

func checkStack(t *testing.T, s *Stack, es []interface{}) {

	i := 0
	for e := s.top; e != nil; e = e.prev {
		se := e.Value.(int)
		if se != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, se, es[i])
		}
		i++
	}
}
