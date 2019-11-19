package addtwonumbers

import (
	"strconv"
)

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := node2Int(l1)
	num2 := node2Int(l2)
	num := *num1 + *num2

	return int2Node(num)
}

func int2Node(num int) *ListNode {
	numstr := strconv.Itoa(num)

	var res *ListNode
	var pre *ListNode
	for i := 1; i <= len(numstr); i++ {
		item := num % 10
		num = num / 10
		node := &ListNode{
			Val: item,
		}

		if res == nil {
			res = node
		}

		if pre != nil {
			pre.Next = node
		}
		pre = node
	}

	return res
}

func node2Int(l *ListNode) *int {
	if l.Next == nil {
		return &l.Val
	}

	str := strconv.Itoa(l.Val)
	next := node2Int(l.Next)
	if next != nil {
		tmp := strconv.Itoa(*next)
		str = tmp + str
	}

	res, _ := strconv.Atoi(str)
	return &res
}
