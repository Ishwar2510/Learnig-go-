package main

import (
	"errors"
	"fmt"
)

// construct an linked list
// it can be of any type
// it should be a doubley linked list

type Node struct {
	next *Node
	data interface{}
	prev *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
}

func (list *Linkedlist) Insert(data interface{}) error {
	newNode := &Node{data: data}
	if list.isEmpty() != nil {
		list.head = newNode
		list.tail = list.head
		return nil
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	list.tail = current.next
	return nil
}
func (list *Linkedlist) isEmpty() error {
	if list.head == nil {
		return errors.New("Empty list")
	}
	return nil
}

func (list *Linkedlist) PeekLast() (interface{}, error) {
	if list.head == nil {
		return nil, errors.New("Empty list")
	}
	return list.head.data, nil
	// or use tail
	// return list.tail.data, nil
}

func (list *Linkedlist) PeekFirst() (interface{}, error) {
	if list.head == nil {
		return nil, errors.New("Empty list ")
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, nil

}
func (list *Linkedlist) RemoveFirst() (interface{}, error) {
	if err := list.isEmpty(); err != nil {
		return nil, err
	}
	current := list.head
	if current.next == nil {
		list.head = nil
		return nil, nil
	}
	toRemove := list.head
	nextNode := current.next
	list.head = nextNode
	list.head.prev = nil
	return toRemove.data, nil

}

func (list *Linkedlist) RemoveLast() (interface{}, error) {
	if err := list.isEmpty(); err != nil {
		return nil, err
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	prev := current.prev
	prev.next = nil
	return current.data, nil

	// or

	// list.tail = list.tail.prev
	// list.tail.next = nil
	// return nil
}
func (list *Linkedlist) GetMiddle() (interface{}, error) {
	// edge cases
	// single element
	// even or odd

	if err := list.isEmpty(); err != nil {
		return nil, err
	}
	if list.head.next == nil {
		return list.head, nil
	}
	firstPointer := list.head
	secondPointer := firstPointer.next

	for secondPointer != nil && secondPointer.next != nil {
		secondPointer = secondPointer.next.next
		firstPointer = firstPointer.next
	}
	return firstPointer.data, nil

	// Since it is doubly linked we can use tail

	// start := list.head
	// end := list.tail

	// for start.next != end || start != end {
	// 	start = start.next
	// 	end = end.prev
	// }
	// return start, nil
}
func (list *Linkedlist) PrintList() {
	if err := list.isEmpty(); err != nil {
		fmt.Println("Empty List")
		return
	}
	current := list.head
	for current != nil {
		if current.next == nil {
			fmt.Println(current.data)
		} else {
			fmt.Print(current.data, "--->")
		}
		current = current.next
	}

}
func (list *Linkedlist) ReverseList() error {
	if err := list.isEmpty(); err != nil {
		return err
	}

	var prev *Node
	current := list.head
	cnext := current.next
	for cnext != nil {
		current.next = prev
		current.prev = cnext
		cnext.prev = nil
		prev = current
		current = cnext
		cnext = cnext.next
	}
	current.next = prev
	list.head = current
	return nil
}

func NewLinkedList() *Linkedlist {
	return &Linkedlist{}
}
