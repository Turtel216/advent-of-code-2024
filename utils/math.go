package utils

// Custom constraint for all number types (integers and floats)
type number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

func Abs[T number](x T) T {
	if x < 0 {
		return -x
	}

	return x
}
