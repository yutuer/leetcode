package _116

import "strconv"

type Node struct {
	val   int;
	left  *Node;
	right *Node;
	next  *Node;
}

func connect(root *Node) *Node {
	var n *Node = root
	for n != nil {

	}
	return nil
}

func (n *Node) String() string {
	return strconv.Itoa(n.val)
}
