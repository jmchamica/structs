package structs

type linkedListMap struct {
	size int
	head *node
}

type node struct {
	key  any
	val  any
	next *node
}

func NewLinkedListMap() Dictionary {
	return &linkedListMap{
		0,
		nil,
	}
}

// Dictionary interface methods

// boolean return -> was key replaced or new?
func (list *linkedListMap) Set(key, val any) bool {
	newNode := &node{
		key,
		val,
		nil,
	}

	prev := &list.head
	for iter := list.head; iter != nil; iter = iter.next {
		if iter.key == key {
			// replace pair
			newNode.next = iter.next
			*prev = newNode
			return true
		}

		prev = &iter.next
	}

	// new node
	list.size++

	newNode.next = list.head
	list.head = newNode
	return false // no replacement
}

// return boolean -> was key present?
func (list *linkedListMap) Remove(key any) (any, bool) {
	prev := &list.head
	for iter := list.head; iter != nil; iter = iter.next {
		if iter.key == key {
			// detach node from list
			*prev = iter.next
			return iter.val, true
		}
	}

	return nil, false
}

// return boolean -> is key present?
func (list *linkedListMap) Get(key any) (any, bool) {
	for iter := list.head; iter != nil; iter = iter.next {
		if iter.key == key {
			return iter.val, true
		}
	}

	return nil, false
}

func (list *linkedListMap) Size() int {
	return list.size
}
