package probs99

import (
	"container/list"
	"fmt"
)

func lastElem(list []int) (int, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("slice is zero size")
	}
	return list[len(list)-1], nil
}
func lastButOne(list []int) (int, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("slice is zero size")
	} else if len(list) == 1 {
		return list[0], fmt.Errorf("slize is only length 1")
	} else {
		return list[len(list)-2], nil
	}

}

//note k=1 is 1st element in the list
func kthElement(list []int, k int) (int, error) {
	if k > len(list) {
		return 0, fmt.Errorf("k is greater than length of list")
	}
	indx := k - 1
	if indx < 0 {
		return 0, fmt.Errorf("element must be greater than 0 ")
	}
	return list[indx], nil

}
func numElements(lst []int) int {
	return len(lst)
}

func reverse(lst []int) []int {
	for i, j := 0, len(lst)-1; i < len(lst)/2; i, j = i+1, j-1 {
		lst[i], lst[j] = lst[j], lst[i]
	}
	return lst
}
func isPalindrome(lst []int) bool {
	for i, j := 0, len(lst)-1; i < len(lst)/2; i, j = i+1, j-1 {
		if lst[i] != lst[j] {
			return false
		}
	}
	return true
}

//note flatten multidim array
func flatten(in [][]int) []int {
	var out []int
	for _, row := range in {
		for _, v := range row {
			out = append(out, v)
		}
	}
	return out
}
func noDups(in []int) []int {
	res := []int{}
	if len(in) == 0 {
		return res
	}
	for i, j := 0, 1; i < len(in); i, j = i+1, j+1 {
		if j == len(in) {
			res = append(res, in[i])

		} else if in[i] != in[j] {
			res = append(res, in[i])

		}
	}
	return res
}

func helper() {
	l := list.New()
	//create some data
	d1 := []int{1, 2, 3}
	d2 := []int{4, 5}
	d3 := []int{6, 7, 8, 9}

	l.PushBack(d1)
	l.PushBack(d2)
	l.PushBack(d3)

	//iterate
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func consecutive2Sublist(data []int) *list.List {
	l := list.New()
	t := []int{}
	for i, j := 0, 1; i < len(data)-1; i, j = i+1, j+1 {
		//last case
		if j == len(data)-1 {
			if data[i] != data[j] {
				t = append(t, data[i])
				l.PushBack(t)
				t = []int{}
				t = append(t, data[j])
				l.PushBack(t)
			} else {
				t = append(t, data[i:]...)
				l.PushBack(t)
			}

		} else if data[i] == data[j] {
			t = append(t, data[i])
		} else {
			t = append(t, data[i])
			l.PushBack(t)
			t = []int{}
		}

	}
	return l
}

/*
func addList(l *list.List, data []int) {
	l.PushBack(data)

}*/

/*
func main() {
	lst := []int{1,2,3,4,5}
	r := reverse(lst)
	fmt.Println(r)
	fmt.Println(lst)
	//	bad := []int{}
	lst := []int{1, 2, 3, 4, 5}
	v, err := lastElem(lst)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	v, err = lastButOne(lst)
	if err != nil {
		fmt.Println(err)
	}
*/
