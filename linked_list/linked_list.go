package main

import (
	"fmt"
)

type Node struct {	
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

type List interface {
	add(val int)
	getHead() *Node
	getTail() *Node
	removeHead() *Node
	removeTail() *Node
}

// add a node to the back of the list
func (list *LinkedList) add(val int) {
	if list.head == nil {
		n := Node{val, nil}
		list.head = &n 
		list.tail = &n
	} else if list.head == list.tail {
		n := Node{val, nil}
		list.tail = &n
		list.head.next = list.tail
	} else {
		n := Node{val, nil}
		list.tail.next = &n
		list.tail = &n
	}
}

func (list *LinkedList) removeHead() *Node {
	old_head := list.head
	if list.head != nil && list.head == list.tail{
		list.head = nil
		list.tail = nil
	} else if list.head != nil {
		list.head = list.head.next
	}
	return old_head
}

func (list *LinkedList) removeTail() *Node {
	old_tail  := list.tail
	if list.head != nil && list.head == list.tail {
		list.head = nil
		list.tail = nil
		return old_tail
	} else if list.head != nil {
		prev := &Node{0, nil}
		curr := &Node{0, nil}
		for i := list.head; i != nil; i = i.next {
			prev = curr
			curr = i
		}
		prev.next = nil
		list.tail = prev
		return curr
	}
	return old_tail
}

func main() {
	fmt.Println("test")
	list := LinkedList{}
	list.add(2)
	list.add(3)
	list.add(4)
	fmt.Println("head: ", list.head.value, ", tail: ", list.tail.value)
	list.removeTail()
	list.add(5)
	fmt.Println("head: ", list.head.value, ", tail: ", list.tail.value)
	list.removeHead()
	// fmt.Println(head.value)
	// fmt.Println(head.next.value)
	for i := list.head; i != nil; i = i.next {
		fmt.Println(i.value)
	}
}
