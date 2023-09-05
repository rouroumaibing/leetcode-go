/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	left, right := head, reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	//&ListNode{}初始化会赋值val为0,导致交换值时覆盖原有数据
	tmp, next, cur := &ListNode{}, &ListNode{}, head
	pre := tmp.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}