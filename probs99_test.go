package probs99

import "testing"

func TestLastElem(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	v, err := lastElem(lst)
	if err != nil {
		t.Errorf("%v", err)
	}
	if v != 5 {
		t.Error("error:Value not equal to last")
	}
	t.Logf("good:value is %v", v)
	nodata := []int{}
	v, err = lastElem(nodata)
	if err == nil {
		t.Errorf("error:should yield err %v", err)
	}
	t.Logf("good: should be %v", err)

}

func TestLastButOne(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	v, err := lastButOne(lst)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if v != 4 {
		t.Errorf("error: should be 4 got %v", v)
	}
	t.Logf("good value is %v", v)

	lst0 := []int{}
	v, err = lastButOne(lst0)
	if err == nil {
		t.Errorf("error: should be error from zero size")
	}
	lst1 := []int{1}
	v, err = lastButOne(lst1)
	if err == nil {
		t.Errorf("error: should be error from one size")
	}

}
func TestKthElement(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	v, err := kthElement(lst, 3)
	if err != nil {
		t.Errorf("error: getting 3rd element")
	}
	if v != 3 {
		t.Errorf("error:Value should be 3")
	}
	v, err = kthElement(lst, 0)
	if err == nil {
		t.Errorf("error:index must be greater than zero")
	}
	v, err = kthElement(lst, 10)
	if err == nil {
		t.Errorf("error index is greater than length of array")
	}
	oth := []int{}

	v, err = kthElement(oth, 9)
	if err == nil {
		t.Errorf("error:index is greater than array length")
	}

}
func TestNumElements(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	v := numElements(lst)
	if v != 5 {
		t.Errorf("Number of elements should be 5")
	}

	lst2 := []int{}
	v = numElements(lst2)
	if v != 0 {
		t.Errorf("Number of elements should be 0")
	}

}
func TestReverse(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5} //odd length
	r := reverse(lst)
	if r[0] != 5 || r[1] != 4 || r[2] != 3 || r[3] != 2 || r[4] != 1 {
		t.Errorf("failed to reverse odd length")
	}
	if lst[0] != 5 || lst[1] != 4 || lst[2] != 3 || lst[3] != 2 || lst[4] != 1 {
		t.Errorf("failed to reverse odd length")
	}
	lst2 := []int{1, 2, 3, 4, 5, 6} //even length
	r2 := reverse(lst2)
	if r2[0] != 6 || r2[1] != 5 || r2[2] != 4 || r2[3] != 3 || r2[4] != 2 || r2[5] != 1 {
		t.Logf("list is %v", r2)
		t.Errorf("failed to reverse even length")
	}
}
func TestPalindrome(t *testing.T) {
	lst := []int{1, 2, 3, 2, 1}
	r := isPalindrome(lst)
	if r != true {
		t.Errorf("should be palindrome")
	}
	lst2 := []int{1, 2, 3, 1, 2}
	r = isPalindrome(lst2)
	if r != false {
		t.Errorf("this is not a palindrome")
	}
}

func TestFlatten(t *testing.T) {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r := flatten(arr)
	if len(r) != 9 {
		t.Errorf("error:length should be 9")
	}
}

func TestNoDuplicates(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 5, 5, 5, 6, 6}
	r := noDups(arr)
	if len(r) != 6 {
		t.Errorf("length should be 6")
	}
	t.Logf("result is %v", r)
	arr2 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	r = noDups(arr2)
	if len(r) != 1 {
		t.Errorf("length should be 1")
	}
	t.Logf("result is %v", r)
}

func TestListHelper(t *testing.T) {
	helper()
}
func TestConsec2SubList(t *testing.T) {
	d := []int{1, 2, 3, 4, 5, 5, 5, 6}
	l := consecutive2Sublist(d)
	count := 0
	for e := l.Front(); e != nil; e = e.Next() {
		count++
		//		t.Logf("%v", e.Value)
	}
	if count != 6 {
		t.Errorf("Number of groups should be 6")
	}
	count = 0
	d2 := []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 6}

	l = consecutive2Sublist(d2)
	for e := l.Front(); e != nil; e = e.Next() {
		count++
		t.Logf("%v", e.Value)
	}
	if count != 6 {
		t.Errorf("Number of groups should be 6")
	}

}
