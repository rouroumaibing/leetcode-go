/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	//快指针走两步，慢指针走一步，同时走，快指针到链表末尾，慢指针就刚好到中点了。
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}