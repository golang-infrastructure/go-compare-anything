package compare_anything

type Comparable interface {
	CompareTo(target any) int
}
