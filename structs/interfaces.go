package structs

type any = interface{}

type List interface {
	InsertHead(any)
	InsertTail(any)
	Insert(int, any) error
	Set(int, any) error

	GetHead() any
	GetTail() any
	Get(int) (any, error)

	RemoveHead() (any, bool)
	RemoveTail() (any, bool)
	Remove(int) (any, error, bool)
	Remove(any) bool

	Push(any)
	Pop() (any, bool)

	Size() int
	IndexOf(any) int
	Contains(any) bool

	Sort(func(any, any) int)

	// Functional
	ForEach(func(int, any))
}

type Dictionary interface {
	Set(any, any) bool
	Remove(any) (any, bool)
	Get(any) (any, bool)
	Size() int
}
