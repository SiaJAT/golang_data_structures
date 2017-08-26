package main

import (
	"testing"
)

func Test_Add(t *testing.T) {
	list := LinkedList{}
	list.add(8)
	if list.head.value != 8 {
		t.Fatalf("Add head failed")
	}
}

func Test_RemoveTail(t *testing.T) {
	list := LinkedList{}
	list.add(8)
	list.add(7)
	list.removeTail()
	if list.head.value != 8 || list.tail.value != 8 {
		t.Fatalf("Add head failed")
	}
}

func Test_RemoveHead(t *testing.T) {
	list := LinkedList{}
	list.add(8)
	list.add(7)
	list.removeHead()
	if list.head.value != 7 || list.tail.value != 7 {
		t.Fatalf("Add head failed")
	}
}

func Test_AddAddRemoveHeadAdd(t *testing.T) {
	list := LinkedList{}
	list.add(8)
	list.add(7)
	list.removeHead()
	list.add(9)
	if list.head.value != 7 || list.tail.value != 9 {
		t.Fatalf("AddAddRemoveHead failed")
	}
}

func Test_AddAddRemoveTailAdd(t *testing.T) {
	list := LinkedList{}
	list.add(8)
	list.add(7)
	list.removeTail()
	list.add(9)
	if list.head.value != 8 || list.tail.value != 9 {
		t.Fatalf("AddAddRemoveHead failed")
	}
}

func Test_RemoveHeadEmpty(t *testing.T) {
	list := LinkedList{}
	return_node := list.removeHead()
	if return_node != nil {
		t.Fatalf("RemoveHeadEmpty failed")
	}
}

func Test_RemoveTailEmpty(t *testing.T) {
	list := LinkedList{}
	return_node := list.removeTail()
	if return_node != nil {
		t.Fatalf("RemoveHeadEmpty failed")
	}
}

func Test_RandomAddRemove_1(t *testing.T) {
	list := LinkedList{}
	list.add(1)
	list.removeTail()
	list.add(2)
	list.add(3)
	list.removeHead()
	list.add(5)
	if list.head.value != 3 || list.tail.value != 5 {
		t.Fatalf("RandomAddRemove_1 failed")
	}
}

func Test_RandomAddRemove_2(t *testing.T) {
	list := LinkedList{}
	list.add(1)
	list.add(2)
	list.removeTail()
	list.add(3)
	list.removeTail()
	list.add(4)
	list.add(5)
	list.removeHead()
	list.removeHead()
	if list.head.value != 5 || list.tail.value != 5{
		t.Fatalf("RandomAddRemove_2 failed")
	}
}

func Test_RandomAddRemove_3(t *testing.T) {
	list := LinkedList{}
	list.add(1)
	list.add(2)
	list.removeHead()
	list.add(3)
	list.add(4)
	list.removeTail()
	list.add(5)
	list.removeTail()
	if list.head.value != 2 || list.tail.value != 3 {
		t.Fatalf("RandomAddRemove_3 failed")
	}
}