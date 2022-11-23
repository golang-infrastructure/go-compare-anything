package compare_anything

import "reflect"

// Comparable 表示是可比较的
type Comparable interface {

	// CompareTo 用来做比较的方法
	CompareTo(target any) int
}

// IsComparable 判断是否是可比较类型
func IsComparable(v any) bool {
	reflectValue := reflect.ValueOf(v)
	if !reflectValue.CanInterface() {
		return false
	}
	_, ok := reflectValue.Interface().(Comparable)
	return ok
}

// CastToComparable 把给定的值转为Comparable接口类型，如果无法转换则会返回error
func CastToComparable(value any) (Comparable, error) {
	reflectValue := reflect.ValueOf(value)
	if !reflectValue.CanInterface() {
		return nil, ErrCastFailed
	}
	v, ok := reflectValue.Interface().(Comparable)
	if !ok {
		return nil, ErrCastFailed
	}
	return v, nil
}
