package main

import (
	"testing"
)

func Test_Add(t *testing.T) {
	list := ArrayList{}
	list.init()
	list.add(2)
	if val,_ := list.get(0); val != 2 {
		t.Fatalf("Basic add failed")
	}
}


func Test_Remove(t *testing.T) {
	list := ArrayList{}
	list.init()
	list.add(1)
	list.add(2)
	list.add(3)
	list.remove(1)
	if val,_ := list.get(1); val != 3 {
		t.Fatalf("Faulty remove")
	}

}

func Test_RemoveLast(t *testing.T) {
	list := ArrayList{}
	list.init()
	list.add(1)
	list.add(2)
	list.add(3)
	list.removeLast()
	if val,_ := list.get(1); val != 2 {
		t.Fatalf("Faulty removeLast")
	}
}

func Test_RemoveEmpty(t *testing.T) {
	list := ArrayList{}
	list.init()
	if err := list.remove(0); err == nil {
		t.Fatalf("Empty remove error not caught")
	}
}

func Test_RemoveLastEmpty(t *testing.T) {
	list := ArrayList{}
	list.init()
	if err := list.removeLast(); err == nil {
		t.Fatalf("Empty removeLast error not caught")
	}
}



func Test_Iterable(t *testing.T) {
	list := ArrayList{}
	list.init()
	list.add(1)
	list.add(2)
	list.add(3)

	iterable, _ := list.getIterableArray()
	for i := 0; i < len(iterable); i++ {
		if iterable[i] != i+1 {
			t.Fatalf("Bad values in iterable")
		}
	}
	
	iterable[2] = 4

	if val, _ := list.get(2); val != 3 {
		t.Fatalf("Iterable array should return array copy")
	}
	
}

func Test_Resize(t *testing.T) {
	list := ArrayList{}
	list.init()
	for i := 0; i < 9; i++ {
		list.add(i)
	}
	list.resize()
	if err := list.add(10); err != nil && list.size <= 10 {
		t.Fatalf("Resize failed")
	}
}

func Test_RandomAddRemove_1(t *testing.T) {
	list := ArrayList{}
	list.init()
	list.add(1)
	list.add(2)
	list.removeLast()
	list.add(3)
	list.add(4)
	list.add(5)
	if val,_ := list.get(1); val != 3 {
		t.Fatalf("RandomAddRemove_1 failed")
	}
}

func Test_RandomAddRemove_2(t *testing.T) {
	list := ArrayList{}
	list.init()
	list.add(1)
	list.add(2)
	list.remove(0)
	list.add(3)
	list.add(4)
	list.add(5)
	list.removeLast()
	if valSecond,_ := list.get(2); valSecond != 4  {
		t.Fatalf("RandomAddRemove_2 failed")
	}
}

func Test_RandomAddRemove_3(t *testing.T) {
	list := ArrayList{}
	list.init()
	list.add(1)
	list.removeLast()
	list.add(2)
	list.add(3)
	list.add(4)
	list.remove(1)
	list.add(5)
	if valSecond,_ := list.get(1); valSecond != 4  {
		t.Fatalf("RandomAddRemove_3 failed")
	}
}


func Test_AddUninit(t *testing.T) {
	list := ArrayList{}
	if err := list.add(8); err == nil {
		t.Fatalf("Add on un-initialized list should fail")
	}
}

func Test_GetUninit(t *testing.T) {
	list := ArrayList{}
	if _, err := list.get(0); err == nil {
		t.Fatalf("Get on un-initialized list should fail")
	}
}

func Test_RemoveEmptyUninit(t *testing.T) {
	list := ArrayList{}
	if err := list.remove(0); err == nil {
		t.Fatalf("Remove on un-initialized list should fail")
	}
}

func Test_RemoveLastEmptyUninit(t *testing.T) {
	list := ArrayList{}
	if err := list.removeLast(); err == nil {
		t.Fatalf("RemoveLast on un-initialized list should fail")
	}
}

func Test_IterableUninit(t *testing.T) {
	list := ArrayList{}
	if _, err := list.getIterableArray(); err == nil {
		t.Fatalf("GetIterableArray on un-initialized list should fail")
	}
}

func Test_ResizeUninit(t *testing.T) {
	list := ArrayList{}
	if err := list.resize; err == nil {
		t.Fatalf("Cannot resize un-initialized list")
	}
}
