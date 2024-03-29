package on299

func ReverseList(head *ListNode) *ListNode {
	var behind *ListNode

	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}

	return behind
}
