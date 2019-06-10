package _999

func numRookCaptures(board [][]byte) int {
	aLen := len(board)
	var total int

	var rook *Node
	var nodes map[int]*Node

	for i := 0; i < aLen; i++ {
		bLen := len(board[i])
		for j := 0; j < bLen; j++ {
			b := board[i][j]
			if b == 'R' {
				rook = &Node{i, j, 'R'}
			} else if b == 'B' {
				nodes[i*aLen+j] = &Node{i, j, 'B'}
			} else if b == 'p' {
				nodes[i*aLen+j] = &Node{i, j, 'p'}
			}
		}
	}

	// 向左
	for i := rook.x*aLen + rook.y - 1; i >= rook.x*aLen; i-- {
		if v, ok := nodes[i]; ok {
			if v.c == 'B' {
				break
			} else if v.c == 'p' {
				total ++
				break
			}
		}
	}

	//向右
	for i := rook.x*aLen + rook.y + 1; i < (rook.x+1)*aLen; i++ {
		if v, ok := nodes[i]; ok {
			if v.c == 'B' {
				break
			} else if v.c == 'p' {
				total ++
				break
			}
		}
	}

	// 向上


	// 向下

	return total
}

type Node struct {
	x, y int
	c    byte
}
