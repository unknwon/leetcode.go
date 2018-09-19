/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	var last *ListNode
	carry := false
	for l1 != nil || l2 != nil {
		l := &ListNode{}
		if l1 == nil {
			last.Next = l2
			break
		} else if l2 == nil {
			last.Next = l1
			break
		} else {
			l.Val = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}

		if carry {
			l.Val++
		}

		if l.Val >= 10 {
			carry = true
			l.Val -= 10
		} else {
			carry = false
		}

		if last == nil {
			last = l
			head.Next = last
		} else {
			last.Next = l
			last = l
		}
	}

	for carry {
		if last.Next == nil {
			last.Next = &ListNode{Val: 1}
			break
		}

		last.Next.Val++
		if last.Next.Val >= 10 {
			carry = true
			last.Next.Val -= 10
			last = last.Next
		} else {
			carry = false
		}
	}
	return head.Next
}