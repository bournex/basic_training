package base

import "errors"

type Comparable interface {
	Compare(interface{}) int
	ToString() string
}

var (
	ErrNoFound = errors.New("no found")
	ErrEmpty   = errors.New("empty")
	ErrFull    = errors.New("full")
)
