package _206

import (
	"bytes"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}

	//var h *ListNode
	//
	//for he := head; he != nil; {
	//	next := he.Next
	//	he.Next = h
	//	h = he
	//	he = next
	//}
	//return h

	first, tail := a(head)
	first.Next = nil
	return tail
}

func a(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	n, tail := a(head.Next)
	n.Next = head
	return head, tail
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var str bytes.Buffer
	for nn := n; nn != nil; nn = nn.Next {
		str.WriteString(strconv.Itoa(nn.Val))
		str.WriteString("->")
	}
	str.WriteString("NULL")
	return str.String()
}

func transfer(arr []int) *ListNode {
	var head *ListNode
	var cur *ListNode
	for _, v := range arr {
		n := &ListNode{v, nil}
		if head == nil {
			head = n
		} else {
			cur.Next = n
		}
		cur = n
	}
	return head
}
