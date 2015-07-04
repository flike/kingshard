package router

import (
	"testing"
)

func testCheckList(t *testing.T, l []int, checkList ...int) {
	if len(l) != len(checkList) {
		t.Fatal("invalid list len", len(l), len(checkList))
	}

	for i := 0; i < len(l); i++ {
		if l[i] != checkList[i] {
			t.Fatal("invalid list item", l[i], i)
		}
	}
}

func TestListSet(t *testing.T) {
	var l1 []int
	var l2 []int
	var l3 []int

	l1 = []int{1, 2, 3}
	l2 = []int{2}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 3}
	l2 = []int{2, 3}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2, 3)

	l1 = []int{1, 2, 4}
	l2 = []int{2, 3}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 4}
	l2 = []int{}

	l3 = interList(l1, l2)
	testCheckList(t, l3)

	l1 = []int{1, 2, 3}
	l2 = []int{2}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3)

	l1 = []int{1, 2, 4}
	l2 = []int{3}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{2, 3, 4}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{2}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 3, 4)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{1, 3, 5}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 2, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{1, 3, 5, 6}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{2, 3}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 4)
}
