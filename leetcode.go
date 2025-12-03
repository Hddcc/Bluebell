package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := ListNode{}
	l3 := &newList
	flag := 0
	for l1 != nil || l2 != nil || flag != 0 {
		if l1 != nil {
			flag += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			flag += l2.Val
			l2 = l2.Next
		}
		l3.Next = &ListNode{Val: flag % 10}
		flag /= 10
		l3 = l3.Next
	}
	return newList.Next
}
