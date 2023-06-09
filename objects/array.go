package objects

import "sync"

type Predicate[T comparable] func(value T) bool
type Callback[T comparable] func(value T, index int)

type Array[T comparable] struct {
	mutex sync.Mutex
	data  []T
}

func (a *Array[T]) First() T {
	return a.data[0]
}

func (a *Array[T]) Last() T {
	return a.data[a.Length()-1]
}

func (a *Array[T]) At(i int) T {
	return a.data[i]
}

func (a *Array[T]) Get(i int) T {
	return a.At(i)
}

func (a *Array[T]) Take(i int) T {
	value := a.At(i)
	a.Splice(i, 1)
	return value
}

func (a *Array[T]) Push(v ...T) int {
	a.data = append(a.data, v...)
	return a.Length() - 1
}

func (a *Array[T]) Pop() T {
	var value T

	lastIndex := len(a.data)

	if lastIndex > 0 {
		lastIndex -= 1

		value = a.data[lastIndex]
		a.data = a.data[:lastIndex]
	}

	return value
}

func (a *Array[T]) Slice(start int, stop ...int) *Array[T] {
	var value []T

	if len(stop) == 0 {
		value = a.data[start:]
	} else {
		value = a.data[start:stop[0]]
	}

	return NewArray(value...)
}

func (a *Array[T]) Splice(offset int, length int, v ...T) *Array[T] {
	a.mutex.Lock()

	defer a.mutex.Unlock()

	endOffset := offset + length

	start := a.data[:offset]
	value := a.data[offset:endOffset]
	end := a.data[endOffset:]

	a.data = append(start, v...)
	a.data = append(a.data, end...)

	return NewArray(value...)
}

func (a *Array[T]) Contains(v T) bool {
	return a.IndexOf(v) >= 0
}

func (a *Array[T]) IndexOf(v T) int {
	for i, data := range a.data {
		if data == v {
			return i
		}
	}

	return -1
}

func (a *Array[T]) Length() int {
	return len(a.data)
}

func (a *Array[T]) Find(fn Predicate[T]) T {
	var value T

	for _, entry := range a.data {
		if ok := fn(entry); ok {
			value = entry
			break
		}
	}

	return value
}

func (a *Array[T]) Filter(fn Predicate[T]) []T {
	value := make([]T, 0)

	for _, entry := range a.data {
		if ok := fn(entry); ok {
			value = append(value, entry)
		}
	}

	return value
}

func (a *Array[T]) ForEach(fn Callback[T]) {
	for i, entry := range a.data {
		fn(entry, i)
	}
}

/*************************************/

func NewArray[T comparable](data ...T) *Array[T] {
	return &Array[T]{
		data: data,
	}
}
