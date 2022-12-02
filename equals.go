package compare_anything

import "reflect"

type EqualsFunc[T any] func(a, b T) bool

type Equalsable[T any] interface {
	Equals(other T) bool
}

// IsEqualsable 判断值是否是可比较的
func IsEqualsable[T any](value T) bool {
	reflectValue := reflect.ValueOf(value)
	if !reflectValue.CanInterface() {
		return false
	}
	_, ok := reflectValue.Interface().(Equalsable[T])
	return ok
}

// Equals 值和类型都要相等
func Equals(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// TypeEquals 比较两个值的类型是否匹配，忽略值
func TypeEquals(a, b any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

// ValueEquals 比较两个值是否匹配，忽略类型
func ValueEquals(a, b any) bool {
	reflectValueA := reflect.ValueOf(a)
	reflectValueB := reflect.ValueOf(b)
	if !reflectValueA.IsValid() {
		return reflectValueA.IsValid() == reflectValueB.IsValid()
	}
	return Equals(reflectValueA.Interface(), reflectValueB.Interface())
}
