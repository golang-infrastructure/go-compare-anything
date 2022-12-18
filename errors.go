package compare_anything

import "errors"

// ErrCastFailed 类型转换失败
var ErrCastFailed = errors.New("type cast failed")

// ErrCanNotGenerateComparator 无法生成比较器
var ErrCanNotGenerateComparator = errors.New("can not generate comparator")
