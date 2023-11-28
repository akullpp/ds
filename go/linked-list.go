package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (ll *LinkedList[T]) String() string {
	var str string

	for node := ll.head; node != nil; node = node.next {
		str += fmt.Sprintf("%v", node.data)
		if node.next != nil {
			str += " -> "
		}
	}
	return str
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

func main() {
	ll := LinkedList[int]{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Prepend(0)
	fmt.Println(ll.String())
}
