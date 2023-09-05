/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow = slow.Next
			slow.Val = fast.Val
		}

		fast = fast.Next
	}
	//断开多余的链表
	slow.Next = nil
	return head
}