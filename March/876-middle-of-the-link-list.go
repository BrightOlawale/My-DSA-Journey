/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package march

type ListNode struct {
    Val int
    Next *ListNode
}
 func middleNode(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    fast := head
    slow := head

    for fast.Next != nil {
        fast = fast.Next

        if fast.Next != nil {
            fast = fast.Next
        }
        slow = slow.Next
    }

    return slow
}