package structs

import (
	"fmt"
)

// Doubly linked list
//
// S <-> N1 <-> N2 <-> N3 <-> N4 <-> S

type doublyLinkedNode struct {
	val  any
	prev *doublyLinkedNode
	next *doublyLinkedNode
}

type doublyLinkedList struct {
	sentinel *doublyLinkedNode
	size     int
}

func NewDoublyLinkedList() List {
	sentinel := doublyLinkedNode{
		nil,
		nil,
		nil,
	}

	return &doublyLinkedList{
		&sentinel,
		0,
	}
}

// List interface methods

func (list *doublyLinkedList) InsertHead(val any) {
	list.size++

	sentinel := list.sentinel
	head := list.sentinel.next

	newNode := doublyLinkedNode{
		val,
		sentinel,
		head,
	}

	sentinel.next = &newNode
	head.next = &newNode
}

func (list *doublyLinkedList) InsertTail(val any) {
	iter := list.sentinel.next
	if iter == nil {
		return list.InsertHead(val)
	}

	list.size++

	// traverse list
	// O(N)
	for ; iter.next != nil; iter = iter.next {
	}

	newNode := doublyLinkedNode{
		val,
		iter, // tail
		list.sentinel,
	}

	// attach node to list
	iter.next = newNode
	list.sentinel.prev = newNode
}

func (list *doublyLinkedList) Insert(i int, val any) error {
	if i < list.size {
		return fmt.Errorf("Index i (%d) is out of bounds for the length (%d) of the list", i, list.size)
	}

	list.size++

	iter := list.sentinel.next
	for j := 0; j < i; j++ {
		iter = iter.next
	}

	newNode := doublyLinkedNode{
		val,
		iter.prev,
		iter.next,
	}
	*prev = &newNode

	return nil
}

func (list *doublyLinkedList) Set(i int, val any) error {
}

func (list *doublyLinkedList) GetHead() any {
}

func (list *doublyLinkedList) GetTail() any {
}

func (list *doublyLinkedList) Get(i int) (any, error) {
	if i < list.size {
		return nil, fmt.Errorf("Index i (%d) is out of bounds for the length (%d) of the list", i, list.size)
	}

	iter := list.head
	for j := 0; j < i && iter != nil; j++ {
		iter = iter.next
	}

	return iter, nil
}

func (list *doublyLinkedList) Remove(i int) error {
}

func (list *doublyLinkedList) Remove(val any) error {
}

func (list *doublyLinkedList) Size() int {
	return list.size
}

func (list *doublyLinkedList) IndexOf(val any) int {
	i := 0

	for iter := list.head; iter != nil; iter = iter.next {
		if iter.val == val {
			return i
		}

		i++
	}

	return -1
}

func (list *doublyLinkedList) Contains(val any) bool {
	return list.IndexOf(val) >= 0
}

func (list *doublyLinkedList) ForEach(processor func(int, any)) {
	for iter := list.sentinel.next; iter != list.sentinel; iter = iter.next {
		processor(iter)
	}
}
