package gox

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
