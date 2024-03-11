package gox

// Ternary is a helper to do ternary type operations where you just want to
// return one or the other value based on a condition without needing to set a
// var in the scope block.
//
// Basically, if you ever went for a ternary to set a value and thought "damn,
// why doesn't go have ternaries again?", then just use this.
func Ternary[Type any](cond bool, t Type, f Type) Type {
	if cond {
		return t
	}
	return f
}

// MustAssert runs the assertion operation on the given object and panics if the
// results are false.
func MustAssert[T any](obj any) T {
	res, ok := obj.(T)
	if !ok {
		panic("could not assert object")
	}
	return res
}

// MakeAny takes a generic argument and makes one, it does this by dereferencing
// the "new" results using the type.
func MakeAny[T any]() T {
	return *(new(T))
}
