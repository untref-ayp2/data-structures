package stack

import "testing"

func TestPushPop(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Pop()
	if err != nil || v != 3 {
		t.Error("Error en Push")
	}

	v, err = s.Pop()
	if err != nil || v != 2 {
		t.Error("Error en Push")
	}

	v, err = s.Pop()
	if err != nil || v != 1 {
		t.Error("Error en Push")
	}

	_, err = s.Pop()
	if err == nil {
		t.Error("Error en Push")
	}
}
