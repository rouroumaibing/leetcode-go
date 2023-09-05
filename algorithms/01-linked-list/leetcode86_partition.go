/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func partition(head *ListNode, x int) *ListNode {
    dummy1 := &ListNode{}
    dummy2 := &ListNode{}
    tmp := &ListNode{}

    //初始化两条空链表，一条用于存放小于x，一条用于存放大于x
    p, p1, p2 := head, dummy1, dummy2

    for p != nil  {
        if p.Val < x {
            p1.Next = p
            p1 = p1.Next
        } else {
            p2.Next = p
            p2 = p2.Next
        }
        //断开原链表中的每个节点的next指针
        tmp = p.Next
        p.Next = nil
        p = tmp

    }
    p1.Next = dummy2.Next

    return dummy1.Next
}