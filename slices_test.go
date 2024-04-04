package gox

import "testing"

func TestCopySlice(t *testing.T) {
	orig := []int{0, 1, 2, 3}
	new := CopySlice(orig)
	orig[0] = -1

	if new[0] == orig[0] {
		t.Error("modifying original hit the new one")
	}
}

func TestPeekSlideOldest(t *testing.T) {
	orig := []int{0, 1, 2, 3, 4}

	sample := PeekSliceOldest(orig, 2)
	if len(sample) != 2 {
		t.Error("incorrect length returned")
	} else if sample[0] != orig[0] || sample[1] != orig[1] {
		t.Error("incorrect values returned")
	}

	orig[0] = -1
	if sample[0] == orig[0] {
		t.Error("not a copy")
	}

	sample = PeekSliceOldest(orig, len(orig)+2)
	if len(sample) != len(orig) {
		t.Error("count did not clamp to min")
	}

	sample = PeekSliceOldest([]int{}, 1)
	if len(sample) != 0 {
		t.Error("returned something weird")
	}
}

func TestPeekSlice(t *testing.T) {
	orig := []int{0, 1, 2, 3, 4}

	sample := PeekSlice(orig, 2)
	if len(sample) != 2 {
		t.Error("incorrect length returned")
	} else if sample[0] != orig[len(orig)-2] || sample[1] != orig[len(orig)-1] {
		t.Error("incorrect values returned")
	}

	orig[0] = -1
	if sample[0] == orig[0] {
		t.Error("not a copy")
	}

	sample = PeekSlice(orig, len(orig)+2)
	if len(sample) != len(orig) {
		t.Error("count did not clamp to min")
	}

	sample = PeekSlice([]int{}, 1)
	if len(sample) != 0 {
		t.Error("returned something weird")
	}
}

func TestShiftSlice(t *testing.T) {
	_, elem := ShiftSlice([]int{})
	if elem != 0 {
		t.Error("did not return zero")
	}

	orig := []int{0, 1, 2, 3, 4}

	orig, elem = ShiftSlice(orig)
	if len(orig) != 4 {
		t.Error("did not modify original")
	} else if orig[0] != 1 {
		t.Error("did not shift original")
	} else if elem != 0 {
		t.Error("did not return correct element")
	}
}

func TestPopSlice(t *testing.T) {
	_, elem := PopSlice([]int{})
	if elem != 0 {
		t.Error("did not return zero")
	}

	orig := []int{0, 1, 2, 3, 4}

	orig, elem = PopSlice(orig)
	if len(orig) != 4 {
		t.Error("did not modify original")
	} else if orig[len(orig)-1] != 3 {
		t.Error("did not trim original")
	} else if elem != 4 {
		t.Error("did not return correct element")
	}
}

func TestShiftMultipleFromSlice(t *testing.T) {
	_, elem := ShiftMultipleFromSlice([]int{}, 2)
	if len(elem) != 0 {
		t.Error("did not return zero")
	}

	orig := []int{0, 1, 2, 3, 4}

	var elems []int
	orig, elems = ShiftMultipleFromSlice(orig, 3)
	if len(orig) != 2 {
		t.Error("did not modify original")
	} else if orig[0] != 3 {
		t.Error("did not shift original")
	} else if len(elems) != 3 || elems[len(elems)-1] != 2 {
		t.Error("did not return correct element")
	}
}

func TestPopMultipleFromSlice(t *testing.T) {
	_, elem := PopMultipleFromSlice([]int{}, 2)
	if len(elem) != 0 {
		t.Error("did not return zero")
	}

	orig := []int{0, 1, 2, 3, 4}

	var elems []int
	orig, elems = PopMultipleFromSlice(orig, 3)
	if len(orig) != 2 {
		t.Error("did not modify original")
	} else if orig[len(orig)-1] != 1 {
		t.Error("did not trim original")
	} else if len(elems) != 3 || elems[len(elems)-1] != 4 {
		t.Error("did not return correct element")
	}
}
