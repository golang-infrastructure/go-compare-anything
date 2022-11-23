package compare_anything

import "reflect"

type Equalsable[T any] interface {
	Equals(other T) bool
}

// Equals 值和类型都要相等
func Equals(a, b any) bool {
	return false
}

// IsEqualsable 判断值是否是可比较的
func IsEqualsable(value any) bool {
	//reflectValue := reflect.ValueOf(value)
	//if !reflectValue.CanInterface() {
	//	return false
	//}
	//reflectValue.Interface().(Equalsable)
	return false
}

func DeepEqual(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// TypeEquals 比较两个值的类型是否匹配
func TypeEquals(a, b any) bool {
	return false
}

// ValueEquals 比较两个值是否匹配
func ValueEquals(a, b any) bool {
	reflectValueA := reflect.ValueOf(a)
	reflectValueB := reflect.ValueOf(b)
	if !reflectValueA.IsValid() {
		return reflectValueA.IsValid() == reflectValueB.IsValid()
	}
	return DeepEqual(reflectValueA.Interface(), reflectValueB.Interface())
}
