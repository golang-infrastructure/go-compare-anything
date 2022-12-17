package compare_anything

import "reflect"

// Comparable 表示是可比较的
type Comparable[T any] interface {

	// CompareTo 用来做比较的方法
	CompareTo(target T) int
}

// IsComparable 判断是否是可比较类型
func IsComparable[T any](v T) bool {
	reflectValue := reflect.ValueOf(v)
	if !reflectValue.CanInterface() {
		return false
	}
	_, ok := reflectValue.Interface().(Comparable[T])
	return ok
}

// CastToComparable 把给定的值转为Comparable接口类型，如果无法转换则会返回error
func CastToComparable[T any](value T) (Comparable[T], error) {
	reflectValue := reflect.ValueOf(value)
	if !reflectValue.CanInterface() {
		return nil, ErrCastFailed
	}
	v, ok := reflectValue.Interface().(Comparable[T])
	if !ok {
		return nil, ErrCastFailed
	}
	return v, nil
}
