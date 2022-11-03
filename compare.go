package compare_anything

//import (
//	"errors"
//	"reflect"
//)
//
//// Compare 比较大小
//func Compare(a, b any) int {
//	// TODO
//
//	reflect.TypeOf("").PkgPath()
//
//	return -1
//}
//
//func GenComparatorFor[T any](value T) (Comparator[T], error) {
//	// TODO
//	return nil, nil
//}
//
//func GenComparatorFromValue() {
//
//}
//
//func GenComparatorFromType[T any](value T, reflectType reflect.Type) (Comparator[T], error) {
//
//}
//
//// GenComparatorFromKind 根据类型来决定用哪个比较器
//func GenComparatorFromKind[T any](kind reflect.Kind) (Comparator[T], error) {
//	switch kind {
//	case reflect.Invalid:
//		return nil, errors.New("")
//	case reflect.Bool:
//		return Comparator[T](BoolComparator()), nil
//	case reflect.Int:
//		return Comparator[T](IntComparator()), nil
//	case reflect.Int8:
//		return Comparator[T](Int8Comparator()), nil
//	case reflect.Int16:
//		return Comparator[T](Int16Comparator()), nil
//	case reflect.Int32:
//		return Comparator[T](Int32Comparator()), nil
//	case reflect.Int64:
//		return Comparator[T](Int64Comparator()), nil
//	case reflect.Uint:
//		return Comparator[T](UIntComparator()), nil
//	case reflect.Uint8:
//		return Comparator[T](UInt8Comparator()), nil
//	case reflect.Uint16:
//		return Comparator[T](UInt16Comparator()), nil
//	case reflect.Uint32:
//		return Comparator[T](UInt32Comparator()), nil
//	case reflect.Uint64:
//		return Comparator[T](UInt64Comparator()), nil
//	case reflect.Uintptr:
//		// TODO
//	case reflect.Float32:
//		return Comparator[T](Float32Comparator()), nil
//	case reflect.Float64:
//		return Comparator[T](Float64Comparator()), nil
//	case reflect.Complex64:
//	case reflect.Complex128:
//		// 复数无法比较大小
//		return nil, errors.New("")
//	case reflect.Array:
//		return Comparator[T](ArrayComparator[T]()), nil
//	case reflect.Chan:
//	case reflect.Func:
//	case reflect.Interface:
//	case reflect.Map:
//	case reflect.Pointer:
//		// 继续访问该指针指向的内容
//		elemType := reflectType.Elem()
//		if elemType != nil {
//			return GenComparatorFromType(nil, elemType)
//		}
//	case reflect.Slice:
//		return Comparator[T](SliceComparator()), nil
//	case reflect.String:
//		return Comparator[T](StringComparator()), nil
//	case reflect.Struct:
//
//	case reflect.UnsafePointer:
//
//	}
//
//	// 不支持的类型，无法生成比较器
//	return nil, errors.New("")
//}
