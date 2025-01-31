package span

type spanOptions[T any] struct {
	Span[T]
}

// Option is the type used to provide additional options during Span creation.
type Option[T any] func(*spanOptions[T])

// WithComparator is the Option used to provide a custom comparator during Span
// creation.
func WithComparator[T any](comparator Comparator[T]) Option[T] {
	return func(o *spanOptions[T]) {
		o.comparator = comparator
	}
}
