/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    //&{0 <nil>}
    p := dummy
    
    p1, p2 := list1, list2

    // 将值较小的的节点接到 p 指针
    for p1 != nil && p2 != nil {
        if p1.Val > p2.Val {
            p.Next = p2
            p2 = p2.Next
        }else {
            p.Next = p1
            p1 = p1.Next
        }
        //p指针不断前进
        p = p.Next
    }

    if p1 == nil{
        p.Next = p2
    }

    if p2 == nil{
        p.Next = p1
    }
    //dummy 包含初始化的阈值
    //dummy = [0, ...]
    //dummy.Next = [...]
    return dummy.Next
}