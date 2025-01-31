package span

import (
	"cmp"
	"time"
)

// Comparator is the interface used to compare elements.
type Comparator[T any] interface {
	// Less returns true if the first element is less than the second.
	Less(T, T) bool
	// Equal returns true if the first element is equal to the second.
	Equal(T, T) bool
}

// OrderedComparator is the type used to compare cmp.Ordered objects.
type OrderedComparator[T cmp.Ordered] struct {
}

// Less implements the Comparator interface for cmp.Ordered elements.
func (OrderedComparator[T]) Less(a, b T) bool {
	return a < b
}

// Equal implements the Comparator interface for cmp.Ordered elements.
func (OrderedComparator[T]) Equal(a, b T) bool {
	return a == b
}

// TimeComparator is the type used to compare time.Time instances.
type TimeComparator struct{}

// Less implements the Comparator interface for time.Time.
func (TimeComparator) Less(a, b time.Time) bool {
	return a.Before(b)
}

// Equal implements the Comparator interface for time.Time.
func (TimeComparator) Equal(a, b time.Time) bool {
	return a.Equal(b)
}
