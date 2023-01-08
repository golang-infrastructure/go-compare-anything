package compare_anything

import (
	"github.com/golang-infrastructure/go-gtypes"
	"math"
	"strings"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// Comparator 比较器的接口定义，一个类型要想参与比较必须先实现比较器
type Comparator[T any] func(a, b T) int

// ------------------------------------------------ ---------------------------------------------------------------------

// OrderedComparator 为支持排序的类型生成比较器
func OrderedComparator[T gtypes.Ordered]() Comparator[T] {
	return func(a, b T) int {
		if a == b {
			return 0
		} else if a < b {
			return -1
		} else {
			return +1
		}
	}
}
//
//// ReverseOrderedComparator 对可比较类型逆序比较的比较器，会把较大的放在前面
//func ReverseOrderedComparator[T gtypes.Ordered]() Comparator[T] {
//	return ReverseComparator[T](OrderedComparator[T]())
//}

// ------------------------------------------------ ---------------------------------------------------------------------
//
//// OrderedPointerComparator 为支持排序的类型生成比较器
//func OrderedPointerComparator[T *gtypes.Ordered]() Comparator[T] {
//	return func(a, b T) int {
//		// 1. 先比较指针
//		if a == nil && b != nil {
//			return -1
//		} else if a != nil && b == nil {
//			return 1
//		} else if a == nil && b == nil {
//			return 0
//		}
//		// 2. 再比较实际的值
//		aValue := *a
//		bValue := *b
//		if aValue == bValue {
//			return 0
//		} else if aValue < bValue {
//			return -1
//		} else {
//			return +1
//		}
//	}
//}
//
//// ReverseOrderedPointerComparator 对可比较类型逆序比较的比较器，会把较大的放在前面
//func ReverseOrderedPointerComparator[T *gtypes.Ordered]() Comparator[T] {
//	return ReverseComparator[T](OrderedPointerComparator[T]())
//}

// ------------------------------------------------ ---------------------------------------------------------------------

//// ReverseComparator 将比较结果按原点取反以达到逆序的效果
//func ReverseComparator[T any](comparator Comparator[T]) Comparator[T] {
//	return func(a, b T) int {
//		return comparator(a, b) * -1
//	}
//}

// ------------------------------------------------ ---------------------------------------------------------------------

// StringComparator 比较字符串
func StringComparator() Comparator[string] {
	return strings.Compare
}

// ------------------------------------------------ ---------------------------------------------------------------------

func IntComparator() Comparator[int] {
	return OrderedComparator[int]()
}

func Int8Comparator() Comparator[int8] {
	return OrderedComparator[int8]()
}

func Int16Comparator() Comparator[int16] {
	return OrderedComparator[int16]()
}

func Int32Comparator() Comparator[int32] {
	return OrderedComparator[int32]()
}

func Int64Comparator() Comparator[int64] {
	return OrderedComparator[int64]()
}

// ------------------------------------------------ ---------------------------------------------------------------------

func UIntComparator() Comparator[uint] {
	return OrderedComparator[uint]()
}

func UInt8Comparator() Comparator[uint8] {
	return OrderedComparator[uint8]()
}

func UInt16Comparator() Comparator[uint16] {
	return OrderedComparator[uint16]()
}

func UInt32Comparator() Comparator[uint32] {
	return OrderedComparator[uint32]()
}

func UInt64Comparator() Comparator[uint64] {
	return OrderedComparator[uint64]()
}

// ------------------------------------------------ ---------------------------------------------------------------------

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

// ------------------------------------------------ ---------------------------------------------------------------------

// BoolComparator 布尔类型的比较器，默认是让false排在前面
func BoolComparator() Comparator[bool] {
	return func(a bool, b bool) int {
		if a == b {
			return 0
		} else if !a && b {
			return -1
		} else {
			return 1
		}
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

//// OrderedArrayComparator 承载可比较元素的数组或切片使用的比较器
//func OrderedArrayComparator[T gtypes.Ordered]() Comparator[[]T] {
//	return func(a, b []T) int {
//		for index := 0; ; index++ {
//			// 对应位置先没元素的认为更小
//
//			// 如果都有元素的话，则看元素的值
//			if index >= len(a) {
//
//			}
//		}
//	}
//}

// ------------------------------------------------ ---------------------------------------------------------------------

//func SliceComparator[T any]() Comparator[T] {
//	return func(a T, b T) int {
//
//		// a和b都是切片，通过反射来获取它们的值
//		reflectA := reflect.ValueOf(a)
//		reflectB := reflect.ValueOf(b)
//		if !reflectA.IsValid() && reflectB.IsValid() {
//			return 0
//		} else if !reflectA.IsValid() {
//			return -1
//		} else if !reflectB.IsValid() {
//			return 1
//		}
//
//		// 两个切片都是可用的，那就开始挨个比较吧
//		//for index := 0; ; index++ {
//		//	valueA := reflectA.Index(index)
//		//	kind, err := GenComparatorFromKind(valueA.Kind())
//		//
//		//}
//
//		// 对应位置先没元素的认为更小
//
//		// 如果都有元素的话，则看元素的值，为元素生成比较器来比较
//
//		return 0
//	}
//}

// ------------------------------------------------ ---------------------------------------------------------------------
