package compare_anything

import (
	"math"
	"strings"
)

type Comparator[T any] func(a T, b T) int

func StringComparator() Comparator[string] {
	return strings.Compare
}

func IntComparator() Comparator[int] {
	return func(a int, b int) int {
		return a - b
	}
}

func UIntComparator() Comparator[uint] {
	return func(a uint, b uint) int {
		return int(a) - int(b)
	}
}

func Float64Comparator() Comparator[float64] {
	return func(a float64, b float64) int {
		if math.Abs(a-b) < 0.00001 {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	}
}

func Float32Comparator() Comparator[float32] {
	return func(a float32, b float32) int {
		if math.Abs(float64(a-b)) < 0.00001 {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	}
}
