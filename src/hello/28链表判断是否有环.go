package main

type ListNode struct {
	val int
	Next *ListNode
}
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil{
		return nil
	}

	p := head
	fast := head.Next.Next
	slow := head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	for p != slow {
		p = p.Next
		slow = slow.Next
	}

	return p
}