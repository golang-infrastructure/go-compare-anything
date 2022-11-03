package compare_anything

import "reflect"

type EqualFunc[T any] func(a, b T) bool

func CompareTo() {

}

// Equals 值和类型都要相等
func Equals(a, b any) bool {
	return false
}

func DeepEqual(x, y any) bool {
	//if x == nil || y == nil {
	//	return x == y
	//}
	//v1 := ValueOf(x)
	//v2 := ValueOf(y)
	//if v1.Type() != v2.Type() {
	//	return false
	//}
	//return deepValueEqual(v1, v2, make(map[visit]bool))
	return reflect.DeepEqual(x, y)
}

// TypeEquals 类型是否匹配
func TypeEquals() {

}

// ValueEquals 值是否匹配
func ValueEquals() {

}
