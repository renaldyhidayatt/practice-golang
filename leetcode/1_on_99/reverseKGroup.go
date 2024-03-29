package on99

func ReverseKGroup(head *ListNode, k int) *ListNode {
	node := head

	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}

		node = node.Next
	}

	newHead := reverse(head, node)

	head.Next = ReverseKGroup(node, k)

	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last

	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}

	return prev
}
