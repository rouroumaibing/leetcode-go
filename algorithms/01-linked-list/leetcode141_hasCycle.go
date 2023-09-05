/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	//快慢指针初始化指向head
	slow, fast := head, head
	//快指针走到末尾停止
	for fast != nil && fast.Next != nil {
		//慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
		//快慢指针相遇，说明有环
		if slow == fast {
			return true
		}
	}
	return false
}