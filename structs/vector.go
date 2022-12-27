package structs

import (
	"fmt"
)

const (
	initialCapacity = 10
)

type vector struct {
	array    []any
	size     int
	capacity int
}

func NewVector() List {
	array := make([]any, initialCapacity)

	return &vector{
		array,
		0,
		initialCapacity,
	}
}

// List interface methods

func (vec *vector) InsertHead(val any) {
	if vec.size >= vec.capacity {
		vec.expand()
	}

	vec.size++
	vec.shiftRight(0)
	vec.array[0] = val
}

func (vec *vector) InsertTail(val any) {
	if vec.size >= vec.capacity {
		vec.expand()
	}

	vec.size++
	vec.array[vec.size - 1] = val
}

func (vec *vector) Insert(i int, val any) error {
}

func (vec *vector) Set(i int, val any) error {
	if i >= vec.size {
		return fmt.Errorf("Index i (%d) out of range for length %d.", i, vec.size)
	}

	if i < 0 {
		return nil, fmt.Errorf("Negative index given (%d).", i)
	}

	vec.array[i] = val
	return nil
}

func (vec *vector) GetHead() any {
	if vec.size <= 0 {
		return nil
	}

	return vec.array[0]
}

func (vec *vector) GetTail() any {
	if vec.size <= 0 {
		return nil
	}

	return vec.array[vec.size - 1]
}

func (vec *vector) Get(i int) (any, error) {
	if i >= vec.size {
		return nil, fmt.Errorf("Index i (%d) out of range for length %d.", i, vec.size)
	}

	if i < 0 {
		return nil, fmt.Errorf("Negative index given (%d).", i)
	}

	return vec.array[i], nil
}

func (vec *vector) RemoveHead() (any, bool) {
	if vec.size <= 0 {
		return nil, false
	}

	head := vec.array[0]
	vec.shiftLeft(0)
	return head, true
}

func (vec *vector) RemoveTail() (any, bool) {
	if vec.size <= 0 {
		return nil, false
	}

	tail := vec.array[vec.size - 1]
	vec.shiftRight(0)

	return head, true
}

func (vec *vector) Remove(i int) error {
	vec.shiftLeft(i)
	vec.size--

	return nil
}

func (vec *vector) Remove(val any) error {
	return nil
}

func (vec *vector) Push(val any) {
	vec.InsertTail(val)
}

func (vec *vector) Pop() (any, bool) {
	return vec.RemoveTail()
}

func (vec *vector) Size() int {
	return vec.size
}

func (vec *vector) IndexOf(val any) int {
	for i := 0; i < vec.size; i++ {
		if vec.array[i] == val {
			return i
		}
	}

	return -1
}

func (vec *vector) Contains(val any) bool {
	return vec.IndexOf(val) >= 0
}

func (vec *vector) Sort(comparator func(any, any) int) {
	algorithms.MergeSort(vec.array, comparator)
}

func (vec *vector) ForEach(processor func(int, any)) {
	for i := 0; i < vec.size; i++ {
		item := vec.array[i]
		processor(item)
	}
}

// Private methods

func (vec *vector) expand() {
	vec.capacity *= 2
	newArray := make([]any, vec.capacity)

	for i := 0; i < vec.size; i++ {
		newArray[i] = vec.array[i]
	}

	vec.array = newArray
}

func (vec *vector) shiftLeft(from, to int) {
	for j := i + 1; j < vec.size; j++ {
		vec.array[j-1] = vec.array[j]
	}
}

func (vec *vector) shiftRight(from, to int) {
	for j := vec.size; j >= i && j > 0; j-- {
		vec.array[j] = vec.array[j-1]
	}
}

