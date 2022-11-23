能够让任何类型参与比较，而不是仅仅只是comparable

Go语言内置的比较支持：



| Type      | Comparable | Ordered | Description                                                  |
| --------- | ---------- | ------- | ------------------------------------------------------------ |
| Boolean   | ✅          | ❌       |                                                              |
| Integer   | ✅          | ✅       |                                                              |
| Float     | ✅          | ✅       |                                                              |
| Complex   | ✅          | ❌       | 分别比较实数和虚数，同时相等则两个复数相等。 如果需要比较大小，需要开发者分别比较实数和虚数。 |
| String    | ✅          | ✅       | 基于字节逐个比较。                                           |
| Pointer   | ✅          | ❌       | 如果两个指针指向同一个对象或者都为 nil，则两者相等。         |
| Channel   | ✅          | ❌       | 类似 Pointer，两个 Channel 变量只有都为 nil，或者指向同一个 Channel 的时候才相等。 |
| Interface | ✅          | ❌       | 两个 interface 的 Type 和 Value 值同时相等时，两者才相等。   |
| Struct    | ⚠️          | ❌       | 仅当 Struct 内所有成员都是 Comparable，这个 Struct 才是 Comparable 的。 如果两个 struct 类型相同，且所有非空成员变量都相等，则两者相等。 |
| Array     | ⚠️          | ❌       | 仅当成员为 Comparable，Array 才是 Comparable 的。 如果两个 Array 中的每一个元素一一相等时，则两个 Array 相等。 |
| Map       | ❌          | ❌       |                                                              |
| Slice     | ❌          | ❌       |                                                              |
| Func      | ❌          | ❌       |                                                              |



对原生的功能做扩展，提供更强大的比较支持：

- 大小
- 相等 









