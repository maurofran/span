package span

import (
	"cmp"
	"fmt"
	"time"
)

type Span[T any] struct {
	start      T
	end        T
	comparator Comparator[T]
}

// Of creates a new Span with only the provided value.
func Of[T cmp.Ordered](value T, options ...Option[T]) Span[T] {
	span, _ := Between(value, value, options...)
	return span
}

// Between creates a new Span for provided start and end values.
// It returns ErrInvalidSpan if the start value is greater than end value.
// It returns ErrInvalidComparator if an option to change the comparator is
// provided but holds a nil value.
func Between[T cmp.Ordered](start, end T, options ...Option[T]) (Span[T], error) {
	opts := spanOptions[T]{Span[T]{
		start:      start,
		end:        end,
		comparator: OrderedComparator[T]{},
	}}
	for _, opt := range options {
		opt(&opts)
	}
	result := opts.Span
	if result.comparator == nil {
		return Span[T]{}, fmt.Errorf("%w: no comparator provided", ErrInvalidComparator)
	}
	if result.comparator.Less(result.end, result.start) {
		return Span[T]{}, fmt.Errorf("%w: start is greater than end", ErrInvalidSpan)
	}
	return result, nil
}

// OfTime creates a new Span of time.Time values with only the provided value.
func OfTime(value time.Time, options ...Option[time.Time]) Span[time.Time] {
	span, _ := BetweenTimes(value, value, options...)
	return span
}

// BetweenTimes creates a new Span of time.Time values with start and end values.
// Returns ErrInvalidSpan if the provided start time is greater than the end time.
// It returns ErrInvalidComparator if an option to change the comparator is
// provided but holds a nil value.
func BetweenTimes(start, end time.Time, options ...Option[time.Time]) (Span[time.Time], error) {
	opts := spanOptions[time.Time]{Span[time.Time]{
		start:      start,
		end:        end,
		comparator: TimeComparator{},
	}}
	for _, opt := range options {
		opt(&opts)
	}
	result := opts.Span
	if result.comparator == nil {
		return Span[time.Time]{}, fmt.Errorf("%w: no comparator provided", ErrInvalidComparator)
	}
	if result.comparator.Less(result.end, result.start) {
		return Span[time.Time]{}, fmt.Errorf("%w: start is greater than end", ErrInvalidSpan)
	}
	return result, nil
}

// New creates a new generic Span with provided start and end values and with
// the provided comparator.
// Returns ErrInvalidSpan if the provided start time is greater than the end time.
// It returns ErrInvalidComparator if an option to change the comparator is
// provided but holds a nil value.
func New[T any](start, end T, comparator Comparator[T]) (Span[T], error) {
	if comparator == nil {
		return Span[T]{}, fmt.Errorf("%w: no comparator provided", ErrInvalidComparator)
	}
	if comparator.Less(end, start) {
		return Span[T]{}, fmt.Errorf("%w: start is greater than end", ErrInvalidSpan)
	}
	return Span[T]{
		start:      start,
		end:        end,
		comparator: comparator,
	}, nil
}

// Start returns the start value of the span.
func (r Span[T]) Start() T {
	return r.start
}

// End returns the end value of the span.
func (r Span[T]) End() T {
	return r.end
}

// Contains Test if the value T is contained in receiver Span.
func (r Span[T]) Contains(value T) bool {
	if r.isZero(value) {
		return false
	}
	return !r.comparator.Less(value, r.start) && !r.comparator.Less(r.end, value)
}

// ContainsSpan check if the receiver Span contains the provided one.
func (r Span[T]) ContainsSpan(other Span[T]) bool {
	return r.Contains(other.start) && r.Contains(other.end)
}

// Fit test if the value is within the Span. If it's lower return the Start
// value of the Span; if it's greater return the End value of the Span, in other
// cases returns the value itself.
func (r Span[T]) Fit(value T) T {
	if r.After(value) {
		return r.start
	}
	if r.Before(value) {
		return r.end
	}
	return value
}

// Intersect the receiver Span with the provided one, creating a new Span with
// the Start and End value resulting from the intersection.
func (r Span[T]) Intersect(other Span[T]) Span[T] {
	if r.Equal(other) {
		return r
	}
	start := r.start
	if r.comparator.Less(r.start, other.start) {
		start = other.start
	}
	end := other.end
	if r.comparator.Less(other.end, r.end) {
		end = r.end
	}
	return Span[T]{
		start:      start,
		end:        end,
		comparator: r.comparator,
	}
}

// After test if the receiver Span is after the provided value.
func (r Span[T]) After(value T) bool {
	if r.isZero(value) {
		return false
	}
	return r.comparator.Less(value, r.start)
}

// AfterSpan test if the receiver Span is after the provided one.
func (r Span[T]) AfterSpan(other Span[T]) bool {
	return r.After(other.end)
}

// Before test if the receiver Span is before the provided value.
func (r Span[T]) Before(value T) bool {
	if r.isZero(value) {
		return false
	}
	return r.comparator.Less(r.end, value)
}

// BeforeSpan test if the receiver Span is before the provided one.
func (r Span[T]) BeforeSpan(other Span[T]) bool {
	return r.Before(other.start)
}

// StartedBy test if the receiver Span is started by provided value.
func (r Span[T]) StartedBy(value T) bool {
	if r.isZero(value) {
		return false
	}
	return r.comparator.Equal(value, r.start)
}

// EndedBy test if the receiver Span is ended by provided value.
func (r Span[T]) EndedBy(value T) bool {
	if r.isZero(value) {
		return false
	}
	return r.comparator.Equal(value, r.end)
}

// OverlappedBy test if receiver Span is overlapped by provided one.
func (r Span[T]) OverlappedBy(other Span[T]) bool {
	return other.Contains(r.start) || other.Contains(r.end) || r.Contains(other.start)
}

// IsZero test if the receiver Span is the zero value.
func (r Span[T]) IsZero() bool {
	return r.isZero(r.start) && r.isZero(r.end)
}

// Equal test if the receiver Span is equal to another one.
func (r Span[T]) Equal(other Span[T]) bool {
	return r.comparator.Equal(r.start, other.start) && r.comparator.Equal(r.end, other.end)
}

// String implements the fmt.Stringer interface.
func (r Span[T]) String() string {
	return fmt.Sprintf("[%s..%s]", r.start, r.end)
}

func (r Span[T]) isZero(value T) bool {
	var zero T
	return r.comparator.Equal(value, zero)
}
