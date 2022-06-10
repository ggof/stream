package stream

// Stream represents the action to
type Stream[T any] interface {
	Next() (T, bool)
}

type Predicate[T any] func(T) bool
type Mapper[I, O any] func(I) O

type StreamFunc[T any] func() (T, bool)

func (s StreamFunc[T]) Next() (T, bool) { return s() }

// OfArray creates a new Stream from an already existing array.
func OfArray[T any](arr []T) Stream[T] {
	size := len(arr)
	i := 0

	return StreamFunc[T](func() (T, bool) {
		if i < size {
			i++
			return arr[i-1], true
		}

		return *new(T), false
	})
}

func Map[I, O any](stream Stream[I], transform func(I) O) Stream[O] {
	return StreamFunc[O](func() (O, bool) {
		next, ok := stream.Next()

		if !ok {
			return *new(O), false
		}

		return transform(next), ok
	})
}

func Filter[T any](stream Stream[T], predicate func(T) bool) Stream[T] {
	return StreamFunc[T](func() (T, bool) {
		next, ok := stream.Next()
		for ok && !predicate(next) {
			next, ok = stream.Next()
		}

		return next, ok

	})
}

// Skip skips the first n elements from the stream.
func Skip[T any](stream Stream[T], n int) Stream[T] {
	return StreamFunc[T](func() (T, bool) {
		for n > 0 {
			n--
			stream.Next()
		}

		return stream.Next()
	})
}

// SkipWhile skips elements from the stream for as long as they match predicate.
func SkipWhile[T any](stream Stream[T], predicate Predicate[T]) Stream[T] {
	var done bool

	return StreamFunc[T](func() (T, bool) {
		next, ok := stream.Next()
		for !done && ok && predicate(next) {
			next, ok = stream.Next()
		}

		done = true
		return stream.Next()
	})
}

// Take returns the first n elements from the stream.
func Take[T any](stream Stream[T], n int) Stream[T] {
	return StreamFunc[T](func() (T, bool) {
		if n > 0 {
			n--
			return stream.Next()
		}

		return *new(T), false
	})
}

// TakeWhile returns elements from stream while they match predicate.
func TakeWhile[T any](stream Stream[T], predicate Predicate[T]) Stream[T] {
	done := false
	return StreamFunc[T](func() (T, bool) {
		next, ok := stream.Next()
		if !done && ok && predicate(next) {
			return stream.Next()
		}

		done = true
		return *new(T), false
	})
}

// Reduce consumes the stream, using reducer and accumulator to transform the whole stream into accumulator.
func Reduce[T, A any](stream Stream[T], accumulator A, reducer func(A, T) A) A {
	next, ok := stream.Next()

	for ok {
		accumulator = reducer(accumulator, next)
		next, ok = stream.Next()
	}

	return accumulator
}

// ToArray applies every transformation that were registered on this stream, yielding an array.
func ToArray[T any](stream Stream[T]) []T {
	return Reduce(stream, []T{}, func(acc []T, next T) []T { return append(acc, next) })
}

//Foreach consumes the stream, running f for every element.
func ForEach[T any](stream Stream[T], f func(T)) {
	next, ok := stream.Next()

	for ok {
		f(next)
		next, ok = stream.Next()
	}
}
