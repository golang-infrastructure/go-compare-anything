package compare_anything

// 这个文件中存放的是一些类型约束之类的

// Signed 有符号整数
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned 无符号整数
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer 整数
type Integer interface {
	Signed | Unsigned
}

// Float 浮点数
type Float interface {
	~float32 | ~float64
}

// Complex 复数
type Complex interface {
	~complex64 | ~complex128
}

// Ordered 有顺序的类型
type Ordered interface {
	Integer | Float | ~string
}
