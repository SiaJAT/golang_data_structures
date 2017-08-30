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
	get(index int)
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
	if list.container == nil {
		return errors.New("Op on non-init list")
	}

	if list.end + 1 >= list.size {
		out_of_capacity_error := errors.New("Add to out-of-capacity list")
		return out_of_capacity_error
	}
	list.container[list.end] = value
	list.end += 1
	return nil
}

func (list *ArrayList) get(index int) (value int, err error) {
	if list.container == nil {
		return 0,errors.New("Op on non-init list")
	}

	if index >= list.end || index < 0 {
		return 0, errors.New("Attempted to get from nonexistent index")
	}
	return list.container[index], nil
}

func (list *ArrayList) remove(index int) (err error) {
	if list.container == nil {
		return errors.New("Op on non-init list")
	}

	if list.end > 0 {
		if index == list.end {
			list.removeLast()
		} else {
			for i := index; i < list.end - 1; i++ {
				list.container[i] = list.container[i+1]
			}
			list.end -= 1
		}
		return nil
	}
	return errors.New("Remove from empty list")
}

func (list *ArrayList) removeLast() (err error) {
	if list.container == nil {
		return errors.New("Op on non-init list")
	}

	if list.end > 0 {
		list.end -= 1
		return nil
	}
	return errors.New("Remove from empty list")
}

func (list *ArrayList) getIterableArray() (arr []int, err error) {
	if list.container == nil {
		return nil, errors.New("Op on non-init list")
	}

	iterable_array := make([]int, list.end)
	for i := 0; i < list.end; i++ {
		iterable_array[i] = list.container[i]
	}
	return iterable_array, nil
}

func (list *ArrayList) resize() (err error) {
	if list.container == nil {
		return errors.New("Op on non-init list")
	}

	old_arr, _ := list.getIterableArray()
	list.size *= 2
	list.container = make([]int, list.size)
	for i := 0; i < list.end; i++ {
		list.container[i] = old_arr[i]
	}
	return nil
}	



func main() {
	fmt.Println("Hello")
}
