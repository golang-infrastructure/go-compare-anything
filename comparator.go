package compare_anything

import (
	"github.com/golang-infrastructure/go-gtypes"
	"math"
	"reflect"
	"strings"
	"unsafe"
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

// ReverseOrderedComparator 对可比较类型逆序比较的比较器，会把较大的放在前面
func ReverseOrderedComparator[T gtypes.Ordered]() Comparator[T] {
	return ReverseComparator[T](OrderedComparator[T]())
}

// ------------------------------------------------ ---------------------------------------------------------------------

// OrderedPointerComparator 为支持排序的类型生成比较器
func OrderedPointerComparator[T *gtypes.Ordered]() Comparator[T] {
	return func(a, b T) int {
		// 1. 先比较指针
		if a == nil && b != nil {
			return -1
		} else if a != nil && b == nil {
			return 1
		} else if a == nil && b == nil {
			return 0
		}
		// 2. 再比较实际的值
		aValue := *a
		bValue := *b
		if aValue == bValue {
			return 0
		} else if aValue < bValue {
			return -1
		} else {
			return +1
		}
	}
}

// ReverseOrderedPointerComparator 对可比较类型逆序比较的比较器，会把较大的放在前面
func ReverseOrderedPointerComparator[T *gtypes.Ordered]() Comparator[T] {
	return ReverseComparator[T](OrderedPointerComparator[T]())
}

// ------------------------------------------------ ---------------------------------------------------------------------

// ReverseComparator 将比较结果按原点取反以达到逆序的效果
func ReverseComparator[T any](comparator Comparator[T]) Comparator[T] {
	return func(a, b T) int {
		return comparator(a, b) * -1
	}
}

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

// OrderedSliceComparator 承载可比较元素的数组或切片使用的比较器
func OrderedSliceComparator[T gtypes.Ordered]() Comparator[[]T] {
	return func(sliceA, sliceB []T) int {
		for index := 0; ; index++ {

			// 如果有任意一个消费到尾部了，则对应位置先没元素的认为更小
			if index >= len(sliceA) || index >= len(sliceB) {
				return len(sliceA) - len(sliceB)
			}

			// 如果都有元素的话，则比较对应位置的元素的值，因为是可比较元素，因此直接比较就好了
			if sliceA[index] == sliceB[index] {
				continue
			} else if sliceA[index] > sliceB[index] {
				return 1
			} else {
				return -1
			}
		}
	}
}

// SliceComparator 用于比较两个切片的大小
func SliceComparator[T any]() Comparator[[]T] {
	return func(sliceA, sliceB []T) int {
		for index := 0; ; index++ {
			// 如果有任意一个消费到尾部了，则对应位置先没元素的认为更小
			if index >= len(sliceA) || index >= len(sliceB) {
				return len(sliceA) - len(sliceB)
			}

			// 如果都有元素的话，则比较对应位置的元素的值，因为是不可以直接比较的元素，因此根据元素类型自动比较
			r, err := CompareE(sliceA[index], sliceB[index])
			if err != nil {
				continue
			}
			if r != 0 {
				return r
			}
		}
	}
}

// SliceComparator2 另一种用于比较两个切片的大小的算法
func SliceComparator2[T any]() Comparator[[]T] {
	return func(sliceA, sliceB []T) int {
		// 先比较切片中的元素的个数，如果元素的个数不想等的话，谁持有的元素个数多久认为谁大
		if len(sliceA) != len(sliceB) {
			return len(sliceA) - len(sliceB)
		}
		for index := 0; index < len(sliceA); index++ {
			// 如果都有元素的话，则比较对应位置的元素的值，因为是不可以直接比较的元素，因此根据元素类型自动比较
			r, err := CompareE(sliceA[index], sliceB[index])
			if err != nil {
				continue
			}
			if r != 0 {
				return r
			}
		}
		// 如果比较完了所有元素还没有结果，则认为是相等的
		return 0
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

// StructComparator 用于比较Struct大小的比较器
func StructComparator[T any]() Comparator[T] {
	return func(a, b T) int {
		// 依此比较struct中的每个field的大小
		// TODO 优先比较Ordered的类型，然后再比较其他类型
		reflectValueA := reflect.ValueOf(a)
		reflectValueB := reflect.ValueOf(b)
		for i := 0; i < reflectValueA.NumField(); i++ {
			fieldValueA := reflectValueA.Field(i).Interface()
			fieldValueB := reflectValueB.Field(i).Interface()
			r, err := CompareE(fieldValueA, fieldValueB)
			if err != nil {
				continue
			}
			if r != 0 {
				return r
			}
		}
		// 如果都比较了还不能区分大小，则只好认为是相等了
		return 0
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

// MapComparator 用于比较Map大小
func MapComparator[K comparable, V any]() Comparator[map[K]V] {
	return func(mapA, mapB map[K]V) int {

		// 谁持有的元素数量多则认为谁大
		lenA := len(mapA)
		lenB := len(mapB)
		if lenA != lenB {
			return lenA - lenB
		}

		// 如果元素个数一样多，则依此比较键值对

	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

// ChanComparator 用于比较两个channel的大小
func ChanComparator[T any]() Comparator[chan T] {
	return func(channelA, channelB chan T) int {

		// 谁持有的元素个数多则认为谁大
		lenA := len(channelA)
		lenB := len(channelB)
		if lenA != lenB {
			return lenA - lenB
		}

	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func FuncComparator() {

}

// ------------------------------------------------ ---------------------------------------------------------------------

// InterfaceComparator 其类型为interface，则看起实际指向的值了
func InterfaceComparator() {

}

// ------------------------------------------------ ---------------------------------------------------------------------

// PtrComparator 比较指针类型
func PtrComparator[T any]() Comparator[*T] {
	return func(a, b *T) int {
		// 先比较指针，比较指针类型的话就是nil为小，都为nil则相等
		if a == nil && b != nil {
			return -1
		} else if a != nil && b == nil {
			return 1
		} else if a == nil && b == nil {
			return 0
		}
		// 都不为nil的话则再比较其指向的内容的大小了
		return Compare(*a, *b)
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

// UnsafePointerComparator 用于比较unsafe类型
func UnsafePointerComparator() Comparator[unsafe.Pointer] {
	return func(a, b unsafe.Pointer) int {
		// 先比较指针，比较指针类型的话就是nil为小，都为nil则相等
		if a == nil && b != nil {
			return -1
		} else if a != nil && b == nil {
			return 1
		} else if a == nil && b == nil {
			return 0
		}
		// 如果都不为nil的话，则直接比较其地址大小
		return int(*a - *b)
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------
