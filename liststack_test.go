package stack

import "testing"

func TestStack(t *testing.T) {
	s1 := New()

	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	checkStack(t, s1, []interface{}{3, 2, 1})
}

func checkStack(t *testing.T, s *Stack, es []interface{}) {

	i := 0
	for e := s.Pop(); e != nil; e = s.Pop() {
		se := e.(int)
		if se != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, se, es[i])
		}
		i++
	}
}
