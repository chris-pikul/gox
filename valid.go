package gox

// Validable is an interface describing a type that can be checked for boolean
// validity.
type Validable interface {
	// Valid returns true if this object is considered valid, and false otherwise.
	Valid() bool
}

// PresentAndValid returns true if the given object is not-nil and passes the
// [Valid] check.
func PresentAndValid(obj Validable) bool {
	if obj == nil {
		return false
	}
	return obj.Valid()
}

// OptionalAndValid returns true if the given object is nil, if it is not it
// returns the [Valid] results.
func OptionalAndValid(obj Validable) bool {
	if obj != nil {
		return obj.Valid()
	}
	return true
}

// AllPresentAndValid runs [PresentAndValid] on variadic arguments to check that
// all objects are not-nil and pass the [Valid] check.
func AllPresentAndValid(objs ...Validable) bool {
	for _, v := range objs {
		if !PresentAndValid(v) {
			return false
		}
	}
	return true
}

// AllOptionalAndValid runs [OptionalAndValid] on variadic arguments to check
// that all objects are either [Valid] or nil.
func AllOptionalAndValid(objs ...Validable) bool {
	for _, v := range objs {
		if !OptionalAndValid(v) {
			return false
		}
	}
	return true
}
