
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//拆解为两部分：第一部分为反转链表前N个节点，第二部分为反转练链表区间

//反转链表前N个节点
var successor = &ListNode{}

func reverseN(head *ListNode, n int) *ListNode {

	if n == 1 {
		successor = head.Next
		return head
	}

	last := &ListNode{}

	last = reverseN(head.Next, n-1)

	head.Next.Next = head
	head.Next = successor
	return last
}

/*
 1->2->3->4->5->nil


 1->2->3->4->5->nil
       |
     successor: 4->5->nil

 1->2<-3  4->5->nil
    |
 successor: 4->5->nil

 1<-2<-3  4->5->nil
 |
 successor: 4->5->nil

*/

//反转练链表区间
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}
 
 