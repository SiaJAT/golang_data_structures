package main

import (
	"fmt"
	"errors"
)

const CONTAINER_CAPACITY = 10

type ArrayList struct {
	container []int
	size int
	end int
}

type List interface {
	add(value int)
	remove(index int)
	removeLast()
	getIterableArray() []int
}

func (list *ArrayList) init() {
	list.container = make([]int, 10)
	list.size = 10
	list.end = 0
}

func (list *ArrayList) add(value int) (err error) {
	if list.end + 1 >= list.size {
		out_of_capacity_error := errors.New("Add to out-of-capacity list")
		return out_of_capacity_error
	}
	list.container[list.end] = value
	list.end += 1
	return nil
}

func (list *ArrayList) remove(index int) {
	
}

func (list *ArrayList) removeLast() {
	if end > 0 {
		end -= 1
	}
}

func (list *ArrayList) getIterableArray() []int {
	return list.container[0:end]
}


func main() {
	fmt.Println("Hello")
}
