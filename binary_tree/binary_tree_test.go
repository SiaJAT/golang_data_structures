package main

import (
	"testing"
)

func Test_contains(t *testing.T) {
	val_arr := genIntArray(15)
	tree := pureIntArr2tree(val_arr)

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
	tree := pureIntArr2tree(val_arr)
	if tree.depth() != 4 {
		t.Fatalf("Depth incorrect")
	}
}

func Test_InOrderTraveral(t *testing.T) {
	arr := []int{7,3,8,1,9,4,10,0,11,5,12,2,13,6,14}

	val_arr := genIntArray(15)
	tree := pureIntArr2tree(val_arr)
	in_trav := tree.inOrderTraversal()
	for index, val := range in_trav {
		if arr[index] != val {
			t.Fatalf("In order traversal broken")
		}
	}
}

func Test_PreOrderTraveral(t *testing.T) {
	arr := []int{0,1,3,7,8,4,9,10,2,5,11,12,6,13,14}

	val_arr := genIntArray(15)
	tree := pureIntArr2tree(val_arr)
	pre_trav := tree.preOrderTraversal()
	for index, val := range pre_trav {
		if arr[index] != val {
			t.Fatalf("Pre order traversal broken")
		}
	}
}

func Test_PostOrderTraveral(t *testing.T) {
	arr := []int{7,8,3,9,10,4,1,11,12,5,13,14,6,2,0}

	val_arr := genIntArray(15)
	tree := pureIntArr2tree(val_arr)
	post_trav := tree.postOrderTraversal()
	for index, val := range post_trav {
		if arr[index] != val {
			t.Fatalf("Post order traversal broken")
		}
	}
}

func Test_Serialize(t *testing.T) {
	impure_test := []interface{}{0,1,2,3,4,"nil",6,7,"nil","nil",10,"nil","nil",13,"nil"}
	impure_tree, _ := impureArr2Tree(impure_test)
	if impure_tree.root.left.right.right.value != 10 {
		t.Fatalf("Test_Serialize: Left, right, right, right failed")
	}

	if impure_tree.root.right.left != nil {
		t.Fatalf("Test_Serialize: Right, left failed")
	}

	if impure_tree.root.right.right.right != nil {
		t.Fatalf("Test_Serialize: Right, right, right failed")
	}
}

func Test_SerializeInvalid(t *testing.T) {
	bad_arr_1 := []interface{}{"nil", 1, 2}
	ok_arr := []interface{}{0, "nil", 2, "nil", "nil"}
	bad_arr_2 := []interface{}{0, 1, "nil", "nil", "nil", 5, 6}
	bad_arr_3 := []interface{}{0,1,2,3}

	if  _, err_1 := impureArr2Tree(bad_arr_1); err_1 == nil {
		t.Fatalf("Serialize Check failed for example 1 (nil pointer)")
	}

	if _, err_2 := impureArr2Tree(bad_arr_2); err_2 == nil {
		t.Fatalf("Serialize Check failed for example 2 (nil pointer)")
	}

	if _, err_3 := impureArr2Tree(ok_arr); err_3 != nil {
		t.Fatalf("Serialize shouldn't fail for valid array")
	}

	if _, err_4 := impureArr2Tree(bad_arr_3); err_4 == nil {
		t.Fatalf("Serialize Check failed for example 3 (incomplete binary tree)")
	}
}

func Test_SerializeDeSerialize(t *testing.T) {
	impure_test := []interface{}{0,1,2,3,4,"nil",6,7,"nil","nil",10,"nil","nil",13,"nil"}
	impure_tree, serialize_err := impureArr2Tree(impure_test)
	if serialize_err == nil {
		back_to_arr, deserialize_err := tree2ImpureArr(impure_tree)
		if deserialize_err == nil {
			for index, _ := range back_to_arr {
				if impure_test[index] != back_to_arr[index] {
					t.Fatalf("Deserialize fails to be inverse operation")
				}
			}
		} else {
			t.Fatalf("Deserialization failed")
		}
	} else {
		t.Fatalf("Serialization failed.")
	}
}

