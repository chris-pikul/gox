package gox

import (
	"reflect"
	"unsafe"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 | byte
}

type Uinteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Uinteger | Float
}

type Primitive interface {
	bool | Number | string
}

func IsPrimitive(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Bool, reflect.String:
		return true
	}
	return false
}

// ForceCast is a DANGEROUS function. It uses [unsafe] to force cast the incoming
// value to the provided generic one. Very bad, no good, last resort.
func ForceCast[Type any](value any) Type {
	return *(*Type)(unsafe.Pointer(&value))
}
