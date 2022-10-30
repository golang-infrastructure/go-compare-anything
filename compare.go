package compare_anything

import "reflect"

// Compare 比较大小
func Compare(a, b any) int {
	// TODO

	reflect.TypeOf("").PkgPath()

	return -1
}

func GenComparatorFor[T any](value T) (Comparator[T], error) {
	// TODO
	return nil, nil
}

//func GenComparatorFromType() {
//	of := reflect.TypeOf(value)
//	switch of.Kind() {
//
//	}
//}
