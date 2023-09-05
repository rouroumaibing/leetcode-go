/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB

	for p1 != p2 {
		//循环完A，循环B,结束后为nil
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		//循环完B，循环A，结束后为nil
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	// 有交点或是空，两个结果
	return p1
}