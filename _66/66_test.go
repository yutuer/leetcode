package _66

import (
	"fmt"
	"reflect"
	"testing"
)

type question66 struct {
	para66
	ans66
}

// para 是参数
// one 代表第一个参数
type para66 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans66 struct {
	one []int
}

func Test_Problem66(t *testing.T) {

	qs := []question66{

		{
			para66{[]int{1, 2, 3}},
			ans66{[]int{1, 2, 4}},
		},

		{
			para66{[]int{4, 3, 2, 1}},
			ans66{[]int{4, 3, 2, 2}},
		},

		{
			para66{[]int{9, 9}},
			ans66{[]int{1, 0, 0}},
		},

		{
			para66{[]int{0}},
			ans66{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 66------------------------\n")

	for _, q := range qs {
		a, p := q.ans66, q.para66

		r := plusOne(p.one)
		if !reflect.DeepEqual(r, a.one) {
			t.Errorf("【input】:%v       【output】:%v   expect:%v \n", p, r, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
