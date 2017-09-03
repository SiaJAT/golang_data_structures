package main

import (
	"strconv"
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

func str2tree (str string) *BinaryTree {
	if str == "" {
		return nil
	}

	root_val := _byte2int(str[0])

	root := Node{root_val, nil, nil}
	tree := BinaryTree{&root}

	node_arr := make([]*Node, len(str))

	// a b c d e f g
	// 0 1 2 3 4 5 6
	// 
	//     0
	//   1   2
	//  3 4 5 6

	return nil 

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
	return nil
}

func (tree *BinaryTree) preOrderTraversal() []*Node  {
	return nil
}

func (tree *BinaryTree) postOrderTraversal() []*Node {
	return nil
}

func main() {
	tree := BinaryTree{}

}

