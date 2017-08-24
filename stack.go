package main

import (
	"fmt"
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
	
}

func (stack *Stack) pop() *Node {
	
}


func main() {

}