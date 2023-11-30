package linkedlist

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	// Append becomes O(1), otherwise O(n)
	tail *Node[T]
}

func (ll *LinkedList[T]) Prepend(data T) {
	node := Node[T]{data: data, next: ll.head}
	ll.head = &node
	if ll.tail == nil {
		ll.tail = &node
	}
}

func (ll *LinkedList[T]) Append(data T) {
	node := Node[T]{data: data}

	if ll.head == nil {
		ll.head = &node
		ll.tail = &node
	} else {
		ll.tail.next = &node
		ll.tail = &node
	}
}

func (ll *LinkedList[T]) Reverse() {
	var prev *Node[T]
	current := ll.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.tail = ll.head
	ll.head = prev
}

func (ll *LinkedList[T]) String() string {
	var str string

	current := ll.head
	for current != nil {
		str += fmt.Sprintf("%v", current.data)
		if current.next != nil {
			str += " -> "
		}
		current = current.next
	}
	return str
}
