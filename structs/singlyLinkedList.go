package structs

import (
	"../algorithms"
	"fmt"
)

// Singly linked list
//
// N1 -> N2 -> N3 -> N4

type singlyLinkedNode struct {
	val  any
	next *singlyLinkedNode
}

type singlyLinkedList struct {
	head *singlyLinkedNode
	size int
}

func NewSinglyLinkedList() List {
	return &singlyLinkedList{
		nil,
		0,
	}
}

// List interface methods

func (list *singlyLinkedList) InsertHead(val any) {
	list.size++

	newNode := singlyLinkedNode{
		val,
		list.head,
	}

	list.head = &newNode
}

func (list *singlyLinkedList) InsertTail(val any) {
	iter := list.head
	if iter == nil {
		return list.InsertHead(val)
	}

	list.size++

	// traverse list
	// O(N)
	for ; iter.next != nil; iter = iter.next {
	}

	// append to tail
	iter.next = &singlyLinkedNode{
		val,
		nil,
	}
}

func (list *singlyLinkedList) Insert(i int, val any) error {
	if i >= list.size {
		return fmt.Errorf("Index i (%d) is out of bounds for the length (%d) of the list", i, list.size)
	}

	list.size++

	newNode := singlyLinkedNode{
		val,
		nil,
	}

	prev := &list.head
	iter := list.head
	for j := 0; j < i; j++ {
		prev = &iter.next
		iter = iter.next
	}

	newNode.next = iter
	*prev = &newNode

	return nil
}

func (list *singlyLinkedList) Set(i int, val any) error {
	if i >= list.size {
		return fmt.Errorf("Index i (%d) is out of bounds for the length (%d) of the list", i, list.size)
	}

	newNode := singlyLinkedNode{
		val,
		nil,
	}

	prev := &list.head
	iter := list.head
	for j := 0; j < i; j++ {
		prev = &iter.next
		iter = iter.next
	}

	newNode.next = iter.next // skipping iter from list
	*prev = &newNode

	return nil
}

func (list *singlyLinkedList) GetHead() any {
	return list.head
}

func (list *singlyLinkedList) GetTail() any {
	iter := list.head

	// traverse until tail
	for ; iter.next != nil; iter = iter.next {
	}

	return iter
}

func (list *singlyLinkedList) Get(i int) (any, error) {
	if i >= list.size {
		return nil, fmt.Errorf("Index i (%d) is out of bounds for the length (%d) of the list", i, list.size)
	}

	iter := list.head
	for j := 0; j < i; j++ {
		iter = iter.next
	}

	return iter, nil
}

func (list *singlyLinkedList) RemoveHead() (any, bool) {
	val, _, removed := list.Remove(0)
	return val, removed
}

func (list *singlyLinkedList) RemoveTail() (any, bool) {
	lastIndex := list.size - 1
	val, _, removed := list.Remove(lastIndex)
	return val, removed
}

func (list *singlyLinkedList) Remove(i int) (any, error, bool) {
	if i >= list.size {
		return nil, fmt.Errorf("Index i (%d) is out of bounds for the length (%d) of the list", i, list.size), false
	}

	// find the ith node
	prev := &list.head
	iter := list.head
	for j := 0; j < i; j++ {
		prev = &iter.next
		iter = iter.next
	}

	// skipping iter from list
	*prev = iter.next

	return iter.val, nil, true
}

func (list *singlyLinkedList) Remove(val any) bool {
	prev := &list.head
	iter := list.head

	for i := 0; i < list.size; j++ {
		if iter.val == val {
			// skipping iter from list
			*prev = iter.next
			return true
		}

		prev = &iter.next
		iter = iter.next
	}

	return false
}

func (list *singlyLinkedList) Push(val any) {
	list.InsertHead(val)
}

func (list *singlyLinkedList) Pop() (any, bool) {
	return list.RemoveHead()
}

func (list *singlyLinkedList) Size() int {
	return list.size
}

func (list *singlyLinkedList) IndexOf(val any) int {
	i := 0

	for iter := list.head; iter != nil; iter = iter.next {
		if iter.val == val {
			return i
		}

		i++
	}

	return -1
}

func (list *singlyLinkedList) Contains(val any) bool {
	return list.IndexOf(val) >= 0
}

func (list *singlyLinkedList) Sort(comparator func(any, any) int) {
	array := list.toArray()

	algorithms.MergeSort(array, comparator)

	// copy sorted array onto linked list
	i := 0
	for iter := list.head; iter != nil; iter = iter.next {
		iter.val = array[i]
		i++
	}

}

func (list *singlyLinkedList) ForEach(processor func(int, any)) {
	for iter := list.head; iter != nil; iter = iter.next {
		processor(iter)
	}
}

// Private methods

func (list *singlyLinkedList) toArray() []any {
	array := make([]any, list.Size())

	i := 0
	for iter := list.head; iter != nil; iter = iter.next {
		array[i] = iter.val
		i++
	}

	return array
}
