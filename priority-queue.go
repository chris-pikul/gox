package gox

import "sync"

// PriorityQueue is a queue type backed by a slice and mutex locks that takes
// new items in and adds them to the end of the queue. It features a limited size
// that when exceeded will trim the contents and shift the queue to make room
// for new items. I didn't have a better name, but basically it "prioritizes"
// incoming items over older indices by shifting the slice left. It is considered
// FIFO in that the oldest items are returned first when accessing.
type PriorityQueue[Type any] struct {
	capacity int
	items    []Type
	lock     sync.RWMutex
}

// Capacity returns the maximum capacity of the slice
func (q *PriorityQueue[Type]) Capacity() int {
	return q.capacity
}

// Length returns the current length of the internal queue slice. This is thread
// safe.
func (q *PriorityQueue[Type]) Length() int {
	q.lock.Lock()
	length := len(q.items)
	q.lock.Unlock()

	return length
}

// Slice returns a copy of the internal slice. This is thread-safe.
func (q *PriorityQueue[Type]) Slice() []Type {
	q.lock.Lock()
	arr := CopySlice(q.items)
	q.lock.Unlock()

	return arr
}

// Peek returns the elements (up to [count] number) of the oldest elements in
// the queue.
func (q *PriorityQueue[Type]) Peek(count int) []Type {
	q.lock.Lock()
	arr := PeekSliceOldest(q.items, count)
	q.lock.Unlock()

	return arr
}

// Pop returns the elements (up to [count] number) of the oldest elements in
// the queue. It adjusts the queue by removing the items chosen and shifting the
// newer ones down.
func (q *PriorityQueue[Type]) Pop(count int) []Type {
	q.lock.RLock()
	defer q.lock.RUnlock()

	var items []Type
	q.items, items = ShiftMultipleFromSlice(q.items, count)
	return items
}

// Push adds items to the end of the queue. If the capacity is exceeded it shifts
// the oldest items out to make room for the new items. It returns the number
// of displaced items (if any). If more items are pushed than there is capacity
// for, then the items being pushed are also trimmed to the [Capacity] limit of
// oldest items.
func (q *PriorityQueue[Type]) Push(items ...Type) int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	// If we are over capacity do the quicker copy operation
	if len(items) >= q.capacity {
		disp := len(q.items)
		if len(q.items) < q.capacity {
			q.items = ResizeSlice(q.items, q.capacity)
		}

		items = items[len(items)-q.capacity:]
		copy(q.items, items)
		return disp
	}

	overlap := 0

	space := q.capacity - len(q.items)
	if len(items) > space {
		// Putting more items in than we have space for
		overlap = len(items) - space
		q.items = q.items[overlap:]
	}

	// Resize the length up to capacity
	q.items = ResizeSlice(q.items, Min(len(q.items)+len(items), q.capacity))

	// Copy in the new items
	copy(q.items[len(q.items)-len(items):], items)
	return overlap
}

// NewPriorityQueue constructs a new [PriorityQueue] using the given capacity.
// This capacity is used to allocate the internal slice's capacity at creation
// time.
func NewPriorityQueue[Type any](capacity int) PriorityQueue[Type] {
	capacity = Max(capacity, 1)

	return PriorityQueue[Type]{
		capacity: capacity,
		items:    make([]Type, 0, capacity),
		lock:     sync.RWMutex{},
	}
}
