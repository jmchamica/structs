package test

import (
	"../structs"
	"testing"
)

func TestSinglyLinkedListFunctions(t *testing.T) {
	var list structs.List
	list = structs.NewSinglyLinkedList()

	n := 10000

	for i := 0; i < n; i++ {
		list.InsertHead(i)
	}

	for i := 0; i < list.Size(); i++ {
		val, err := arrayList.Get(i)
		if err != nil {
			t.Fatal(err)
		}

		if val != i {
			t.Fatal("Value mismatch")
		}
	}

	for i := 0; i < n; i++ {
		err := arrayList.Remove(0)
		if err != nil {
			t.Fatal(err)
		}

		for j := 0; j < arrayList.Size(); j++ {
			val, err := arrayList.Get(j)
			if err != nil {
				t.Fatal(err)
			}

			if val != j+i+1 {
				t.Fatal("Value mismatch")
			}
		}
	}
}
