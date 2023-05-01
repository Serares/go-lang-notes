package generics

import (
	"fmt"
)

// Node represents a singly-linked list that holds
// values of any type.
type Node[T any] struct {
	next *Node[T]
	val  T
}

func (l *Node[int]) InsertNumber(v int) {
	new_node := Node[int]{val: v}
	l.next = &new_node
}

func UseList() {
	list := Node[int]{val: 1}
	list.InsertNumber(2)
	number2 := list.next
	fmt.Printf("The list %d, ", number2.val)
	number2.InsertNumber(3)
}
