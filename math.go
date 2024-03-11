package gox

import (
	"reflect"
	"unsafe"
)

var (
	INT_SIZE  = unsafe.Sizeof(int(0))
	UINT_SIZE = unsafe.Sizeof(uint(0))

	MIN_INT   int   = -1 << (INT_SIZE - 1)
	MIN_INT8  int8  = -1 << 7
	MIN_INT16 int16 = -1 << 15
	MIN_INT32 int32 = -1 << 31
	MIN_INT64 int64 = -1 << 63

	MIN_FLOAT32 float32 = -1 << 31
	MIN_FLOAT64 float64 = -1 << 63

	MAX_INT   int   = 1<<(INT_SIZE-1) - 1
	MAX_INT8  int8  = 1<<7 - 1
	MAX_INT16 int16 = 1<<15 - 1
	MAX_INT32 int32 = 1<<31 - 1
	MAX_INT64 int64 = 1<<63 - 1

	MAX_UINT   uint   = 1<<(UINT_SIZE-1) - 1
	MAX_UINT8  uint8  = 1<<7 - 1
	MAX_UINT16 uint16 = 1<<15 - 1
	MAX_UINT32 uint32 = 1<<31 - 1
	MAX_UINT64 uint64 = 1<<63 - 1

	MAX_FLOAT32 float32 = 1<<31 - 1
	MAX_FLOAT64 float64 = 1<<63 - 1
)

func Min[Type Number](vals ...Type) Type {
	if len(vals) == 0 {
		return 0
	}

	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[Type Number](vals ...Type) Type {
	if len(vals) == 0 {
		return 0
	}

	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

// MinValue returns the minimum value for the given numeric type, this is the
// smallest the memory will allow.
func MinValue[Type Number]() Type {
	test := new(Type)
	typeOf := reflect.Indirect(reflect.ValueOf(test)).Type()
	switch typeOf.Kind() {
	case reflect.Int:
		return ForceCast[Type](MIN_INT)
	case reflect.Int8:
		return ForceCast[Type](MIN_INT8)
	case reflect.Int16:
		return ForceCast[Type](MIN_INT16)
	case reflect.Int32:
		return ForceCast[Type](MIN_INT32)
	case reflect.Int64:
		return ForceCast[Type](MIN_INT64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Type(0)
	case reflect.Float32:
		return ForceCast[Type](MIN_FLOAT32)
	case reflect.Float64:
		return ForceCast[Type](MIN_FLOAT64)
	}
	panic("unsupported numeric type for [utils.MinValue]")
}

// MaxValue returns the maximum value for the given numeric type, this is the
// largest the memory will allow.
func MaxValue[Type Number]() Type {
	test := new(Type)
	typeOf := reflect.Indirect(reflect.ValueOf(test)).Type()
	switch typeOf.Kind() {
	case reflect.Int:
		return ForceCast[Type](MAX_INT)
	case reflect.Int8:
		return ForceCast[Type](MAX_INT8)
	case reflect.Int16:
		return ForceCast[Type](MAX_INT16)
	case reflect.Int32:
		return ForceCast[Type](MAX_INT32)
	case reflect.Int64:
		return ForceCast[Type](MAX_INT64)
	case reflect.Uint:
		return ForceCast[Type](MAX_UINT)
	case reflect.Uint8:
		return ForceCast[Type](MAX_UINT8)
	case reflect.Uint16:
		return ForceCast[Type](MAX_UINT16)
	case reflect.Uint32:
		return ForceCast[Type](MAX_UINT32)
	case reflect.Uint64:
		return ForceCast[Type](MAX_UINT64)
	case reflect.Float32:
		return ForceCast[Type](MAX_FLOAT32)
	case reflect.Float64:
		return ForceCast[Type](MAX_FLOAT64)
	}
	panic("unsupported numeric type for [utils.MinValue]")
}
