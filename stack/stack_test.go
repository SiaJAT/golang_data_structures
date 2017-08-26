package main

import (
	"testing"
)

func Test_Push(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	if stack.top.value != 1 {
		t.Fatalf("Push failed")
	}
}

func Test_Pop(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.pop() 
	if stack.top.value != 1 {
		t.Fatalf("Pop failed")
	}
}

func Test_PopEmpty(t *testing.T) {
	stack := Stack{}
	_, err := stack.pop()
	if err == nil {
		t.Fatalf("Bad pop")
	}
}

func Test_PushThenPop(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	stack.pop()
	stack.push(2)
	if stack.top.value != 2 {
		t.Fatalf("Failed push then pop")
	}
}
