package priority_queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	n1 := NewNode(1, 1, 2)
	n2 := NewNode(2, 2, 3)
	n3 := NewNode(3, 3, 4)
	pq.Push(n1)
	pq.Push(n2)
	pq.Push(n3)

	v := pq.Pop()
	if v.GetKey().(int) != 1 {
		t.Fatal()
	}
	v = pq.Pop()
	if v.GetKey().(int) != 2 {
		t.Fatal()
	}
	v = pq.Pop()
	if v.GetKey().(int) != 3 {
		t.Fatal()
	}
	v = pq.Pop()
	if v != nil {
		t.Fatal()
	}

	pq.Push(n1)
	pq.Push(n2)
	pq.Push(n3)

	pq.Remove(n2.GetIndex())
	pq.Remove(3)

	v = pq.Pop()
	if v.GetKey().(int) != 1 {
		t.Fatal()
	}
	v = pq.Pop()
	if v.GetKey().(int) != 3 {
		t.Fatal()
	}
	v = pq.Pop()
	if v != nil {
		t.Fatal()
	}
}
