/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	//假设相遇点距环的起点的距离为 m，想遇距离为k，
	//环的起点距头结点 head 的距离为 k - m，从相遇点继续前进 k - m 步，也恰好到达环起点
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		//相交，退出循环
		if fast == slow {
			break
		}
	}

	//循环结束，还是没有交点
	if fast == nil || fast.Next == nil {
		return nil
	}

	//循环遇到交点，再走k-m 步
	slow = head

	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}

	return slow

}