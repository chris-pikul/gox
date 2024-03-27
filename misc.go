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

// AnyNil returns true if any of the variadic arguments is nil
func AnyNil(opts ...any) bool {
	for _, v := range opts {
		if v == nil {
			return true
		}
	}
	return false
}

// AnyNotNil returns true if any item in the variadic arguments is NOT nil.
func AnyNotNil(opts ...any) bool {
	for _, v := range opts {
		if v != nil {
			return true
		}
	}
	return false
}

// AllNil returns true if all the variadic arguments are nil.
func AllNil(opts ...any) bool {
	return !AnyNotNil(opts...)
}

// AllNotNil returns true if all the variadic arguments are NOT nil.
func AllNotNil(opts ...any) bool {
	return !AnyNil(opts...)
}

// AllMatchNil returns true if the nil-ability of all the given variadic arguments
// matches. Basically this is a way to say "if any of these are not nil, they must
// all be not nil", and the inverse of that statement, "if any of these are nil, they
// must all be nil".
func AllMatchNil(opts ...any) bool {
	if AnyNil(opts...) {
		return AllNil(opts...)
	}
	return true
}
