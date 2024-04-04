package gox

// CopySlice returns a copy of the given slice
func CopySlice[Type any](slice []Type) []Type {
	arr := make([]Type, len(slice))
	copy(arr, slice)
	return arr
}

// PeekSliceOldest safely returns the [count] of items at the beginning of the
// slice without modifying the original. The returned slice is a copy from the
// original and should not be tied by reference to the original.
func PeekSliceOldest[Type any](slice []Type, count int) []Type {
	count = Min(count, len(slice))
	if count == 0 {
		return []Type{}
	}

	arr := make([]Type, count)
	copy(arr, slice)
	return arr
}

// PeekSlice safely returns the [count] of items at the end of the array without
// modifying the original. The returned slice is a copy from the original and
// should not be tied by reference to the original.
func PeekSlice[Type any](slice []Type, count int) []Type {
	count = Min(count, len(slice))
	if count == 0 {
		return []Type{}
	}

	arr := make([]Type, count)
	copy(arr, slice[len(slice)-count:])
	return arr
}

// ShiftSlice shifts the contents of the given slice left removing the first
// element and returning it. If the slice is empty, it returns a zero value.
func ShiftSlice[Type any](slice []Type) (remaining []Type, element Type) {
	if len(slice) == 0 {
		remaining = slice
		return
	}

	element, remaining = slice[0], slice[1:]
	return
}

// ShiftMultipleFromSlice removes the given [count] of elements from the front
// of the slice and returns them. It also returns the modified slice of the remaining
// elements. If the slice was empty, it returns an empty slice.
func ShiftMultipleFromSlice[Type any](slice []Type, count int) (remaining []Type, elements []Type) {
	count = Min(count, len(slice))
	if count == 0 {
		return slice, []Type{}
	}

	elements, remaining = slice[:count], slice[count:]
	return
}

// PopSlice removes the last element in the slice and returns it. If the slice
// is empty, it returns a zero value.
func PopSlice[Type any](slice []Type) (remaining []Type, element Type) {
	if len(slice) == 0 {
		remaining = slice
		return
	}

	element, remaining = slice[len(slice)-1], slice[:len(slice)-1]
	return
}

func PopMultipleFromSlice[Type any](slice []Type, count int) (remaining []Type, elements []Type) {
	count = Min(count, len(slice))
	if count == 0 {
		return slice, []Type{}
	}

	elements, remaining = slice[len(slice)-count:], slice[:len(slice)-count]
	return
}

func ResizeSlice[Type any](slice []Type, capacity int) []Type {
	if capacity < len(slice) {
		return slice[:capacity]
	}

	arr := make([]Type, capacity)
	copy(arr, slice)
	return arr
}

// JoinSlices combines all the values in each slice together.
func JoinSlices[Value any](slices ...[]Value) (ret []Value) {
	ret = make([]Value, 0)
	for _, s := range slices {
		ret = append(ret, s...)
	}
	return
}

// SliceContains runs a matcher on all values of a slice and returns true when
// the matching function does.
func SliceContains[Type any](slice []Type, matcher func(v Type) bool) bool {
	for _, v := range slice {
		if matcher(v) {
			return true
		}
	}
	return false
}

func SliceFindIndex[Type any](slice []Type, matcher func(v Type) bool) int {
	for i, v := range slice {
		if matcher(v) {
			return i
		}
	}
	return -1
}

// SliceFindFirst runs a matcher on all values of a slice and returns the value
// that responds with true.
func SliceFindFirst[Type any](slice []Type, matcher func(v Type) bool) (*Type, bool) {
	for _, v := range slice {
		if matcher(v) {
			return &v, true
		}
	}
	return nil, false
}

// FilterSlice returns a new slice by running a matcher on all values in the
// provided slice. If the matcher returns true, it is added to the new return
// slice.
func FilterSlice[Type any](slice []Type, matcher func(v Type) bool) (ret []Type) {
	for _, v := range slice {
		if matcher(v) {
			ret = append(ret, v)
		}
	}
	return
}

// SplitSlice runs a matcher on every entry in the supplied slice. When the matcher
// returns true it is added to a "trues" slice, when false to a "falses" slice.
// Both slices are returned.
func SplitSlice[Type any](slice []Type, matcher func(v Type) bool) (falses []Type, trues []Type) {
	for _, v := range slice {
		if matcher(v) {
			trues = append(trues, v)
		} else {
			falses = append(falses, v)
		}
	}
	return
}

// SliceAny returns true if any element in the given slice passes the matcher
// function provided.
func SliceAny[Type any](slice []Type, matcher func(v Type) bool) bool {
	for _, v := range slice {
		if matcher(v) {
			return true
		}
	}
	return false
}

// SliceEvery returns true if every element in the given slice passes the matcher
// function provided.
func SliceEvery[Type any](slice []Type, matcher func(v Type) bool) bool {
	for _, v := range slice {
		if !matcher(v) {
			return false
		}
	}
	return true
}
