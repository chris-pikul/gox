package gox

// FixedArray is a data structure for holding a fixed amount of elements. Any new
// elements pushed will evict the older elements.
//
// There is no  mutexes here as that is up to your implementation. The oldest
// element will be at index 0.
type FixedArray[Type any] struct {
	size int
	arr  []Type
	ind  int
}

func (a FixedArray[Type]) safeInd() int {
	if a.ind >= 0 && a.ind < a.size {
		return a.ind
	} else if a.ind < 0 {
		return 0
	}
	return a.size - 1
}

func (a FixedArray[Type]) endInd() int {
	end := a.ind + 1
	if end > a.size {
		end = a.size
	}
	return end
}

// Size returns the maximum size of the array
func (a FixedArray[Type]) Size() int {
	return a.size
}

// Count returns the current size of the array as it is filled and will never
// exceed [Size].
func (a FixedArray[Type]) Count() int {
	return a.ind
}

// IsFull returns true if the array is considered full
func (a FixedArray[Type]) IsFull() bool {
	return a.ind >= a.size
}

// Reset clears the array
func (a *FixedArray[Type]) Reset() {
	a.arr = a.arr[:0]
	a.ind = -1
}

// Oldest returns the first element being the oldest in the array
func (a FixedArray[Type]) Oldest() *Type {
	if a.ind >= 0 {
		return &a.arr[0]
	}
	return nil
}

// Youngest returns the last element being the youngest in the array
func (a FixedArray[Type]) Youngest() *Type {
	if a.ind >= 0 {
		return &a.arr[a.safeInd()]
	}
	return nil
}

// Elements copies the internal array and returns it as a slice.
func (a FixedArray[Type]) Elements() (ret []Type) {
	ret = make([]Type, a.endInd())
	copy(ret, a.arr[:a.endInd()])
	return
}

// shift moves all the elements down in the array
func (a *FixedArray[Type]) shift(count int) {
	a.arr = a.arr[count:]
	a.ind -= count
	if a.ind < -1 {
		a.ind = -1
	}
}

func (a *FixedArray[Type]) pushElement(elem Type) {
	if a.ind >= a.size {
		a.shift(1)
		a.arr[a.size-1] = elem
		a.ind = a.size
	} else {
		a.ind++
		a.arr[a.safeInd()] = elem
	}
}

// Push adds new elements on-top of this array. They are added in the order they
// are provided.
//
// NOTE: An optimization can be made to pre-shift the whole array to the number of
// elements being added instead of individually.
func (a *FixedArray[Type]) Push(elems ...Type) {
	for _, e := range elems {
		a.pushElement(e)
	}
}

func (a FixedArray[Type]) Copy() (ret FixedArray[Type]) {
	ret.size = a.size
	ret.arr = make([]Type, ret.size)
	if a.ind >= 0 {
		copy(ret.arr, a.arr[:a.endInd()])
	}
	ret.ind = a.ind
	return
}

func (a FixedArray[Type]) MarshalJSON() ([]byte, error) {
	return JSONMarshaler(a.Elements())
}

// NewFixedArray creates an instantiates a new FixedArray of the given size.
func NewFixedArray[Type any](size int, elems ...Type) FixedArray[Type] {
	return FixedArray[Type]{
		size: size,
		arr:  make([]Type, size),
		ind:  -1,
	}
}
