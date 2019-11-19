package addtwonumbers

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	var carry int
	for l1 != nil || l2 != nil {
		val1 := func() int {
			if l1 == nil {
				return 0
			}
			return l1.Val
		}()
		val2 := func() int {
			if l2 == nil {
				return 0
			}
			return l2.Val
		}()

		sum := val1 + val2 + carry
		carry = sum / 10
		sum %= 10

		cur.Next = &ListNode{
			Val: sum,
		}
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		cur.Next = &ListNode{
			Val: 1,
		}
	}

	return pre.Next
}
