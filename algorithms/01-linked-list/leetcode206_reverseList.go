/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := &ListNode{}

	last = reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil
	return last
}

/*
1->2->3->4->5->nil

1->2->3->4->5<-nil
            |
            nil

1->2->3->4<-5<-nil
         |
         nil

1->2->3<-4<-5<-nil
      |
     nil

1->2<-3<-4<-5<-nil
   |
  nil

1<-2<-3<-4<-5<-nil
|
nil
*/