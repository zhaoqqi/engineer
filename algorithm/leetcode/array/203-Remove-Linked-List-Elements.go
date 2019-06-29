
/* Definition for singly-linked list.
*type ListNode struct {
*   Val int
*   Next *ListNode
*}
*/
 
func removeElements(head *ListNode, val int) *ListNode {
    
    // remove head element if equals to val
    for head!=nil && head.Val==val {
        head = head.Next
    }
    
    // remove narmal element if equals to val
    var cur *ListNode = head
    for cur != nil {
        if cur.Next == nil {
            break
        }
        if cur.Next.Val == val {
            if cur.Next.Next == nil {
                cur.Next = nil
                break
            } else {
                cur.Next = cur.Next.Next
            }
        } else {
            cur = cur.Next
        }
    }
    
    return head
}

