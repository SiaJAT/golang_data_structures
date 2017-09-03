package main

import (
	"strconv"
	"fmt"
	"errors"
	"reflect"
	"math"
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
	contains(val int) bool
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

func pureIntArr2tree(arr []int) *BinaryTree {
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

func (tree *BinaryTree) contains(val int) bool {
	if tree == nil {
		return false
	}
	return _contains(tree.root, val)
}

func _contains(node *Node, val int) bool {
	if node == nil {
		return false
	}
	if node.value == val {
		return true
	} else {
		return _contains(node.left, val) || _contains(node.right, val)
	}
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

// invalid tree == there is some nil.left --> node || nil.right --> node
// e.g. [nil, 1, 2]
func validateArrTreeRepresentation(arr []interface{}) bool {
	for index,_ := range arr {
		if arr[index] == "nil" {
			if 2*index + 1 < len(arr) && arr[2*index+1] != "nil" {
				return false
			}

			if 2*index + 2 < len(arr) && arr[2*index+2] != "nil" {
				return false
			}
		}
	}
	return true
}

// TODO
// build a tree from a list containing ints and nils
// e.g., [1,2,nil]
func impureArr2Tree(arr []interface{}) (tree *BinaryTree, err error) {
	if arr == nil || len(arr) == 0 {
		return nil, errors.New("Nil or 0 length representation passed. Aborting.")
	}

	if validateArrTreeRepresentation(arr) == false {
		return nil, errors.New("Bad array representation: Nil pointer")
	}

	if a := math.Mod(float64(len(arr)+1),2); a != 0 {
		return nil, errors.New("Bad array representaiton: Not complete binary tree")
	}
	
	node_arr := make([]*Node, len(arr))

	// create unlinked nodes with value equal to their index in arr
	for index, _ := range node_arr {
		if _, ok := arr[index].(int); ok {
			node_arr[index] = &Node{arr[index].(int), nil, nil}
		} else {
			node_arr[index] = nil
		}
	}

	// 2*index + 1 == left
	// 2*index + 2 == right
	for index,_ := range node_arr {
		if 2*index + 1 < len(arr) && node_arr[index] != nil {
			node_arr[index].left = node_arr[2*index+1]
		}
		
		if 2*index + 2 < len(arr) && node_arr[index] != nil {
			node_arr[index].right = node_arr[2*index+2]
		}		
	}

	tree = &BinaryTree{node_arr[0]}

	return tree, nil
}

func tree2ImpureArr(tree *BinaryTree) (arr []interface{}, err error) {
	if tree == nil {
		return nil, errors.New("Nil tree. Aborting.")
	}

	arr = make([]interface{}, _powVal(2,tree.depth())-1)
	for i, _ := range arr {
		arr[i] = "nil"
	}

	_tree2ImpureArr(tree.root, 0, &arr)

	return arr, nil
}

func _tree2ImpureArr(node *Node, nodeNum int, arr *[]interface{}) {
	if node != nil {
		(*arr)[nodeNum] = node.value
		_tree2ImpureArr(node.left, 2*nodeNum+1, arr)
		_tree2ImpureArr(node.right, 2*nodeNum+2, arr)
	}
}



func main() {
	arr := genIntArray(15)
	//            0  
	//     1              2
	//  3     4       5       6
	// 7 8   9 10   11 12   13 14
	fmt.Println("testing")
	tree := pureIntArr2tree(arr)
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

	fmt.Println("contains test")
	does_contain_0 := tree.contains(0)
	does_contain_11 := tree.contains(21)
	fmt.Printf("%v\n", does_contain_0)
	fmt.Printf("%v\n", does_contain_11)

	mixed := make([]interface{}, 10)
	mixed[0] = "nil"
	mixed[1] = 1
	fmt.Printf("%v\n", mixed[0])
	fmt.Printf("%v\n", mixed[1])
	fmt.Printf("%v\n", reflect.TypeOf(mixed[0]))
	fmt.Printf("%v\n", reflect.TypeOf(mixed[1]))
	_, ok := mixed[1].(int)
	fmt.Printf("%v\n", ok)

	//            0  
	//     1             2
	//  3     4      !      6
	// 7 !   ! 10  !  !   13  !
	impure_test := []interface{}{0,1,2,3,4,"nil",6,7,"nil","nil",10,"nil","nil",13,"nil"}
	impure_tree, err := impureArr2Tree(impure_test)
	if err == nil {
		impure_trav := impure_tree.inOrderTraversal()
		for _, val := range impure_trav {
			fmt.Printf("%v\n", val)
		}

		back_to_arr, _ := tree2ImpureArr(impure_tree)
		fmt.Printf("%#v\n", back_to_arr)
	} else {
		fmt.Println("Err")
	}

	b := math.Mod(4,2)

	fmt.Printf("%v\n", b)
	
}

