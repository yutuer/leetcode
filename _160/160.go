package _160

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var bOver, aOver = false, false
	curA, curB := headA, headB

	for ; curA != curB && curA != nil && curB != nil; {
		curA = curA.Next
		if curA == nil {
			if !bOver {
				curA = headB
				bOver = true
			}
		}

		curB = curB.Next
		if curB == nil {
			if !aOver {
				curB = headA
				aOver = true
			}
		}
	}
	return curA
}

type ListNode struct {
	Val  int
	Next *ListNode
}
