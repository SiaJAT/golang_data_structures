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

func _powVal(base int, pow int) int {
	accum := 1
	for i := 0; i < pow; i++ {
		accum *= base
	}
	return accum
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

func (tree *BinaryTree) inOrderTraversal() []int {
	if tree == nil {
		return nil
	}
	// max # nodes in binary tree of "d" depth = 2^d - 1
	arr := make([]int, 0, _powVal(2,tree.depth())-1)
	
	_inOrderTraversal(tree.root, &arr)
	return arr
}

func _inOrderTraversal(node *Node, trav *[]int) {
	if node == nil {
		return
	}

	if node.left != nil {
		_inOrderTraversal(node.left, trav)
	}

	if node != nil {
		*trav = append(*trav, node.value)
	}

	if node.right != nil {
		_inOrderTraversal(node.right, trav)
	}
}


func (tree *BinaryTree) preOrderTraversal() []int  {
	if tree == nil {
		return nil
	}
	// max # nodes in binary tree of "d" depth = 2^d - 1
	arr := make([]int, 0, _powVal(2,tree.depth())-1)
	
	_preOrderTraversal(tree.root, &arr)
	return arr
}

func _preOrderTraversal(node *Node, trav *[]int) {
	if node == nil {
		return
	}

	if node != nil {
		*trav = append(*trav, node.value)
	}

	if node.left != nil {
		_preOrderTraversal(node.left, trav)
	}

	if node.right != nil {
		_preOrderTraversal(node.right, trav)
	}
}

func (tree *BinaryTree) postOrderTraversal() []int {
	if tree == nil {
		return nil
	}
	// max # nodes in binary tree of "d" depth = 2^d - 1
	arr := make([]int, 0, _powVal(2,tree.depth())-1)
	
	_postOrderTraversal(tree.root, &arr)
	return arr
}

func _postOrderTraversal(node *Node, trav *[]int) {
	if node == nil {
		return
	}

	if node.left != nil {
		_postOrderTraversal(node.left, trav)
	}

	if node.right != nil {
		_postOrderTraversal(node.right, trav)
	}

	if node != nil {
		*trav = append(*trav, node.value)
	}
}

func main() {
	arr := genIntArray(16)
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

	test_pow := _powVal(2, 5)
	fmt.Printf("%v\n", test_pow)	

	fmt.Println("in trav...")
	in_trav := tree.inOrderTraversal()
	for _, val := range in_trav {
		fmt.Printf("%v\n", val)
	}

	fmt.Println("pre trav...")
	pre_trav := tree.preOrderTraversal()
	for _, val := range pre_trav {
		fmt.Printf("%v\n", val)
	}

	fmt.Println("post trav...")
	post_trav := tree.postOrderTraversal()
	for _, val := range post_trav {
		fmt.Printf("%v\n", val)
	}




}

