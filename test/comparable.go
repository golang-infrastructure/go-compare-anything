package test

type ComparableTest struct {
}

//var _ compare_anything.Comparable[ComparableTest] = &ComparableTest{}

func (x *ComparableTest) CompareTo(target *ComparableTest) int {
	return -1
}
