/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(left *ListNode, right *ListNode) *ListNode {
    pre, cur, nxt := &ListNode{}, left, left

    for cur != right {
        // 1 -> 2 -> 3 ->4 ->5
        // pre = nil, cur = 1, nxt = 1, cur.Next= 2
        nxt = cur.Next
        //pre = nil, cur = 1, nxt = 2, cur.Next = 2
        cur.Next = pre
        /*  1  2->3->4 ->5
            |
           nil
        pre = nil, cur = 1, nxt = 2, cur.Next = nil
        */
        pre = cur
        //pre = 1 , cur = 1, nxt = 2, cur.Next = nil
        cur = nxt
        //pre = 1, cur = 2, nxt = 2, cur.Next = 3
        // 1->nil  2->3->4 ->5
        // 2->1->nil
    }
    return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    left, right := head, head
    
    //不足K个不需要反转，直接返回
    for i := 0; i < k; i++ {
        if right == nil {
            return head
        }
        right = right.Next 
    }

    newHead := reverse(left, right)
    //递归后的值作为头，和后继链表连接起来
    left.Next = reverseKGroup(right, k)
    return newHead

}