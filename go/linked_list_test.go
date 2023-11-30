package linkedlist

import "testing"

func TestLinkedListPrepend(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Prepend(1)
	ll.Prepend(2)
	ll.Prepend(3)

	t.Log(ll.String())

	if ll.String() != "3 -> 2 -> 1" {
		t.Errorf("Expected 3 -> 2 -> 1, got %v", ll.String())
	}
}

func TestLinkedListAppend(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	t.Log(ll.String())

	if ll.String() != "1 -> 2 -> 3" {
		t.Errorf("Expected 1 -> 2 -> 3, got %v", ll.String())
	}
}
