package gox

import "testing"

func TestPriorityQueueCapacity(t *testing.T) {
	q := PriorityQueue[int]{
		capacity: 5,
	}

	if q.Capacity() != 5 {
		t.Error("incorrect capacity")
	}
}

func TestPriorityQueueLength(t *testing.T) {
	q := PriorityQueue[int]{
		capacity: 5,
		items:    []int{0, 1, 2},
	}

	if q.Length() != 3 {
		t.Error("incorrect length")
	}
}

func TestPriorityQueueSlice(t *testing.T) {
	q := PriorityQueue[int]{
		capacity: 5,
		items:    []int{0, 1, 2},
	}

	slice := q.Slice()

	if len(slice) != 3 {
		t.Error("incorrect length")
	}

	if slice[0] != 0 && slice[len(slice)-1] != 2 {
		t.Error("incorrect data")
	}

	q.items[0] = -1
	if slice[0] == q.items[0] {
		t.Error("modifying original changed new")
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	q := PriorityQueue[int]{
		capacity: 5,
		items:    []int{0, 1, 2},
	}

	elems := q.Peek(2)

	if len(elems) != 2 {
		t.Error("wrong length")
	}

	if elems[0] != 0 {
		t.Error("incorrect first element")
	} else if elems[1] != 1 {
		t.Error("incorrect second element")
	}

	q.items[1] = -1
	if elems[1] == q.items[1] {
		t.Error("modifying original changed new")
	}
}

func TestPriorityQueuePop(t *testing.T) {
	q := PriorityQueue[int]{
		capacity: 5,
		items:    []int{0, 1, 2},
	}

	elems := q.Pop(2)

	if len(elems) != 2 {
		t.Error("wrong length returned")
	} else if elems[0] != 0 {
		t.Error("incorrect first element")
	} else if elems[1] != 1 {
		t.Error("incorrect second element")
	}

	if len(q.items) != 1 {
		t.Error("original did not resize")
	} else if q.items[0] != 2 {
		t.Error("original did not shift")
	}
}

func TestPriorityQueuePush(t *testing.T) {
	q := NewPriorityQueue[int](5)

	// Starting empty
	disp := q.Push(1, 2, 3)
	if disp != 0 {
		t.Error("says there was displaced")
	}
	if q.items[0] != 1 || q.items[1] != 2 || q.items[2] != 3 {
		t.Error("wrong data in queue")
	}

	// Overlapping by 2
	disp = q.Push(4, 5, 6, 7)
	if disp != 2 {
		t.Error("incorrect number displaced")
	}
	if q.items[0] != 3 || q.items[1] != 4 || q.items[2] != 5 || q.items[3] != 6 || q.items[4] != 7 {
		t.Error("wrong data in queue", q.items)
	}

	// More pushes than capacity
	q = NewPriorityQueue[int](2)
	q.Push(1, 2, 3)
	if len(q.items) != 2 {
		t.Error("incorrect length", len(q.items))
	} else if q.items[0] != 2 || q.items[1] != 3 {
		t.Error("incorrect data", q.items)
	}
}
