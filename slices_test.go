package slices

import (
	"testing"
	"fmt"
)

func Test_All(t *testing.T) {
	Test_Int_Append(t)
	Test_Int_Insert(t)
	Test_Int_Extract(t)
}

func Test_Int_Insert(t *testing.T) {
	a := []int{1, 2, 3}
	a = Int.Insert(a, 2, 10, 11, 12)
	fmt.Printf("%v\n", a)

	if len(a) == 6 && a[1] == 2 && a[2] == 10 && a[4] == 12 && a[5] == 3 {
		// Pass
	} else {
		t.Fail()
	}

	a = []int{1, 2, 3}
	a = Int.InsertSlice(a, 2, []int{10, 11, 12})
	fmt.Printf("%v\n", a)

	if len(a) == 6 && a[1] == 2 && a[2] == 10 && a[4] == 12 && a[5] == 3 {
		// Pass
	} else {
		t.Fail()
	}
}

func Test_Int_Append(t *testing.T) {
	a := []int{1, 2, 3}
	a = Int.Append(a, 10, 11, 12)
	fmt.Printf("%v\n", a)

	if len(a) == 6 && a[2] == 3 && a[3] == 10 && a[4] == 11 && a[5] == 12 {
		// Pass
	} else {
		t.Fail()
	}

	a = []int{1, 2, 3}
	a = Int.AppendSlice(a, []int{10, 11, 12})
	fmt.Printf("%v\n", a)

	if len(a) == 6 && a[2] == 3 && a[3] == 10 && a[4] == 11 && a[5] == 12 {
		// Pass
	} else {
		t.Fail()
	}
}

func Test_Int_Extract(t *testing.T) {
	a := []int{1, 2, 3}
	e, a := Int.Extract(a, 1)
	fmt.Printf("%v, %v\n", e, a)

	if len(a) == 2 && a[0] == 1 && a[1] == 3 && e == 2 {
		// Pass
	} else {
		t.Fail()
	}

	a = []int{1, 2, 3}
	l, a := Int.ExtractBy(a, (func(m int) bool {
		return m < 3
	}))
	fmt.Printf("%v, %v\n", l, a)

	if len(l) == 2 && len(a) == 1 && a[0] == 3 && l[0] == 1 && l[1] == 2 {
		// Pass
	} else {
		t.Fail()
	}
}
