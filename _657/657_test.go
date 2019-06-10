package _657

import (
	"fmt"
	"testing"
)

func TestJudgeCircle1(t *testing.T) {
	circle := judgeCircle("UD")
	fmt.Println(circle)
}

func TestJudgeCircle2(t *testing.T) {
	circle := judgeCircle("LL")
	fmt.Println(circle)
}

func TestJudgeCircle3(t *testing.T) {
	circle := judgeCircle("RLUURDDDLU")
	fmt.Println(circle)
}
