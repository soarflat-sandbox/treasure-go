package stack

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	s := &Stack{limit: 2}
	s.Push(1)
	if r := s.Pop(); r != 1 {
		t.Fatalf("want 1, got %v", r)
	}
	if r := s.Pop(); r != nil {
		t.Fatalf("want nil, got %v", r)
	}
	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	if r := s.Pop(); r != "dataC" {
		t.Fatalf("want dataC, got %v", r)
	}
	if r := s.Pop(); r != "dataB" {
		t.Fatalf("want dataB, got %v", r)
	}
	if r := s.Pop(); r != nil {
		t.Fatalf("want nil, got %v", r)
	}
}
