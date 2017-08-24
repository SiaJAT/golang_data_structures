package main

import (
	"fmt"
	"errors"
)

type Node struct {	
	value int
	next *Node
}

type Stack struct {
	top *Node
}

type StackOps interface {
	push(val int)
	pop() *Node
}

func (stack *Stack) push(val int) {
	n := Node{val, nil}
	if stack.top == nil {	
		stack.top = &n
	} else {
		n.next = stack.top
		stack.top = &n
	}
}

func (stack *Stack) pop() *Node {
	if stack.top == nil {
		empty_stack_error := errors.New("Pop from empty stack!")
		fmt.Println(empty_stack_error)
	} 
	return_node := stack.top
	stack.top = stack.top.next
	return return_node
}


func main() {

}