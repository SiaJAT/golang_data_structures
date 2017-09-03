package main

import (
	"strconv"
	"fmt"
)

type Node struct {
	value int
	left *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

type TreeOps interface {
	depth() int
	inOrderTraversal() []*Node
	preOrderTraversal() []*Node
	postOrderTraversal() []*Node
}

func genIntArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	return arr
}

func str2tree(arr []int) *BinaryTree {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	node_arr := make([]*Node, len(arr))

	// create unlinked nodes with value equal to their index in arr
	for index, _ := range node_arr {
		node_arr[index] = &Node{arr[index], nil, nil}
	}

	// 2*index + 1 == left
	// 2*index + 2 == right
	for index,_ := range node_arr {
		if 2*index + 1 < len(arr) {
			node_arr[index].left = node_arr[2*index+1]
		}
		
		if 2*index + 2 < len(arr) {
			node_arr[index].right = node_arr[2*index+2]
		}		
	}

	tree := BinaryTree{node_arr[0]}

	return &tree
}

func _byte2int(b byte) int {
	val, _ := strconv.Atoi(string(b))
	return val
}

func (tree *BinaryTree) depth() int {
	if tree.root == nil {
		return 0
	}
	return _depth(tree.root)
}

func _depth(node *Node) int {
	if node == nil {
		return 0
	}

	if node.left == nil && node.right == nil {
		return 1
	}

	left_depth := _depth(node.left)
	right_depth := _depth(node.right)

	if left_depth >= right_depth {
		return left_depth + 1
	} else {
		return right_depth + 1
	}
}

func (tree *BinaryTree) inOrderTraversal() []*Node {
	
}

func _inOrderTraversal(node *Node, trav []*Node) {

}


func (tree *BinaryTree) preOrderTraversal() []*Node  {
	return nil
}

func _preOrderTraversal(node *Node, trav []*Node) {
	
}

func (tree *BinaryTree) postOrderTraversal() []*Node {
	return nil
}

func _postOrderTraversal(node *Node, trav []*Node) {
	
}

func main() {
	arr := genIntArray(16)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\n", arr[i])
	}

	//            0  
	//     1              2
	//  3     4       5       6
	// 7 8   9 10   11 12   13 14
	fmt.Println("testing")
	tree := str2tree(arr)
	fmt.Printf("%v\n", tree.root.value)
	fmt.Printf("%v\n", tree.root.left.value)
	fmt.Printf("%v\n", tree.root.left.right.value)
	fmt.Printf("%v\n", tree.root.left.right.right.value)

	DNE := tree.root.left.right.right.left == nil
	fmt.Printf("%v\n", DNE)

	tree_depth := tree.depth()
	fmt.Printf("%v\n", tree_depth)

}

