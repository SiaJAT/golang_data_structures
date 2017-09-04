package main

import (

)

type Heap struct {
	container []int
	size int
	min int
}

type HeapOps interface {
	perc_up(val int)
	perc_down(val int)
	insert(val int)
	min_child(val int)
	delete_min()
	build_heap()
}

func (heap *Heap) perc_up(val int) {

}

func (heap *Heap) perc_down(val int) {

}

func (heap *Heap) insert(val int) {

}

func (heap *Heap) min_child(val int) {

}

func (heap *Heap) delete_min() {

}

func (heap *Heap) build_heap() {
	
}

func main() {

}