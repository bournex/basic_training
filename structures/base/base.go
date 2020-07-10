package base

type Comparable interface {
	Compare(interface{}) int
	ToString() string
}
