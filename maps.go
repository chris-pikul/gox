package gox

// CopyMap deeply copies a given map by reconstructing it over iteration.
func CopyMap[Key comparable, Value any](target map[Key]Value) (ret map[Key]Value) {
	ret = make(map[Key]Value, len(target))
	for k, v := range target {
		ret[k] = v
	}
	return
}

// MapKeys extracts all the keys of a map as a slice
func MapKeys[Key comparable, Value any](target map[Key]Value) (ret []Key) {
	ret = make([]Key, len(target))
	ind := 0
	for key := range target {
		ret[ind] = key
		ind++
	}
	return
}

// MapValues extracts all the values of a map as a slice
func MapValues[Key comparable, Value any](target map[Key]Value) (ret []Value) {
	ret = make([]Value, len(target))
	ind := 0
	for _, val := range target {
		ret[ind] = val
		ind++
	}
	return
}

// TransformMapKeys returns a new map with each key of the [target] map ran
// through the [transform] function.
//
// It does not edit in-place in order to protect against clashing key overwrites.
func TransformMapKeys[Key comparable, Value any](target map[Key]Value, transform func(Key) Key) (ret map[Key]Value) {
	ret = make(map[Key]Value)
	for key, val := range target {
		ret[transform(key)] = val
	}
	return
}

// TransformMapValues returns the same map with each value of the [target] map ran
// through the [transform] function.
//
// This operates in-place for the map since the keys are the same.
func TransformMapValues[Key comparable, Value any](target map[Key]Value, transform func(Value) Value) map[Key]Value {
	for key, val := range target {
		target[key] = transform(val)
	}
	return target
}

// JoinMaps takes a variadic amount of maps and joins them into a new one. They
// are processed in order, and any conflicting keys are overridden.
func JoinMaps[Key comparable, Value any](maps ...map[Key]Value) (ret map[Key]Value) {
	ret = make(map[Key]Value)
	for _, m := range maps {
		for k, v := range m {
			ret[k] = v
		}
	}
	return
}

// FlattenMap takes a map and turns it into a slice by laying out keys and values
// in order. Due to generic limitations it returns as any.
func FlattenMap[Key comparable, Value any](target map[Key]Value) (ret []any) {
	ret = make([]any, len(target)*2)
	ind := 0
	for k, v := range target {
		ret[ind] = k
		ret[ind+1] = v
		ind += 2
	}
	return
}

// SplitKeyValue takes a map and splits the keys and values into two separate
// slices and returns them.
func SplitKeyValue[Key comparable, Value any](target map[Key]Value) (keys []Key, vals []Value) {
	keys = make([]Key, len(target))
	vals = make([]Value, len(target))
	ind := 0
	for key, val := range target {
		keys[ind] = key
		vals[ind] = val
		ind++
	}
	return
}
