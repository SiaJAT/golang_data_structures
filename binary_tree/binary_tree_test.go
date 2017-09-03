package main

import (
	"testing"
)

func Test_contains(t *testing.T) {
	val_arr := genIntArray(15)
	tree := str2tree(val_arr)

	for i := 0; i < 15; i++ {
		if tree.contains(i) == false {
			t.Fatalf("Contains not finding value %v\n", i)
		}
	}

	if tree.contains(100) == true {
		t.Fatalf("Tree contains something outside of range")
	}
}

func Test_depth(t *testing.T) {
	val_arr := genIntArray(15)
	tree := str2tree(val_arr)
	if tree.depth() != 4 {
		t.Fatalf("Depth incorrect")
	}
}

func Test_InOrderTraveral(t *testing.T) {
	arr := make([]int, 15)
	arr[0] = 7
	arr[1] = 3
	arr[2] = 8
	arr[3] = 1
	arr[4] = 9
	arr[5] = 4
	arr[6] = 10
	arr[7] = 0
	arr[8] = 11
	arr[9] = 5
	arr[10] = 12
	arr[11] = 2
	arr[12] = 13
	arr[13] = 6 
	arr[14] = 14

	val_arr := genIntArray(15)
	tree := str2tree(val_arr)
	in_trav := tree.inOrderTraversal()
	for index, val := range in_trav {
		if arr[index] != val {
			t.Fatalf("In order traversal broken")
		}
	}
}

func Test_PreOrderTraveral(t *testing.T) {
	arr := make([]int, 15)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 3
	arr[3] = 7
	arr[4] = 8
	arr[5] = 4
	arr[6] = 9
	arr[7] = 10
	arr[8] = 2
	arr[9] = 5
	arr[10] = 11
	arr[11] = 12
	arr[12] = 6
	arr[13] = 13 
	arr[14] = 14

	val_arr := genIntArray(15)
	tree := str2tree(val_arr)
	pre_trav := tree.preOrderTraversal()
	for index, val := range pre_trav {
		if arr[index] != val {
			t.Fatalf("Pre order traversal broken")
		}
	}
}

func Test_PostOrderTraveral(t *testing.T) {
	arr := make([]int, 15)
	arr[0] = 7
	arr[1] = 8
	arr[2] = 3
	arr[3] = 9
	arr[4] = 10
	arr[5] = 4
	arr[6] = 1
	arr[7] = 11
	arr[8] = 12
	arr[9] = 5
	arr[10] = 13
	arr[11] = 14
	arr[12] = 6
	arr[13] = 2 
	arr[14] = 0

	val_arr := genIntArray(15)
	tree := str2tree(val_arr)
	post_trav := tree.postOrderTraversal()
	for index, val := range post_trav {
		if arr[index] != val {
			t.Fatalf("Post order traversal broken")
		}
	}
}

