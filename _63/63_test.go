package _63

import (
	"fmt"
	"testing"
)

type question63 struct {
	para63
	ans63
}

// para 是参数
// one 代表第一个参数
type para63 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans63 struct {
	one int
}

func Test_Problem63(t *testing.T) {

	qs := []question63{

		{
			para63{[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			ans63{2},
		},

		{
			para63{[][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			}},
			ans63{0},
		},

		{
			para63{[][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
			}},
			ans63{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 63------------------------\n")

	for _, q := range qs {
		a, p := q.ans63, q.para63

		r := uniquePathsWithObstacles(p.og)
		if r != a.one {
			t.Errorf("【input】:%v       【output】:%d    expect:%d \n", p, r, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
